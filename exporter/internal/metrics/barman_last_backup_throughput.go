package metrics

import (
	"barman-exporter/internal/integration"
	"context"
	"log/slog"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
)

// Reports how much time it took to finish the last backup, as reported by barman show command.
type BarmanLastBackupThroughput struct {
	metric      metric.Float64Gauge
	integration integration.BarmanIntegration
}

func NewBarmanLastBackupThroughput(meter metric.Meter, integration integration.BarmanIntegration) (BarmanMetric, error) {
	blbtp := &BarmanLastBackupThroughput{integration: integration}
	err := blbtp.Init(meter)
	return blbtp, err
}
func (blbtp *BarmanLastBackupThroughput) Init(meter metric.Meter) error {
	m, err := meter.Float64Gauge("barman_last_backup_throughput", metric.WithDescription("Outputs the throughput(in bytes/second) during the last backup creation"))
	if err != nil {
		slog.Warn("Failed to create barman_last_backup_throughput metric. %v", err)
		return err
	}
	blbtp.metric = m
	return nil
}

// Gets show for all servers latest backups using the integration GetShowForLatestBackupForEachServer()
// and then records the WalRatePerSecond as provided by Barman
func (blbtp *BarmanLastBackupThroughput) Update(ctx context.Context) {
	showForLatest, err := blbtp.integration.GetShowForLatestBackupForEachServer(ctx)
	if err != nil {
		slog.Warn("Failed to update barman_last_backup_throughput metric. Error in retrieving data from barman integration: %v", err)
		return
	}
	for _, server := range showForLatest {
		serverAttribute := metric.WithAttributes(attribute.String("server", string(server.Server)))
		blbtp.metric.Record(ctx, server.Show.ThroughputBytes, serverAttribute)
	}
}
