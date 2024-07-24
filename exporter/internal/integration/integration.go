// The package integration provides means to gather data from barman and output it in go struct.
// The go structs in question are in the outputs package.
// The command execution is handled by CommandExecutor. Currently, there are two - Shell and Mock executor.
// It also manages caching in very "primitive" way, basically if the value is cache for longer than
// integration.LAST_RETRIEVAL_VAR it proceeds with retrieving it from the integration
package integration

import (
	"barman-exporter/internal/outputs"
	"context"
	"log/slog"
	"time"

	"github.com/gookit/config/v2"
)

// Abstracts away the details about the command which needs to be executed to retrieve the JSON output
type CommandExecutor interface {
	GetAllBackups(ctx context.Context) (string, error)
	GetAllServerChecks(ctx context.Context) (string, error)
	GetAllServerStatuses(ctx context.Context) (string, error)
	GetShowForBackup(ctx context.Context, serverName string, backupName string) (string, error)
}

// Glues together the CommandExecutor methods and the structs in the outputs package.
// Also provides caching functionality.
type BarmanIntegration struct {
	executor CommandExecutor
	// Last retrival to compare with cache TTL
	lastRetrivalOfBackups   time.Time
	lastRetrivalOfChecks    time.Time
	lastRetrivalOfShowLast  time.Time
	lastRetrivalOfShowFirst time.Time
	lastRetrivalOfStatuses  time.Time
	// Cached values
	allBackupsCache      []outputs.ServerBackups
	allChecksCache       []outputs.ServerChecks
	latestShowBackups    []outputs.ServerBackupShow
	firstShowBackups     []outputs.ServerBackupShow
	serversStatusesCache []outputs.ServerStatus
}

func NewIntegration(executor CommandExecutor) BarmanIntegration {
	dateInThePastForInvalidCache := time.Now().Add(time.Duration(-60) * time.Second)
	return BarmanIntegration{executor: executor, lastRetrivalOfBackups: dateInThePastForInvalidCache}
}

// Extracts all backups for all servers as reported by Barman. No filtration is applied.
func (integration *BarmanIntegration) GetAllBackups(ctx context.Context) ([]outputs.ServerBackups, error) {
	cacheTTL := time.Duration(config.Int("cache_ttl")) * time.Second

	if time.Since(integration.lastRetrivalOfBackups) < cacheTTL {
		return integration.allBackupsCache, nil
	}
	slog.Info("Retrieving all backups")
	jsonOutputs, err := integration.executor.GetAllBackups(ctx)

	if err != nil {
		return []outputs.ServerBackups{}, err
	}
	integration.allBackupsCache, err = outputs.UnmarshalServersBackups(jsonOutputs)

	if err != nil {
		return []outputs.ServerBackups{}, err
	}

	integration.lastRetrivalOfBackups = time.Now()

	return integration.allBackupsCache, nil
}

// Extacts all checks for all servers
func (integration *BarmanIntegration) GetAllServerChecks(ctx context.Context) ([]outputs.ServerChecks, error) {
	cacheTTL := time.Duration(config.Int("checks_cache_ttl")) * time.Second

	if time.Since(integration.lastRetrivalOfChecks) < cacheTTL {
		return integration.allChecksCache, nil
	}
	slog.Info("Retrieving all server checks")

	jsonOutputs, err := integration.executor.GetAllServerChecks(ctx)

	if err != nil {
		return []outputs.ServerChecks{}, err
	}

	integration.allChecksCache, err = outputs.UnmarshallServerCheck(jsonOutputs)

	if err != nil {
		return []outputs.ServerChecks{}, err
	}

	integration.lastRetrivalOfChecks = time.Now()

	return integration.allChecksCache, nil
}

// Gets the latest backups as indicated by `barman status all` command and then runs `barman show` for each of the backups
// NOTE: The shows is first stored in temporary method, because error may happen during its creation. Do not optimize this method by working directly on integration.latestShowBackups variable
func (integration *BarmanIntegration) GetShowForLatestBackupForEachServer(ctx context.Context) ([]outputs.ServerBackupShow, error) {
	cacheTTL := time.Duration(config.Int("cache_ttl")) * time.Second

	if time.Since(integration.lastRetrivalOfShowLast) < cacheTTL {
		return integration.latestShowBackups, nil
	}
	slog.Info("Retrieving the data about latest backup from 'barman status all'")

	latestBackups, err := integration.getLatestBackupsNames(ctx)
	if err != nil {
		return []outputs.ServerBackupShow{}, err
	}

	var shows []outputs.ServerBackupShow = make([]outputs.ServerBackupShow, 0)

	for srvName, backupId := range latestBackups {
		if backupId == "None" {
			continue
		}
		json, err := integration.executor.GetShowForBackup(ctx, string(srvName), backupId)
		if err != nil {
			return []outputs.ServerBackupShow{}, err
		}
		show, err := outputs.UnmarshallBackupShow(json)
		if err != nil {
			return []outputs.ServerBackupShow{}, err
		}
		shows = append(shows, show[0])
	}

	integration.latestShowBackups = shows
	integration.lastRetrivalOfShowLast = time.Now()
	return integration.latestShowBackups, nil
}

// Gets the first backups as indicated by `barman status all` command and then runs `barman show` for each of the backups
func (integration *BarmanIntegration) GetShowForFirstBackupForEachServer(ctx context.Context) ([]outputs.ServerBackupShow, error) {
	cacheTTL := time.Duration(config.Int("cache_ttl")) * time.Second

	if time.Since(integration.lastRetrivalOfShowFirst) < cacheTTL {
		return integration.firstShowBackups, nil
	}
	slog.Info("Retrieving data about the first backup from 'barman status all'")

	firstBackups, err := integration.getFirstBackupsNames(ctx)
	if err != nil {
		return []outputs.ServerBackupShow{}, err
	}

	var shows []outputs.ServerBackupShow = make([]outputs.ServerBackupShow, 0)

	for srvName, backupId := range firstBackups {
		if backupId == "None" {
			continue
		}
		json, err := integration.executor.GetShowForBackup(ctx, string(srvName), backupId)
		if err != nil {
			return []outputs.ServerBackupShow{}, err
		}
		show, err := outputs.UnmarshallBackupShow(json)
		if err != nil {
			return []outputs.ServerBackupShow{}, err
		}
		shows = append(shows, show[0])
	}

	integration.firstShowBackups = shows
	integration.lastRetrivalOfShowFirst = time.Now()

	return integration.firstShowBackups, nil
}

// Gets the output for `barman server status all`
func (integration *BarmanIntegration) GetAllServerStatuses(ctx context.Context) ([]outputs.ServerStatus, error) {
	cacheTTL := time.Duration(config.Int("cache_ttl")) * time.Second
	if time.Since(integration.lastRetrivalOfStatuses) < cacheTTL {
		return integration.serversStatusesCache, nil
	}
	slog.Info("Retrieving all server statuses")

	jsonOutputs, err := integration.executor.GetAllServerStatuses(ctx)

	if err != nil {
		return []outputs.ServerStatus{}, err
	}

	integration.lastRetrivalOfStatuses = time.Now()
	serverStatuses, err := outputs.UnmarshalServerStatus(jsonOutputs)

	if err != nil {
		return []outputs.ServerStatus{}, err
	}
	integration.lastRetrivalOfStatuses = time.Now()
	integration.serversStatusesCache = serverStatuses
	return integration.serversStatusesCache, nil
}

// Auxiliry method which returns the BACKUP_NAME for the last backup as reported by barman status command
func (integration *BarmanIntegration) getLatestBackupsNames(ctx context.Context) (map[outputs.ServerName]string, error) {
	var latestBackups = make(map[outputs.ServerName]string)
	statuses, err := integration.GetAllServerStatuses(ctx)
	if err != nil {
		return map[outputs.ServerName]string{}, err
	}
	for _, srvStatus := range statuses {
		latestBackups[srvStatus.Server] = srvStatus.Status.LastBackup
	}
	return latestBackups, nil
}

// Auxiliry method which returns the BACKUP_NAME for the first backup as reported by barman status command
func (integration *BarmanIntegration) getFirstBackupsNames(ctx context.Context) (map[outputs.ServerName]string, error) {
	var firstBackups = make(map[outputs.ServerName]string)
	statuses, err := integration.GetAllServerStatuses(ctx)
	if err != nil {
		return map[outputs.ServerName]string{}, err
	}
	for _, srvStatus := range statuses {
		firstBackups[srvStatus.Server] = srvStatus.Status.FirstBackup
	}
	return firstBackups, nil
}
