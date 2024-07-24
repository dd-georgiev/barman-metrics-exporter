<!-- TOC start (generated with https://github.com/derlin/bitdowntoc) -->

- [Implemented ](#implemented)
   * [barman_backups_total](#barman_backups_total)
   * [barman_backups_failed](#barman_backups_failed)
   * [barman_up](#barman_up)
   * [barman_backup_size](#barman_backup_size)
   * [barman_backup_wal_size](#barman_backup_wal_size)
   * [barman_last_backup_copy_time](#barman_last_backup_copy_time)
   * [barman_last_backup](#barman_last_backup)
   * [barman_first_backup](#barman_first_backup)
   * [barman_metrics_update](#barman_metrics_update)
- [Planned metrics](#planned-metrics)
      + [barman_last_backup_throughput](#barman_last_backup_throughput)
      + [barman_last_backup_wal_rate_per_second](#barman_last_backup_wal_rate_per_second)
      + [barman_active_streaming_clients](#barman_active_streaming_clients)
      + [barman_streaming_client_sent_lsn_diff](#barman_streaming_client_sent_lsn_diff)
      + [barman_streaming_client_write_lsn_diff](#barman_streaming_client_write_lsn_diff)
      + [barman_streaming_client_flush_lsn_diff](#barman_streaming_client_flush_lsn_diff)
- [Metric name](#metric-name)

<!-- TOC end -->

<!-- TOC --><a name="implemented"></a>
# Implemented 
<!-- TOC --><a name="barman_backups_total"></a>
## barman_backups_total
 Metric name: `barman_backups_total`
 Metric labels: server
 Metric type: gauge
 Metric unit: scalar
 Metric description: Outputs the current number of available backups for specific server
 Metric source: `barman -f json list-backups SERVER`, count the number of entries for each SERVER in the JSON output. Irrelevent form the backup `status`.

<!-- TOC --><a name="barman_backups_failed"></a>
## barman_backups_failed
 Metric name: `barman_backups_failed`
 Metric labels: server
 Metric type: gauge
 Metric unit: scalar
 Metric description: Outputs the current number of failed backups for specific server
 Metric source: `barman -f json list-backups SERVER`, count the number of entries for each SERVER in the JSON output. The backup `status` of the metrics must be FAILED

<!-- TOC --><a name="barman_up"></a>
## barman_up
Metric name: `barman_up`
Metric labels: check, server
Metric type: guage
Metric unit: bool 1.0 - true, 0.0 false
Metric description: translates that different checks provided by barman into set of guage metrics. Provides two new labels compared to the old exporter - `backup_minimum_age` and `pg_basebackup_supports_tablespaces_mapping`
Metric source: `barman -f json check all` convert each json struct to metric, with the field key as label and 1.0 as value for `status: "OK"` and 0.0 for any other value

<!-- TOC --><a name="barman_backup_size"></a>
## barman_backup_size
 Metric name: `barman_backup_size`
 Metric labels: server, backup position in array
 Metric type: counter
 Metric unit: bytes
 Metric description: Outputs the size of a specific backup(for specific server) in bytes. Only successful backups are taken into account. Backup is counted as successful if its status is `DONE`.
 Metric source: `barman -f json list-backups SERVER`, the `size_bytes` filed from the json out

<!-- TOC --><a name="barman_backup_wal_size"></a>
## barman_backup_wal_size
 Metric name: `barman_backup_wal_size`
 Metric labels: server, backup position in array
 Metric type: guage
 Metric unit: bytes
 Metric description: Outputs the size of a the wals for specific backup(for specific server) in bytes. Only successful backups are taken into account. Backup is counted as successful if its status is `DONE`.
 Metric source: `barman -f json list-backups SERVER`, the `wal_size_bytes` filed from the json out

<!-- TOC --><a name="barman_last_backup_copy_time"></a>
## barman_last_backup_copy_time
 Metric name: `barman_last_backup_copy_time`
 Metric labels: server
 Metric type: guage
 Metric unit: seconds
 Metric description: Outputs the time it took to get the latest backup in seconds
 Metric source: `barman -f json list-backups SERVER`, sort out by `end_time_timestamp`, run `barman -f show-backup BACKUP SERVER`, take the `copy_time_seconds` field. 
<!-- TOC --><a name="barman_last_backup"></a>
## barman_last_backup
 Metric name: `barman_last_backup`
 Metric labels: server
 Metric type: guage
 Metric unit: seconds
 Metric description: Time since the last backup was taken
 Metric source: `barman -f json status all`, take the `last_backup` field for each server
<!-- TOC --><a name="barman_first_backup"></a>
## barman_first_backup
 Metric name: `barman_first_backup`
 Metric labels: server
 Metric type: guage
 Metric unit: seconds
 Metric description: Time since the first backup was taken
 Metric source: `barman -f json status all`, take the `first_backup` field for each server
<!-- TOC --><a name="barman_metrics_update"></a>
## barman_metrics_update
 Metric name: `barman_metrics_update`
 Metric labels: none
 Metric type: guage
 Metric unit: seconds
 Metric description: Outputs the time it took finish the last run of `UpdateAll`
 Metric source:

<!-- TOC --><a name="planned-metrics"></a>
# Planned metrics

<!-- TOC --><a name="barman_last_backup_throughput"></a>
### barman_last_backup_throughput
 Metric name: `barman_last_backup_throughput`
 Metric labels: server
 Metric type: guage
 Metric unit: bytes/seconds
 Metric description: outputs the throughput during the last backup creation
 Metric source: `barman -f json list-backups SERVER`, sort out by `end_time_timestamp`, run `barman -f show-backup BACKUP SERVER`, take the `throughput_bytes` field. 


<!-- TOC --><a name="barman_last_backup_wal_rate_per_second"></a>
### barman_last_backup_wal_rate_per_second
 Metric name: `barman_last_backup_wal_rate_per_second`
 Metric labels: server
 Metric type: guage
 Metric unit: scalar/seconds
 Metric description: outputs the throughput during the last backup creation
 Metric source: `barman -f json list-backups SERVER`, sort out by `end_time_timestamp`, run `barman -f show-backup BACKUP SERVER`, take the `wal_rate_per_second` field. 


<!-- TOC --><a name="barman_active_streaming_clients"></a>
### barman_active_streaming_clients
 Metric name: `barman_active_streaming_clients`
 Metric labels: state - can be async or sync
 Metric type: guage
 Metric unit: scalar
 Metric description: outputs the current number of streaming clients
 Metric source: `barman -f json barman replication-status all`, iterate through all the entries in the json, and increment countes based on the `current_sync_state` field.

<!-- TOC --><a name="barman_streaming_client_sent_lsn_diff"></a>
### barman_streaming_client_sent_lsn_diff
 Metric name: `barman_streaming_client_sent_lsn_diff`
 Metric labels: server
 Metric type: guage
 Metric unit: bytes
 Metric description: Outputs the difference between the current master LSN and the send LSN in bytes
 Metric source: `barman -f json barman replication-status all`, iterate through all the entries in the json and take the `sent_lsn_diff_bytes` field

<!-- TOC --><a name="barman_streaming_client_write_lsn_diff"></a>
### barman_streaming_client_write_lsn_diff
 Metric name: `barman_streaming_client_write_lsn_diff`
 Metric labels: server
 Metric type: guage
 Metric unit: bytes
 Metric description: Outputs the difference between the current master LSN and the written LSN in bytes
 Metric source: `barman -f json barman replication-status all`, iterate through all the entries in the json and take the `write_lsn_diff_bytes` field

<!-- TOC --><a name="barman_streaming_client_flush_lsn_diff"></a>
### barman_streaming_client_flush_lsn_diff
 Metric name: `barman_streaming_client_flush_lsn_diff`
 Metric labels: server
 Metric type: guage
 Metric unit: bytes
 Metric description: Outputs the difference between the current master LSN and the flushed LSN in bytes
 Metric source: `barman -f json barman replication-status all`, iterate through all the entries in the json and take the `flush_lsn_diff_bytes` field


 # Template
<!-- TOC --><a name="metric-name"></a>
# Metric name
 Metric name
 Metric type
 Metric description
 Metric source
