package outputs

import (
	"os"
	"testing"
)

func TestUnmarshallingSingleServerSingleBackupsStatus(t *testing.T) {
	jsonEntry, err := os.ReadFile("testdata/server_status/singleserver/one_backup.json")
	if err != nil {
		t.Errorf("TERR server_status_0001: Failed opening json entry.\nReason: %+v", err)
		t.FailNow()
	}
	status, err := UnmarshalServerStatus(string(jsonEntry))
	if err != nil {
		t.Errorf("TERR server_status_0002: Failed unmarshalling servers status entry.\nReason: %+v", err)
		t.FailNow()
	}
	if len(status) != 1 {
		t.Errorf("TERR server_status_0002: Expected status len to be 1, got %d", len(status))
	}
	if status[0].Server != "pg" {
		t.Errorf("TERR server_status_0003: Invalid server name, expected 'pg' got: %+v", status[0].Server)
	}
	if status[0].Status.LastBackup != "20240707T183345" {
		t.Errorf("TERR server_status_0004: Invalid last backup id, expected '20240707T183345' got: %+v", status[0].Status.LastBackup)
	}
	if status[0].Status.FirstBackup != "20240707T183345" {
		t.Errorf("TERR server_status_0005: Invalid first backup id, expected '20240707T183345' got: %+v", status[0].Status.FirstBackup)
	}
}

func TestUnmarshallingSingleServerTwoBackupsStatus(t *testing.T) {
	jsonEntry, err := os.ReadFile("testdata/server_status/singleserver/two_backups.json")
	if err != nil {
		t.Errorf("TERR server_status_0006: Failed opening json entry.\nReason: %+v", err)
		t.FailNow()
	}
	status, err := UnmarshalServerStatus(string(jsonEntry))
	if err != nil {
		t.Errorf("TERR server_status_0007: Failed unmarshalling servers status entry.\nReason: %+v", err)
		t.FailNow()
	}
	if len(status) != 1 {
		t.Errorf("TERR server_status_0008: Expected status len to be 1, got %d", len(status))
	}
	if status[0].Server != "pg" {
		t.Errorf("TERR server_status_0009: Invalid server name, expected 'pg' got: %+v", status[0].Server)
	}
	if status[0].Status.LastBackup != "20240707T183345" {
		t.Errorf("TERR server_status_0010: Invalid last backup id, expected '20240707T183345' got: %+v", status[0].Status.LastBackup)
	}
	if status[0].Status.FirstBackup != "20240707T183151" {
		t.Errorf("TERR server_status_0011: Invalid first backup id, expected '20240707T183151' got: %+v", status[0].Status.FirstBackup)
	}
}

func TestUnmarshallingTwoServersTwoBackupsStatus(t *testing.T) {
	jsonEntry, err := os.ReadFile("testdata/server_status/multiserver/two_backups.json")
	if err != nil {
		t.Errorf("TERR server_status_0012: Failed opening json entry.\nReason: %+v", err)
		t.FailNow()
	}
	status, err := UnmarshalServerStatus(string(jsonEntry))
	if err != nil {
		t.Errorf("TERR server_status_0013: Failed unmarshalling servers status entry.\nReason: %+v", err)
		t.FailNow()
	}
	if len(status) != 2 {
		t.Errorf("TERR server_status_0014: Expected status len to be 2, got %d", len(status))
	}
	if status[0].Server != "pg-0" && status[0].Server != "pg-1" {
		t.Errorf("TERR server_status_0015: Invalid server name, expected 'pg-0' or 'pg-1' got: %+v", status[0].Server)
	}
	if status[1].Server != "pg-0" && status[1].Server != "pg-1" {
		t.Errorf("TERR server_status_0016: Invalid server name, expected 'pg-0' or 'pg-1' got: %+v", status[0].Server)
	}
	if status[1].Server == status[0].Server {
		t.Errorf("TERR server_status_0017: Invalid server name, expected the two servers to have different name, got single name: %+v", status[0].Server)
	}
	var pg_0, pg_1 ServerStatus
	if status[0].Server == "pg-0" {
		pg_0 = status[0]
		pg_1 = status[1]
	} else {
		pg_0 = status[1]
		pg_1 = status[0]
	}

	if pg_0.Status.LastBackup != "20240707T183345" {
		t.Errorf("TERR server_status_0010: Invalid last backup id, expected '20240707T183345' got: %+v", pg_0.Status.LastBackup)
	}
	if pg_0.Status.FirstBackup != "20240707T183151" {
		t.Errorf("TERR server_status_0011: Invalid first backup id, expected '20240707T183151' got: %+v", pg_0.Status.FirstBackup)
	}
	if pg_1.Status.LastBackup != "20240707T183351" {
		t.Errorf("TERR server_status_0020: Invalid last backup id, expected '20240707T183351' got: %+v", pg_1.Status.LastBackup)
	}
	if pg_1.Status.FirstBackup != "20240707T183156" {
		t.Errorf("TERR server_status_0021: Invalid first backup id, expected '20240707T183156' got: %+v", pg_1.Status.FirstBackup)
	}
}
