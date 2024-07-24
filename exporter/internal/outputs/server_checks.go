package outputs

import (
	"encoding/json"
	"strings"
)

// List with json keys for different properties of barman check command output.
var ServerChecksKeys = [...]string{"archiver_errors", "backup_maximum_age", "compression_settings", "directories", "failed_backups", "minimum_redundancy_requirements", "pg_basebackup", "pg_basebackup_compatible", "pg_basebackup_supports_tablespaces_mapping", "pg_receivexlog", "pg_receivexlog_compatible", "postgresql", "postgresql_streaming", "receive_wal_running", "replication_slot", "retention_policy_settings", "superuser_or_standard_user_with_backup_privileges", "systemid_coherence", "wal_level", "wal_maximum_age", "wal_size"}

type CheckEntry map[string]bool

// Deals with presenting a server check returned by `barman check` command as boolean, which later is exportered as 0 or 1.
type ServerChecks struct {
	Server ServerName
	Check  CheckEntry
}

func UnmarshallServerCheck(jsonInput string) ([]ServerChecks, error) {
	var checksMap map[string]jsonKeyObjectPair
	jsonBytes := []byte(jsonInput)
	err := json.Unmarshal(jsonBytes, &checksMap)
	if err != nil {
		return []ServerChecks{}, err
	}
	checks := make([]ServerChecks, 0)
	for srvName, srvCheck := range checksMap {
		check := newCheckEntryFromMap(srvCheck)
		bsc := ServerChecks{
			Server: ServerName(srvName),
			Check:  check,
		}
		checks = append(checks, bsc)
	}
	return checks, nil
}

func okTextToBool(input string) bool {
	return strings.ToUpper(input) == "OK"
}

func newCheckEntryFromMap(srvCheck jsonKeyObjectPair) CheckEntry {
	var entry CheckEntry = make(CheckEntry)
	for _, key := range ServerChecksKeys {
		entry[key] = okTextToBool(srvCheck[key]["status"])
	}
	return entry
}
