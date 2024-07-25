package integration

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"log/slog"
	"math/rand/v2"
	"os/exec"
	"strings"
	"time"
)

// Uses shell to execute barman commands.
type ShellExecutor struct {
}

// Abstracts the command execution its self, the result is stdout and error.
// The error is returned if cmd.Run method retuns an error of stderr is not empty
func (sh *ShellExecutor) executeShellCommand(ctx context.Context, flags ...string) (string, error) {
	logId := generateRandomLogId(6)

	defer trackExecutionTime(logId, flags...)()

	slog.Info("Executing barman", slog.String("flags", strings.Join(flags, " ")), slog.String("logId", logId))

	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.CommandContext(ctx, "barman", flags...)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		errorMessage := fmt.Sprintf("=====Error while executing 'barman %v'=====\nGolang Error: %v\nStderr: %v\nStdout: %v\n==========", flags, err.Error(), stderr.String(), stdout.String())
		slog.Error(errorMessage, slog.String("logId", logId))
		return stdout.String(), err
	}
	if stderr.String() != "" {
		return stdout.String(), errors.New(stderr.String())
	}
	return stdout.String(), err
}

func (sh *ShellExecutor) GetAllBackups(ctx context.Context) (string, error) {
	return sh.executeShellCommand(ctx, "-f", "json", "list-backups", "all")
}

func (sh *ShellExecutor) GetAllServerChecks(ctx context.Context) (string, error) {
	result, err := sh.executeShellCommand(ctx, "-f", "json", "check", "all")
	if err != nil && strings.Contains(err.Error(), "1") { // Suppress exist status 1 for this command, as if single check fails the status is 1 yet we still want to generate metric.
		slog.Error("!!!Supressing error from 'barman check all' command, most likely there are failing checks!!!")
		return result, nil
	}
	return result, err
}

func (sh *ShellExecutor) GetShowForBackup(ctx context.Context, serverName string, backupName string) (string, error) {
	return sh.executeShellCommand(ctx, "-f", "json", "show-backup", serverName, backupName)
}

func (sh *ShellExecutor) GetAllServerStatuses(ctx context.Context) (string, error) {
	return sh.executeShellCommand(ctx, "-f", "json", "status", "all")
}

// Generates random string, which is attached to all logs from executeShellCommand function
func generateRandomLogId(length int) string {
	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)
	for i := 0; i < length; i++ {
		result[i] = chars[rand.IntN(len(chars))]
	}
	return string(result)
}

// Tracks how much time a function execution takes and prints the command with debug log level
func trackExecutionTime(logId string, barmanFlags ...string) func() {
	start := time.Now()
	return func() {
		msg := fmt.Sprintf("'barman %s' took %v\n", barmanFlags, time.Since(start))
		slog.Debug(msg, slog.String("logId", logId))

	}
}
