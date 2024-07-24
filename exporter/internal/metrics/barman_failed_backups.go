package metrics

import (
	"barman-exporter/internal/integration"
	"context"
	"log/slog"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
)

// Outputs the current number of failed backups for specific server
type BarmanFailedBackups struct {
	metric      metric.Int64Gauge
	integration integration.BarmanIntegration
}

func NewBarmanFailedBackups(meter metric.Meter, integration integration.BarmanIntegration) (BarmanMetric, error) {
	bfb := &BarmanFailedBackups{integration: integration}
	err := bfb.Init(meter)
	return bfb, err
}

func (bfb *BarmanFailedBackups) Init(meter metric.Meter) error {
	m, err := meter.Int64Gauge("barman_backups_failed", metric.WithDescription("Outputs the current number of backups with status FAILED for specific server"))
	if err != nil {
		slog.Warn("Failed to create barman_backups_failed metric. %v", err)
		return err
	}
	bfb.metric = m
	return nil
}

// integration.GetAllBackups() and then iterates through the returned objects(range server.Backups), if the Status field is FAILED counter is incremented.
// Finally the counter Record in the metric
func (bfb *BarmanFailedBackups) Update(ctx context.Context) {
	allServersBackups, err := bfb.integration.GetAllBackups(ctx)
	if err != nil {
		slog.Warn("Failed to update barman_failed_backups metric. Error in retrieving data from barman integration: %v", err)
		return
	}
	for _, server := range allServersBackups {
		serverAttribute := metric.WithAttributes(attribute.String("server", string(server.Server)))
		var failedBackupsCount int64 = 0
		for _, backup := range server.Backups {
			if backup.Status == "FAILED" {
				failedBackupsCount++
			}
		}
		bfb.metric.Record(ctx, failedBackupsCount, serverAttribute)
	}
}
