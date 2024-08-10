package outputs

import (
	"encoding/json"
	"strconv"
)

type BackupShow struct {
	CopyTimeSeconds  float64
	EndTimeTimestamp int64
	WalFiles         int64
	ThroughputBytes  float64
	WalRatePerSecond float64
}

type ServerBackupShow struct {
	Server ServerName
	Show   BackupShow
}

func UnmarshallBackupShow(jsonInput string) ([]ServerBackupShow, error) {
	type showBackup map[string]jsonKeyInterfacePair
	var showBackupMap showBackup
	jsonBytes := []byte(jsonInput)
	err := json.Unmarshal(jsonBytes, &showBackupMap)
	if err != nil {
		return []ServerBackupShow{}, err
	}
	backupsShow := make([]ServerBackupShow, 0)
	for srvName, srvShow := range showBackupMap {
		baseBackupInfoMap := srvShow["base_backup_information"].(map[string]interface{})
		walInfoMap := srvShow["wal_information"].(map[string]interface{})
		backupEndTime, err := strconv.Atoi(baseBackupInfoMap["end_time_timestamp"].(string))
		if err != nil {
			return []ServerBackupShow{}, err
		}
		show := ServerBackupShow{
			Server: ServerName(srvName),
			Show: BackupShow{
				CopyTimeSeconds:  baseBackupInfoMap["copy_time_seconds"].(float64),
				ThroughputBytes:  baseBackupInfoMap["throughput_bytes"].(float64),
				WalRatePerSecond: walInfoMap["wal_rate_per_second"].(float64),
				EndTimeTimestamp: int64(backupEndTime),
				WalFiles:         int64(walInfoMap["no_of_files"].(float64)),
			},
		}
		backupsShow = append(backupsShow, show)
	}
	return backupsShow, nil
}
