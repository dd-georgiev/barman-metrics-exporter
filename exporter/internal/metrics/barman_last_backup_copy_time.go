package metrics

import (
	"barman-exporter/internal/integration"
	"context"
	"log/slog"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
)

// Reports how much time it took to finish the last backup, as reported by barman show command.
type BarmanLastBackupCopyTime struct {
	metric      metric.Float64Gauge
	integration integration.BarmanIntegration
}

func NewBarmanLastBackupCopyTime(meter metric.Meter, integration integration.BarmanIntegration) (BarmanMetric, error) {
	blbct := &BarmanLastBackupCopyTime{integration: integration}
	err := blbct.Init(meter)
	return blbct, err
}
func (blbct *BarmanLastBackupCopyTime) Init(meter metric.Meter) error {
	m, err := meter.Float64Gauge("barman_last_backup_copy_time", metric.WithDescription("Outputs the time it took to get the latest backup in seconds"))
	if err != nil {
		slog.Warn("Failed to create barman_last_backup_copy_time metric. %v", err)
		return err
	}
	blbct.metric = m
	return nil
}

// Gets show for all servers latest backups using the integration GetShowForLatestBackupForEachServer()
// and then records the CopyTimeSeconds as provided by Barman
func (blbct *BarmanLastBackupCopyTime) Update(ctx context.Context) {
	showForLatest, err := blbct.integration.GetShowForLatestBackupForEachServer(ctx)
	if err != nil {
		slog.Warn("Failed to update barman_last_backup_copy_time metric. Error in retrieving data from barman integration: %v", err)
		return
	}
	for _, server := range showForLatest {
		serverAttribute := metric.WithAttributes(attribute.String("server", string(server.Server)))
		blbct.metric.Record(ctx, server.Show.CopyTimeSeconds, serverAttribute)
	}
}
