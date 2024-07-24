package outputs

import (
	"os"
	"testing"
)

func TestUnmarshallingSingleServerCheckWithAllSuccessful(t *testing.T) {
	jsonEntry, err := os.ReadFile("testdata/server_checks/single_server/successful_checks.json")
	if err != nil {
		t.Errorf("TERR server_checks_test_0001: Failed opening json entry.\nReason: %+v", err)
		t.FailNow()
	}
	checks, err := UnmarshallServerCheck(string(jsonEntry))
	if err != nil {
		t.Errorf("TERR server_checks_test_0002: Failed unmarshalling servers check entry.\nReason: %+v", err)
		t.FailNow()
	}
	if checks[0].Server != "pg" {
		t.Errorf("TERR server_checks_test_0003: Invalid server name, expected 'pg' got: %+v", checks[0].Server)
	}
	for _, key := range ServerChecksKeys {
		if checks[0].Check[key] != true {
			t.Errorf("TERR server_checks_test_0004: Invalid %s, expected 'true' got: %+v", key, checks[0].Check[key])
		}
	}
}
func TestUnmarshallingSingleServerCheckWithAllFailing(t *testing.T) {
	jsonEntry, err := os.ReadFile("testdata/server_checks/single_server/failing_checks.json")
	if err != nil {
		t.Errorf("TERR server_checks_test_0005: Failed opening json entry.\nReason: %+v", err)
		t.FailNow()
	}
	checks, err := UnmarshallServerCheck(string(jsonEntry))
	if err != nil {
		t.Errorf("TERR server_checks_test_0006: Failed unmarshalling servers check entry.\nReason: %+v", err)
		t.FailNow()
	}
	if checks[0].Server != "pg" {
		t.Errorf("TERR server_checks_test_0007: Invalid server name, expected 'pg' got: %+v", checks[0].Server)
	}
	for _, key := range ServerChecksKeys {
		if checks[0].Check[key] != false {
			t.Errorf("TERR server_checks_test_0008: Invalid %s, expected 'false' got: %+v", key, checks[0].Check[key])
		}
	}
}

func TestUnmarshallingWithMissingCheckMustSetTheCheckToFalse(t *testing.T) {
	jsonEntry, err := os.ReadFile("testdata/server_checks/single_server/missing_check.json")
	if err != nil {
		t.Errorf("TERR server_checks_test_0009: Failed opening json entry.\nReason: %+v", err)
		t.FailNow()
	}
	checks, err := UnmarshallServerCheck(string(jsonEntry))
	if err != nil {
		t.Errorf("TERR server_checks_test_0010: Failed unmarshalling servers check entry.\nReason: %+v", err)
		t.FailNow()
	}
	if checks[0].Check["archiver_errors"] == true {
		t.Errorf("TERR server_checks_test_00011: Invalid archiver_errors, expected 'false' got: %+v", checks[0].Check["archiver_errors"])
	}
	for i, key := range ServerChecksKeys {
		if i == 0 {
			continue
		}
		if checks[0].Check[key] != true {
			t.Errorf("TERR server_checks_test_0012: Invalid %s, expected 'true' got: %+v", key, checks[0].Check[key])
		}
	}
}
func TestUnmarshallingSingleServerCheckWithAllChecksMissing(t *testing.T) {
	jsonEntry, err := os.ReadFile("testdata/server_checks/single_server/missing_all_checks.json")
	if err != nil {
		t.Errorf("TERR server_checks_test_0013: Failed opening json entry.\nReason: %+v", err)
		t.FailNow()
	}
	checks, err := UnmarshallServerCheck(string(jsonEntry))
	if err != nil {
		t.Errorf("TERR server_checks_test_0014: Failed unmarshalling servers check entry.\nReason: %+v", err)
		t.FailNow()
	}
	if checks[0].Server != "pg" {
		t.Errorf("TERR server_checks_test_0015: Invalid server name, expected 'pg' got: %+v", checks[0].Server)
	}
	for _, key := range ServerChecksKeys {
		if checks[0].Check[key] != false {
			t.Errorf("TERR server_checks_test_0016: Invalid %s, expected 'true' got: %+v", key, checks[0].Check[key])
		}
	}
}

func TestUnmarshallingTwoServersCheckWithAllSuccessful(t *testing.T) {
	jsonEntry, err := os.ReadFile("testdata/server_checks/multi_server/successful_checks.json")
	if err != nil {
		t.Errorf("TERR server_checks_test_0017: Failed opening json entry.\nReason: %+v", err)
		t.FailNow()
	}
	checks, err := UnmarshallServerCheck(string(jsonEntry))
	if err != nil {
		t.Errorf("TERR server_checks_test_0018: Failed unmarshalling servers check entry.\nReason: %+v", err)
		t.FailNow()
	}
	if checks[0].Server != "pg-0" && checks[0].Server != "pg-1" {
		t.Errorf("TERR server_checks_test_0019: Invalid server name, expected 'pg-0' or 'pg-1' got: %+v", checks[0].Server)
	}
	if checks[1].Server != "pg-0" && checks[1].Server != "pg-1" {
		t.Errorf("TERR server_checks_test_0020: Invalid server name, expected 'pg-0' or 'pg-1' got: %+v", checks[0].Server)
	}
	if checks[0].Server == checks[1].Server {
		t.Errorf("TERR server_checks_test_0028: Expected servers to have different name got: %+v", checks[1].Server)
	}
	for _, key := range ServerChecksKeys {
		if checks[0].Check[key] != true {
			t.Errorf("TERR server_checks_test_0021: Invalid %s for server %s, expected 'true' got: %+v", key, checks[0].Server, checks[0].Check[key])
		}
		if checks[1].Check[key] != true {
			t.Errorf("TERR server_checks_test_0022: Invalid %s for server %s, expected 'true' got: %+v", key, checks[1].Server, checks[1].Check[key])
		}

	}
}

func TestUnmarshallingTwoServersCheckWithAllFailing(t *testing.T) {
	jsonEntry, err := os.ReadFile("testdata/server_checks/multi_server/failing_checks.json")
	if err != nil {
		t.Errorf("TERR server_checks_test_0023: Failed opening json entry.\nReason: %+v", err)
		t.FailNow()
	}
	checks, err := UnmarshallServerCheck(string(jsonEntry))
	if err != nil {
		t.Errorf("TERR server_checks_test_0024: Failed unmarshalling servers check entry.\nReason: %+v", err)
		t.FailNow()
	}
	if checks[0].Server != "pg-0" && checks[0].Server != "pg-1" {
		t.Errorf("TERR server_checks_test_0025: Invalid server name, expected 'pg-0' or 'pg-1' got: %+v", checks[0].Server)
	}
	if checks[1].Server != "pg-0" && checks[1].Server != "pg-1" {
		t.Errorf("TERR server_checks_test_0025: Invalid server name, expected 'pg-0' or 'pg-1' got: %+v", checks[1].Server)
	}
	if checks[0].Server == checks[1].Server {
		t.Errorf("TERR server_checks_test_0028: Expected servers to have different name got: %+v", checks[1].Server)
	}
	for _, key := range ServerChecksKeys {
		if checks[0].Check[key] != false {
			t.Errorf("TERR server_checks_test_0026: Invalid %s for server %s, expected 'false' got: %+v", key, checks[0].Server, checks[0].Check[key])
		}
		if checks[0].Check[key] != false {
			t.Errorf("TERR server_checks_test_0027: Invalid %s for server %s, expected 'false' got: %+v", key, checks[1].Server, checks[1].Check[key])
		}
	}
}
