package integration

import (
	"barman-exporter/internal/outputs"
	"context"
	"testing"
)

func TestGetAllBackups(t *testing.T) {
	ctx := context.Background()

	mockExecutor := MockExecutor{}
	integration := NewIntegration(&mockExecutor)
	backups, err := integration.GetAllBackups(ctx)
	if err != nil {
		t.Errorf("TERR integration_0001 Failed to get all backups. Reason: %s", err)
	}
	if len(backups) != 1 {
		t.Errorf("TERR integration_0002: Expected backups len to be 1, got %d", len(backups))
	}
	if backups[0].Server != "pg" {
		t.Errorf("TERR integration_0003: Invalid server name, expected 'pg' got: %+v", backups[0].Server)
	}
	if backups[0].Backups[0].Status != "DONE" {
		t.Errorf("TERR integration_0004: Invalid backup status, expected 'DONE' got: %+v", backups[0].Backups[0].Status)
	}
}

func TestGetAllServerChecks(t *testing.T) {
	ctx := context.Background()

	mockExecutor := MockExecutor{}
	integration := NewIntegration(&mockExecutor)

	checks, err := integration.GetAllServerChecks(ctx)
	if err != nil {
		t.Errorf("TERR integration_0005 Failed to get all backups. Reason: %s", err)
	}
	if len(checks) != 1 {
		t.Errorf("TERR integration_0006: Expected backups len to be 1, got %d", len(checks))
	}
	if checks[0].Server != "pg" {
		t.Errorf("TERR integration_0007: Invalid server name, expected 'pg' got: %+v", checks[0].Server)
	}
	for _, key := range outputs.ServerChecksKeys {
		if checks[0].Check[key] != true {
			t.Errorf("TERR integration_0015: Invalid %s, expected 'true' got: %+v", key, checks[0].Check[key])
		}
	}
}

func TestGetShowForLatestBackupForEachServer(t *testing.T) {
	ctx := context.Background()

	mockExecutor := MockExecutor{}
	integration := NewIntegration(&mockExecutor)
	shows, err := integration.GetShowForLatestBackupForEachServer(ctx)
	if err != nil {
		t.Errorf("TERR integration_0008 Failed to get all backups. Reason: %s", err)
	}
	if shows[0].Server != "pg" {
		t.Errorf("TERR integration_0009: Invalid server name, expected 'pg' got: %v", shows[0].Server)
	}
	if shows[0].Show.CopyTimeSeconds != 8.239298 {
		t.Errorf("TERR integration_0010: Invalid backup copy time seconds, expected '8.239298' got: %+v", shows[0].Show.CopyTimeSeconds)
	}
}

func TestGetStatuses(t *testing.T) {
	ctx := context.Background()

	mockExecutor := MockExecutor{}
	integration := NewIntegration(&mockExecutor)
	shows, err := integration.GetAllServerStatuses(ctx)
	if err != nil {
		t.Errorf("TERR integration_0011 Failed to get all backups. Reason: %s", err)
	}
	if shows[0].Server != "pg" {
		t.Errorf("TERR integration_0012: Invalid server name, expected 'pg' got: %v", shows[0].Server)
	}
	if shows[0].Status.LastBackup != "20240622T084702" {
		t.Errorf("TERR integration_0013: Invalid last backup: %+v", shows[0].Status.LastBackup)
	}
	if shows[0].Status.FirstBackup != "20240622T084702" {
		t.Errorf("TERR integration_0014: Invalid first backups %+v", shows[0].Status.FirstBackup)
	}
}
