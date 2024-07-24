package metrics

import (
	"barman-exporter/internal/integration"
	"context"
	"fmt"
	"log/slog"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
)

// OOutputs the size of a the wals for each server and each backup of the said server in bytes
type BarmanBackupWalSize struct {
	metric      metric.Int64Gauge
	integration integration.BarmanIntegration
}

func NewBarmanBackupWalSize(meter metric.Meter, integration integration.BarmanIntegration) (BarmanMetric, error) {
	bbws := &BarmanBackupWalSize{integration: integration}
	err := bbws.Init(meter)
	return bbws, err
}

// Creates Int64Gauge
func (bbws *BarmanBackupWalSize) Init(meter metric.Meter) error {
	m, err := meter.Int64Gauge("barman_backup_wal_size", metric.WithDescription("Outputs the size of a the wals for specific backup(for specific server) in bytes"))
	if err != nil {
		slog.Warn("Failed to create barman_backup_wal_size metric. %v", err)
		return err
	}
	bbws.metric = m
	return nil
}

// Gets all backups using the integration, filters all backups which are in "DONE" state and records the WalSizeBytes for them
func (bbws *BarmanBackupWalSize) Update(ctx context.Context) {
	allServersBackups, err := bbws.integration.GetAllBackups(ctx)
	if err != nil {
		slog.Warn("Failed to update barman_backup_wal_size metric. Error in retrieving data from barman integration: %v", err)
		return
	}
	for _, server := range allServersBackups {
		serverAttribute := metric.WithAttributes(attribute.String("server", string(server.Server)))
		for i, backup := range server.Backups {
			if backup.Status != "DONE" {
				continue
			}
			backupNumberAttribute := metric.WithAttributes(attribute.String("number", fmt.Sprint(i)))
			bbws.metric.Record(ctx, backup.WalSizeBytes, serverAttribute, backupNumberAttribute)
		}
	}
}
