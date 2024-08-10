package outputs

import (
	"encoding/json"
	"os"
	"testing"
)

func TestUnmarshallingSingleServer(t *testing.T) {
	jsonEntry, err := os.ReadFile("testdata/server_backups/barman_backup_list_single_server.json")
	if err != nil {
		t.Errorf("TERR server_backups_test_001: Failed opening json entry.\nReason: %+v", err)
		t.FailNow()
	}
	backups, err := UnmarshalServersBackups(string(jsonEntry))
	if err != nil {
		t.Errorf("TERR server_backups_test_002: Failed unmarshalling servers backup entry.\nReason: %+v", err)
		t.FailNow()
	}
	if len(backups) != 1 {
		t.Errorf("TERR server_backups_test_003: Expected backups len to be 1, got %d", len(backups))
	}
	if backups[0].Server != "pg" {
		t.Errorf("TERR server_backups_test_004: Invalid server name, expected 'pg' got: %+v", backups[0].Server)
	}
	if backups[0].Backups[0].Status != "DONE" {
		t.Errorf("TERR server_backups_test_005: Invalid backup status, expected 'DONE' got: %+v", backups[0].Backups[0].Status)
	}
}

func TestUnmarshallingSingleBackupEntry(t *testing.T) {
	jsonEntry, err := os.ReadFile("testdata/server_backups/single_backup_entry.json")
	if err != nil {
		t.Errorf("TERR server_backups_test_006: Failed opening json entry.\nReason: %+v", err)
		t.FailNow()
	}
	bytes := []byte(jsonEntry)
	var testEntry BackupEntries = *new(BackupEntries)
	json.Unmarshal(bytes, &testEntry)

	if testEntry[0].Status != "DONE" {
		t.Errorf("TERR server_backups_test_007: Expected entry status to be DONE, got: %+v", testEntry[0].Status)
	}
}

func TestUnmarshallingMultipleBackupEntries(t *testing.T) {
	jsonEntry, err := os.ReadFile("testdata/server_backups/multiple_backup_entries.json")
	if err != nil {
		t.Errorf("TERR server_backups_test_008: Failed opening json entry.\nReason: %+v", err)
		t.FailNow()
	}
	bytes := []byte(jsonEntry)
	var testEntry BackupEntries = *new(BackupEntries)
	json.Unmarshal(bytes, &testEntry)

	if testEntry[0].Status != "DONE" {
		t.Errorf("TERR server_backups_test_009: Expected entry status to be DONE, got: %+v", testEntry[0].Status)
	}

	if testEntry[1].Status != "FAILED" {
		t.Errorf("TERR server_backups_test_010: Expected entry status to be FAILED, got: %+v", testEntry[1].Status)
	}
}
