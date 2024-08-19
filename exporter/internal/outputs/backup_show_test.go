package outputs

import (
	"os"
	"testing"
)

func TestUnmarshallingSingleServerSingleBackupShow(t *testing.T) {
	jsonEntry, err := os.ReadFile("testdata/backup_show/single_server_single_backup_show.json")
	if err != nil {
		t.Errorf("TERR backup_show_test_001: Failed opening json entry.\nReason: %+v", err)
		t.FailNow()
	}
	backupsShow, err := UnmarshallBackupShow(string(jsonEntry))

	if err != nil {
		t.Errorf("TERR backup_show_test_002: Failed unmarshalling backup show entries.\nReason: %+v", err)
		t.FailNow()
	}
	if backupsShow[0].Server != "pg" {
		t.Errorf("TERR backup_show_test_003: Invalid server name, expected 'pg' got: %+v", backupsShow[0].Server)
	}
	if backupsShow[0].Show.CopyTimeSeconds != 8.239298 {
		t.Errorf("TERR backup_show_test_004: Invalid server name, expected '8.239298' got: %+v", backupsShow[0].Show.CopyTimeSeconds)

	}
	if backupsShow[0].Show.WalFiles != 2 {
		t.Errorf("TERR backup_show_test_005: Invalid number of WAL files, expected '2' got: %+v", backupsShow[0].Show.CopyTimeSeconds)
	}
	if backupsShow[0].Show.WalRatePerSecond != 0.05102330047234898 { // not the greatest floating point comparison
		t.Errorf("TERR backup_show_test_006: Invalid number of WAL files, expected '2' got: %+v", backupsShow[0].Show.CopyTimeSeconds)
	}
	if backupsShow[0].Show.ThroughputBytes != 4664618.514829783 { // not the greatest floating point comparison
		t.Errorf("TERR backup_show_test_007: Invalid throughput for backup, expected '4664618.514829783' got: %+v", backupsShow[0].Show.CopyTimeSeconds)
	}
}
