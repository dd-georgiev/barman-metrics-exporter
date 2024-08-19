package metrics

import (
	"barman-exporter/internal/integration"
	"context"
	"log/slog"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
)

// Reports how much time it took to finish the last backup, as reported by barman show command.
type BarmanLastBackupWalFiles struct {
	metric      metric.Int64Gauge
	integration integration.BarmanIntegration
}

func NewBarmanLastBackupWalFiles(meter metric.Meter, integration integration.BarmanIntegration) (BarmanMetric, error) {
	blbwf := &BarmanLastBackupWalFiles{integration: integration}
	err := blbwf.Init(meter)
	return blbwf, err
}
func (blbwf *BarmanLastBackupWalFiles) Init(meter metric.Meter) error {
	m, err := meter.Int64Gauge("barman_last_backup_wal_files", metric.WithDescription("outputs the number of wals for the last backup"))
	if err != nil {
		slog.Warn("Failed to create barman_last_backup_wal_files metric. %v", err)
		return err
	}
	blbwf.metric = m
	return nil
}

// Gets show for all servers latest backups using the integration GetShowForLatestBackupForEachServer()
// and then records the WalRatePerSecond as provided by Barman
func (blbwf *BarmanLastBackupWalFiles) Update(ctx context.Context) {
	showForLatest, err := blbwf.integration.GetShowForLatestBackupForEachServer(ctx)
	if err != nil {
		slog.Warn("Failed to update barman_last_backup_wal_files metric. Error in retrieving data from barman integration: %v", err)
		return
	}
	for _, server := range showForLatest {
		serverAttribute := metric.WithAttributes(attribute.String("server", string(server.Server)))
		blbwf.metric.Record(ctx, server.Show.WalFiles, serverAttribute)
	}
}
