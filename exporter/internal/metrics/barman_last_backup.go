package metrics

import (
	"barman-exporter/internal/integration"
	"context"
	"log/slog"
	"time"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
)

// Outputs the time which has passed since the last backup was taken
type BarmanLastBackup struct {
	metric      metric.Int64Gauge
	integration integration.BarmanIntegration
}

func NewBarmanLastBackup(meter metric.Meter, integration integration.BarmanIntegration) (BarmanMetric, error) {
	blb := &BarmanLastBackup{integration: integration}
	err := blb.Init(meter)
	return blb, err
}

func (blb *BarmanLastBackup) Init(meter metric.Meter) error {
	m, err := meter.Int64Gauge("barman_last_backup", metric.WithDescription("Time since the last backup was taken in seconds"))
	if err != nil {
		slog.Warn("Failed to create barman_last_backup metric. %v", err)
		return err
	}
	blb.metric = m
	return nil
}

// Gets `show BACKUP` from the ingtegration, where BACKUP is the last one provided by the status command.
// The for each of the servers, the last backup EndTimeTimestamp is taken and substracted from time.Now(using time.Since method)
func (blb *BarmanLastBackup) Update(ctx context.Context) {
	lastBackups, err := blb.integration.GetShowForLatestBackupForEachServer(ctx)
	if err != nil {
		slog.Warn("Failed to update barman_last_backup metric. Error in retrieving data from barman integration: %v", err)
		return
	}
	for _, server := range lastBackups {
		serverAttribute := metric.WithAttributes(attribute.String("server", string(server.Server)))
		t := time.Unix(server.Show.EndTimeTimestamp, 0)
		diff := time.Since(t)
		blb.metric.Record(ctx, (diff.Milliseconds() / int64(1000)), serverAttribute)
	}
}
