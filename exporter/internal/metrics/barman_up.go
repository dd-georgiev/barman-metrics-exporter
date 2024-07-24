package metrics

import (
	"barman-exporter/internal/integration"
	"barman-exporter/internal/outputs"
	"context"
	"log/slog"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
)

// Outputs the current number of available backups for specific server
type BarmanUp struct {
	metric      metric.Int64Gauge
	integration integration.BarmanIntegration
}

func NewBarmanUp(meter metric.Meter, integration integration.BarmanIntegration) (BarmanMetric, error) {
	bu := &BarmanUp{integration: integration}
	err := bu.Init(meter)
	return bu, err
}

func (bu *BarmanUp) Init(meter metric.Meter) error {
	m, err := meter.Int64Gauge("barman_up", metric.WithDescription("Barman status checks"))
	if err != nil {
		slog.Warn("Failed to create barman_up metric. %v", err)
		return err
	}
	bu.metric = m
	return nil
}

// integration.GetAllServerChecks() and then iterates through the returned objects(range server), checks every bool in the map and sets it to 1 for true and 0 for any other value
// Finally the counter Record in the metric
func (bu *BarmanUp) Update(ctx context.Context) {
	allServersCheck, err := bu.integration.GetAllServerChecks(ctx)
	if err != nil {
		slog.Warn("Failed to update barman_up metric. Error in retrieving data from barman integration: %v", err)
		return
	}
	for _, server := range allServersCheck {
		serverAttribute := metric.WithAttributes(attribute.String("server", string(server.Server)))
		for _, checkName := range outputs.ServerChecksKeys {
			checkAttribute := metric.WithAttributes(attribute.String("check", checkName))
			if server.Check[checkName] {
				bu.metric.Record(ctx, 1, serverAttribute, checkAttribute)
			} else {
				bu.metric.Record(ctx, 0, serverAttribute, checkAttribute)
			}
		}

	}
}
