package metrics

import (
	"barman-exporter/internal/integration"
	"context"
	"fmt"
	"log/slog"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
)

// Outputs the current number of available backups for specific server
type BarmanBackupSize struct {
	metric      metric.Int64Gauge
	integration integration.BarmanIntegration
}

func NewBarmanBackupSize(meter metric.Meter, integration integration.BarmanIntegration) (BarmanMetric, error) {
	bbs := &BarmanBackupSize{integration: integration}
	err := bbs.Init(meter)
	return bbs, err
}

// Creates Int64Gauge
func (bbs *BarmanBackupSize) Init(meter metric.Meter) error {
	m, err := meter.Int64Gauge("barman_backup_size", metric.WithDescription("Outputs the size of a specific backup(for specific server) in bytes"))
	if err != nil {
		slog.Warn("Failed to create barman_backup_size metric. %v", err)
		return err
	}
	bbs.metric = m
	return nil
}

func (bbs *BarmanBackupSize) Update(ctx context.Context) {
	allServersBackups, err := bbs.integration.GetAllBackups(ctx)
	if err != nil {
		slog.Warn("Failed to update barman_backup_size metric. Error in retrieving data from barman integration: %v", err)
		return
	}
	for _, server := range allServersBackups {
		serverAttribute := metric.WithAttributes(attribute.String("server", string(server.Server)))
		for i, backup := range server.Backups {
			if backup.Status != "DONE" {
				continue
			}
			backupNumberAttribute := metric.WithAttributes(attribute.String("number", fmt.Sprint(i)))
			bbs.metric.Record(ctx, backup.SizeBytes, serverAttribute, backupNumberAttribute)
		}
	}
}
