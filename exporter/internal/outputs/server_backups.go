package outputs

import "encoding/json"

// Structure representing a specific server with all its backups irrelevent of thier parameters
type ServerBackups struct {
	Server  ServerName
	Backups BackupEntries
}

// Single backup entry without relationship with to the server it belonds.
// The only field ommited from the human-readable output it tablespaces.
type BackupEntries []struct {
	BackupID         string `json:"backup_id"`
	EndTime          string `json:"end_time"`
	EndTimeTimestamp string `json:"end_time_timestamp"`
	RetentionStatus  string `json:"retention_status"`
	Size             string `json:"size"`
	SizeBytes        int64  `json:"size_bytes"`
	Status           string `json:"status"`
	WalSize          string `json:"wal_size"`
	WalSizeBytes     int64  `json:"wal_size_bytes"`
}

// Converts JSON output from barman to ServerBackups and BackupEntries
// Its possible to throw error during unmarshalling, in which case the output is empty struct.
// Basically it first unmarshals the json into map of string(the server name is key in the output object) and BackupEntries
// Following the unmarshalling, the map is converted to slice of ServerBackups, using the map key as server name.
func UnmarshalServersBackups(jsonInput string) ([]ServerBackups, error) {
	var backupsMap map[string]BackupEntries
	jsonBytes := []byte(jsonInput)
	err := json.Unmarshal(jsonBytes, &backupsMap)
	if err != nil {
		return []ServerBackups{}, err
	}

	backups := make([]ServerBackups, 0)
	for srvName, srvBackupEntries := range backupsMap {
		backups = append(backups, ServerBackups{Server: ServerName(srvName), Backups: srvBackupEntries})
	}
	return backups, nil
}
