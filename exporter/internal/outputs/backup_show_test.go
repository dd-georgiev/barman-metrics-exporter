package outputs

import (
	"os"
	"testing"
)

func TestUnmarshallingSingleServerSingleBackupShow(t *testing.T) {
	jsonEntry, err := os.ReadFile("testdata/single_server_single_backup_show.json")
	if err != nil {
		t.Errorf("TERR outputs_0015: Failed opening json entry.\nReason: %+v", err)
		t.FailNow()
	}
	backupsShow, err := UnmarshallBackupShow(string(jsonEntry))

	if err != nil {
		t.Errorf("TERR outputs_0016: Failed unmarshalling backup show entries.\nReason: %+v", err)
		t.FailNow()
	}
	if backupsShow[0].Server != "pg" {
		t.Errorf("TERR outputs_0017: Invalid server name, expected 'pg' got: %+v", backupsShow[0].Server)
	}
	if backupsShow[0].Show.CopyTimeSeconds != 8.239298 {
		t.Errorf("TERR outputs_0018: Invalid server name, expected '8.239298' got: %+v", backupsShow[0].Show.CopyTimeSeconds)
	}
}
