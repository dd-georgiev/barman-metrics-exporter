package outputs

import (
	"encoding/json"
	"strconv"
)

type BackupShow struct {
	CopyTimeSeconds  float64
	EndTimeTimestamp int64
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
		backupEndTime, err := strconv.Atoi(baseBackupInfoMap["end_time_timestamp"].(string))
		if err != nil {
			return []ServerBackupShow{}, err
		}
		show := ServerBackupShow{
			Server: ServerName(srvName),
			Show: BackupShow{
				CopyTimeSeconds:  baseBackupInfoMap["copy_time_seconds"].(float64),
				EndTimeTimestamp: int64(backupEndTime),
			},
		}
		backupsShow = append(backupsShow, show)
	}
	return backupsShow, nil
}
