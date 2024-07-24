package outputs

import "encoding/json"

type Status struct {
	LastBackup  string
	FirstBackup string
}

type ServerStatus struct {
	Server ServerName
	Status Status
}

func UnmarshalServerStatus(jsonInput string) ([]ServerStatus, error) {
	var statusAllMap map[string]jsonKeyInterfacePair
	jsonBytes := []byte(jsonInput)
	err := json.Unmarshal(jsonBytes, &statusAllMap)
	if err != nil {
		return []ServerStatus{}, err
	}
	statuses := make([]ServerStatus, 0)
	for srvName, srvStatus := range statusAllMap {

		status := ServerStatus{
			Server: ServerName(srvName),
			Status: Status{
				LastBackup:  extractStringFromJsonMessageKey(srvStatus, "last_backup"),
				FirstBackup: extractStringFromJsonMessageKey(srvStatus, "first_backup"),
			},
		}

		statuses = append(statuses, status)
	}
	return statuses, nil
}

func extractStringFromJsonMessageKey(val jsonKeyInterfacePair, key string) string {
	return val[key].(map[string]interface{})["message"].(string)
}
