package metrics

import (
	"barman-exporter/internal/integration"
	"context"
	"log/slog"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
)

// Returns the total number of backups for each server, including failed ones
type BarmanTotalBackups struct {
	metric      metric.Int64Gauge
	integration integration.BarmanIntegration
}

func NewBarmanTotalBackups(meter metric.Meter, integration integration.BarmanIntegration) (BarmanMetric, error) {
	btb := &BarmanTotalBackups{integration: integration}
	err := btb.Init(meter)
	return btb, err
}

func (btb *BarmanTotalBackups) Init(meter metric.Meter) error {
	m, err := meter.Int64Gauge("barman_backups_total", metric.WithDescription("Outputs the current number of available backups for specific server"))
	if err != nil {
		slog.Warn("Failed to create barman_backups_total metric. %v", err)
		return err
	}
	btb.metric = m
	return nil
}

// integration.GetAllBackups() and then iterates through the returned objects(range server), gets the length of the Backups slice and Records it
// Finally the counter Record in the metric
func (btb *BarmanTotalBackups) Update(ctx context.Context) {
	allServersBackups, err := btb.integration.GetAllBackups(ctx)
	if err != nil {
		slog.Warn("Failed to update barman_backups_total metric. Error in retrieving data from barman integration: %v", err)
		return
	}
	for _, server := range allServersBackups {
		serverAttribute := metric.WithAttributes(attribute.String("server", string(server.Server)))
		currentBackupsCount := int64(len(server.Backups))
		btb.metric.Record(ctx, currentBackupsCount, serverAttribute)
	}
}
