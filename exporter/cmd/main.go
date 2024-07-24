package main

import (
	"barman-exporter/internal/integration"
	"barman-exporter/internal/metrics"
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/yamlv3"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/otel/exporters/prometheus"
	"go.opentelemetry.io/otel/sdk/metric"
)

const UpdateIntervalInSeconds = 5

/*
Handles:
 1. Configuration setup (loads config file passed via CLI argument)
 2. Setups logging level
 3. Setups metrics and integration
 4. Starts metrics update loop
 5. Handles OS signals
    5.1 SIGUSR1 - Reload configuration without restart. Doesn't change executor type. This action restarts the update loop.
    5.2 SIGTERM, & interrupt(ctrl+c) - kill the program.
*/
func main() {
	ctx := context.Background()

	initConfig()

	initLogging()

	allMetrics := initMetrics(ctx)

	go serveMetrics()
	// This variable is used to stop the update loop in case of config reload. This is needed since the 'refresh_interval' config val may be changed
	continueUpdating := make(chan string)

	go startUpdateLoop(ctx, continueUpdating, &allMetrics)

	// hot reload - the app reloads the config on SIGUSR1.
	// This is because it may be runned as normal executable in shell.
	// SystemD sends SIGHUP to trigger config reload, this behavior must be overwritten in the systemd unit file
	// NOTE: Changing the executor type is not supported. The only use case for this is testing, but there are easier ways to test if the configuration will be updated
	// 		 Such as changing the log level up/down/up(i.e. debug, warn, debug)
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGUSR1)
	go func() {
		for range sigs {
			slog.Warn("Detected SIGUSR1, will reload config")
			// reload config and reinit logging to change level[if the config has changed]
			err := config.LoadFiles(config.String("config"))
			if err != nil {
				log.Panic(err)
			}

			initLogging()
			// Stop and start update loop, in case the refresh_interval changed.
			// No need to block until the original/first loop finishes as the caching means that the impact of having two updates one after another means there is very little performance impact
			continueUpdating <- "config reload"
			go startUpdateLoop(ctx, continueUpdating, &allMetrics)
		}
	}()

	// handle OS interrupts(ctrl+c) and SIGTERM
	ctx, _ = signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
	<-ctx.Done()
}
func initConfig() {
	config.WithOptions(config.ParseEnv)
	config.AddDriver(yamlv3.Driver)
	keys := []string{"config"}
	err := config.LoadFlags(keys)
	if !config.Exists("config") {
		log.Fatal("Missing -config option")
	}
	if err != nil {
		log.Panic(err)
	}
	err = config.LoadFiles(config.String("config"))
	if err != nil {
		log.Panic(err)
	}
}
func initLogging() {
	switch logLevelFromConfig := strings.ToUpper(config.String("log_level")); logLevelFromConfig {
	case "ERROR":
		slog.SetLogLoggerLevel(slog.LevelError)
	case "WARN":
		slog.SetLogLoggerLevel(slog.LevelWarn)
	case "INFO":
		slog.SetLogLoggerLevel(slog.LevelInfo)
	case "DEBUG":
		slog.SetLogLoggerLevel(slog.LevelDebug)
	default:
		slog.SetLogLoggerLevel(slog.LevelWarn)
	}
}

func initMetrics(ctx context.Context) []metrics.BarmanMetric {
	exporter, err := prometheus.New(prometheus.WithoutScopeInfo(), prometheus.WithoutUnits())
	if err != nil {
		log.Fatal(err)
	}
	provider := metric.NewMeterProvider(metric.WithReader(exporter))
	meter := provider.Meter("barman-expoter")
	barmanIntegration := integration.NewIntegration(getExecutorFromConfig())

	allMetrics := metrics.InitAllMetrics(meter, barmanIntegration)
	return allMetrics
}
func serveMetrics() {
	addr := fmt.Sprintf("%s:%s", config.String("address"), config.String("port"))

	slog.Info("Started serving metrics", slog.String("address", addr))
	http.Handle("/metrics", promhttp.Handler())
	err := http.ListenAndServe(addr, nil) //nolint:gosec // Ignoring G114: Use of net/http serve function that has no support for setting timeouts.
	if err != nil {
		fmt.Printf("error serving http: %v", err)
		return
	}
}

func startUpdateLoop(ctx context.Context, stop chan string, allMetrics *[]metrics.BarmanMetric) {
	slog.Info("Starting update loop")
	updateInterval := time.Duration(config.Int("refresh_interval")) * time.Second
	go metrics.UpdateAllMetrics(ctx, *allMetrics) // initial retrival
	for range time.Tick(updateInterval) {
		select {
		case s := <-stop:
			slog.Warn("Stopping update loop, ", "reason", s)
			return
		default:
			go metrics.UpdateAllMetrics(ctx, *allMetrics)
		}
	}
}

func getExecutorFromConfig() integration.CommandExecutor {
	switch desiredIntegration := config.String("integration_type"); desiredIntegration {
	case "mock":
		return &integration.MockExecutor{}
	case "shell":
		return &integration.ShellExecutor{}
	default:
		panic("Unknown exeuctor type")
	}
}
