// Manages OTEL metrics.
package metrics

import (
	"barman-exporter/internal/integration"
	"context"
	"log/slog"
	"sync"
	"time"

	"go.opentelemetry.io/otel/metric"
)

// Abstraction on top of the otel types which focuses only on Updating and Initing(in registry) the metrics
// In addition to the abstraction it provides the barman_metrics_update. The metric is stored as global variable in this module.
type BarmanMetric interface {
	// Updates the metric to reflect the latest value.
	// For example the Update method of barman_total_backups metric takes all backups via the barman integration,
	// counts them and returns stores the count in the Int64Gauge created by Init.
	Update(ctx context.Context)
	// The Init metric  creates the metris using otel Meter, any auxiliry actions related to succesful metric creation are done here.
	Init(meter metric.Meter) error
}

// Type signature for function which returns new BarmanMetric, this must be the constructor for all types implementing the BarmanMetric interface
type NewMetricFunc func(meter metric.Meter, integration integration.BarmanIntegration) (BarmanMetric, error)

// Global variable used for reporting how long it took to execute the UpdateAll method.
// Other than the fact that this metric is definied in this file, nothing else differes - its inited in the InitAllMetrics and updated by UpdateAllMetrics
var updateTimeMetric metric.Int64Gauge

// Creates instances of all BarmanMetric types and returns am slice with them
func InitAllMetrics(meter metric.Meter, integration integration.BarmanIntegration) []BarmanMetric {
	metrics := make([]BarmanMetric, 0)
	slog.Debug("Initializing all metrics")
	start := time.Now()

	initMetric(meter, integration, &metrics, NewBarmanFailedBackups)
	initMetric(meter, integration, &metrics, NewBarmanTotalBackups)
	initMetric(meter, integration, &metrics, NewBarmanUp)
	initMetric(meter, integration, &metrics, NewBarmanBackupSize)
	initMetric(meter, integration, &metrics, NewBarmanBackupWalSize)
	initMetric(meter, integration, &metrics, NewBarmanLastBackupCopyTime)
	initMetric(meter, integration, &metrics, NewBarmanLastBackup)
	initMetric(meter, integration, &metrics, NewBarmanFirstBackup)
	initMetric(meter, integration, &metrics, NewBarmanLastBackupWalRatePerSecond)
	initMetric(meter, integration, &metrics, NewBarmanLastBackupThroughput)
	initMetric(meter, integration, &metrics, NewBarmanLastBackupWalFiles)
	initMetricUpdateTimeMetric(meter)

	elapsed := time.Since(start)
	slog.Debug("Finished initializing all metrics", slog.Duration("in", elapsed))
	return metrics
}

// Wrapper which eases error handling, if an metric fails to initialize warning is logged.
// Having problem with some of the metric is not a reason to stop the whole exporter.
func initMetric(meter metric.Meter, integration integration.BarmanIntegration, currentMetrics *[]BarmanMetric, newFunc NewMetricFunc) {
	newMetric, err := newFunc(meter, integration)
	if err != nil {
		slog.Warn("Failed to init metric. %v", err)
		return
	}
	*currentMetrics = append(*currentMetrics, newMetric)
}

// Goes through all metrics in the provided argument and call their Update method using go routine.
// WaitGroup is used so that the function will return when all metrics have finished updating
// The method takes as long as the slowest metric update to finish.
// The duration of each execution is repored as barman_metrics_update in ms.
func UpdateAllMetrics(ctx context.Context, metrics []BarmanMetric) {
	var allMetricsWaitGroup sync.WaitGroup
	slog.Debug("Updating all metrics")
	start := time.Now()
	for _, metric := range metrics {

		allMetricsWaitGroup.Add(1)
		go func() {
			defer allMetricsWaitGroup.Done()
			metric.Update(ctx)
		}()
	}
	allMetricsWaitGroup.Wait()
	elapsed := time.Since(start)
	updateTimeUpdateMetric(ctx, elapsed.Milliseconds())
	slog.Debug("Finished updating all metrics", slog.Duration("in", elapsed))
}

// Takes care of creating the barman_metrics_update with proper description
func initMetricUpdateTimeMetric(meter metric.Meter) error {
	m, err := meter.Int64Gauge("barman_metrics_update", metric.WithDescription("Outputs the time(in ms) it took update all metrics exporter by this exporter"))
	if err != nil {
		slog.Warn("Failed to create barman_metrics_update metric. %v", err)
		return err
	}
	updateTimeMetric = m
	return nil
}

// Records the duration of UpdateAll method as passed by the val argument. The val its self is measurement done in the UpdateAll method.
func updateTimeUpdateMetric(ctx context.Context, val int64) {
	updateTimeMetric.Record(ctx, val)
}
