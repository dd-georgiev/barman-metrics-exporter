package metrics

import (
	"barman-exporter/internal/integration"
	"context"
	"log/slog"
	"time"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
)

// Reports how much time has passed since the first backup was taken, for a given server
type BarmanFirstBackup struct {
	metric      metric.Int64Gauge
	integration integration.BarmanIntegration
}

func NewBarmanFirstBackup(meter metric.Meter, integration integration.BarmanIntegration) (BarmanMetric, error) {
	bfb := &BarmanFirstBackup{integration: integration}
	err := bfb.Init(meter)
	return bfb, err
}

func (bfb *BarmanFirstBackup) Init(meter metric.Meter) error {
	m, err := meter.Int64Gauge("barman_first_backup", metric.WithDescription("Time since the first backup was taken in seconds"))
	if err != nil {
		slog.Warn("Failed to create barman_first_backup metric. %v", err)
		return err
	}
	bfb.metric = m
	return nil
}

// Gets `show BACKUP` from the ingtegration, where BACKUP is the first one provided by the status command.
// The for each of the servers, the first backup EndTimeTimestamp is taken and substracted from time.Now(using time.Since method)
func (bfb *BarmanFirstBackup) Update(ctx context.Context) {
	lastBackups, err := bfb.integration.GetShowForFirstBackupForEachServer(ctx)
	if err != nil {
		slog.Warn("Failed to update barman_first_backup metric. Error in retrieving data from barman integration: %v", err)
		return
	}
	for _, server := range lastBackups {
		serverAttribute := metric.WithAttributes(attribute.String("server", string(server.Server)))
		t := time.Unix(server.Show.EndTimeTimestamp, 0)
		diff := time.Since(t)
		bfb.metric.Record(ctx, (diff.Milliseconds() / int64(1000)), serverAttribute)
	}
}
