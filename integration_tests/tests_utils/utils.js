const ALL_METRICS_NAMES = [
    "barman_up",
    "barman_last_backup_copy_time",
    "barman_first_backup",
    "barman_last_backup",
    "barman_backups_total",
    "barman_backups_failed",
    "barman_backup_wal_size",
    "barman_backup_size",
    "barman_last_backup_wal_rate_per_second",
    "barman_last_backup_wal_files",
    "barman_last_backup_throughput"
]
const BARMAN_UP_LABELS = [
    "archiver_errors",
    "backup_maximum_age",
    "compression_settings",
    "directories",
    "failed_backups",
    "minimum_redundancy_requirements",
    "pg_basebackup",
    "pg_basebackup_compatible",
    "pg_basebackup_supports_tablespaces_mapping",
    "pg_receivexlog",
    "pg_receivexlog_compatible",
    "postgresql",
    "postgresql_streaming",
    "receive_wal_running",
    "replication_slot",
    "retention_policy_settings",
    "superuser_or_standard_user_with_backup_privileges",
    "systemid_coherence",
    "wal_level",
    "wal_maximum_age",
    "wal_size"
]

module.exports = {
    ALL_METRICS_NAMES, BARMAN_UP_LABELS
}