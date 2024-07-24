# Version 3.10.0
## List servers
Command: `barman -f json list-servers`  
Output:
```json
{
    "pg": {
        "description": ""
    }
}
```
## Server check
Command: `barman -f json check all`
Output: 
```json
{
  "pg": {
    "archiver_errors": {
      "hint": "",
      "status": "OK"
    },
    "backup_maximum_age": {
      "hint": "no last_backup_maximum_age provided",
      "status": "OK"
    },
    "backup_minimum_size": {
      "hint": "36.7 MiB",
      "status": "OK"
    },
    "compression_settings": {
      "hint": "",
      "status": "OK"
    },
    "directories": {
      "hint": "",
      "status": "OK"
    },
    "failed_backups": {
      "hint": "there are 0 failed backups",
      "status": "OK"
    },
    "minimum_redundancy_requirements": {
      "hint": "have 1 backups, expected at least 0",
      "status": "OK"
    },
    "pg_basebackup": {
      "hint": "",
      "status": "OK"
    },
    "pg_basebackup_compatible": {
      "hint": "",
      "status": "OK"
    },
    "pg_basebackup_supports_tablespaces_mapping": {
      "hint": "",
      "status": "OK"
    },
    "pg_receivexlog": {
      "hint": "",
      "status": "OK"
    },
    "pg_receivexlog_compatible": {
      "hint": "",
      "status": "OK"
    },
    "postgresql": {
      "hint": "",
      "status": "OK"
    },
    "postgresql_streaming": {
      "hint": "",
      "status": "OK"
    },
    "receive_wal_running": {
      "hint": "",
      "status": "OK"
    },
    "replication_slot": {
      "hint": "",
      "status": "OK"
    },
    "retention_policy_settings": {
      "hint": "",
      "status": "OK"
    },
    "superuser_or_standard_user_with_backup_privileges": {
      "hint": "",
      "status": "OK"
    },
    "systemid_coherence": {
      "hint": "",
      "status": "OK"
    },
    "wal_level": {
      "hint": "",
      "status": "OK"
    },
    "wal_maximum_age": {
      "hint": "no last_wal_maximum_age provided",
      "status": "OK"
    },
    "wal_size": {
      "hint": "32.0 MiB",
      "status": "OK"
    }
  }
}
```

## Replication status check 
Command: `barman -f json replication-status pg`  
Output:
```json
{
    "last_position": 116,
    "last_name": "000000010000000000000005",
    "backups": {
        "20240622T084702": {
            "backup_label": "'START WAL LOCATION: 0/3000028 (file 000000010000000000000003)\\nCHECKPOINT LOCATION: 0/3000060\\nBACKUP METHOD: streamed\\nBACKUP FROM: primary\\nSTART TIME: 2024-06-22 08:47:06 UTC\\nLABEL: pg_basebackup base backup\\nSTART TIMELINE: 1\\n'",
            "begin_offset": 40,
            "begin_time": "Sat Jun 22 08:47:02 2024",
            "begin_wal": "000000010000000000000003",
            "begin_xlog": "0/3000028",
            "compression": null,
            "config_file": "/etc/postgresql/postgresql.conf",
            "copy_stats": {
                "copy_time": 8.239298,
                "total_time": 8.239298
            },
            "deduplicated_size": 38433182,
            "end_offset": 0,
            "end_time": "Sat Jun 22 08:47:10 2024",
            "end_wal": "000000010000000000000003",
            "end_xlog": "0/4000000",
            "error": null,
            "hba_file": "/etc/postgresql/pg_hba.conf",
            "ident_file": "/var/lib/postgresql/data/pg_ident.conf",
            "included_files": null,
            "mode": "postgres",
            "pgdata": "/var/lib/postgresql/data",
            "server_name": "pg",
            "size": 38433182,
            "status": "DONE",
            "systemid": "7383246261373567014",
            "tablespaces": null,
            "timeline": 1,
            "version": 160003,
            "xlog_segment_size": 16777216,
            "backup_id": "20240622T084702"
        }
    },
    "wals": [
        {
            "compression": null,
            "name": "000000010000000000000003",
            "size": 16777216,
            "time": 1719046026.8233335
        },
        {
            "compression": null,
            "name": "000000010000000000000004",
            "size": 16777216,
            "time": 1719046065.0100002
        },
        {
            "compression": null,
            "name": "000000010000000000000005",
            "size": 16777216,
            "time": 1719046085.6200004
        }
    ],
    "version": "3.10.0",
    "config": {
        "msg_list": [],
        "name": "pg",
        "barman_home": "/var/lib/barman",
        "barman_lock_directory": "/var/lib/barman",
        "lock_directory_cleanup": true,
        "config_changes_queue": "/var/lib/barman/cfg_changes.queue",
        "active": true,
        "archiver": false,
        "archiver_batch_size": 0,
        "autogenerate_manifest": false,
        "aws_profile": null,
        "aws_region": null,
        "azure_credential": null,
        "azure_resource_group": null,
        "azure_subscription_id": null,
        "backup_compression": null,
        "backup_compression_format": null,
        "backup_compression_level": null,
        "backup_compression_location": null,
        "backup_compression_workers": null,
        "backup_directory": "/var/lib/barman/pg",
        "backup_method": "postgres",
        "backup_options": "concurrent_backup",
        "bandwidth_limit": null,
        "basebackup_retry_sleep": 30,
        "basebackup_retry_times": 0,
        "basebackups_directory": "/var/lib/barman/pg/base",
        "check_timeout": 30,
        "cluster": "pg",
        "compression": null,
        "conninfo": "host=database user=tester password=tester dbname=postgres application_name=barman",
        "custom_compression_filter": null,
        "custom_decompression_filter": null,
        "custom_compression_magic": null,
        "description": null,
        "disabled": false,
        "errors_directory": "/var/lib/barman/pg/errors",
        "forward_config_path": false,
        "gcp_project": null,
        "gcp_zone": null,
        "immediate_checkpoint": false,
        "incoming_wals_directory": "/var/lib/barman/pg/incoming",
        "last_backup_maximum_age": null,
        "last_backup_minimum_size": null,
        "last_wal_maximum_age": null,
        "max_incoming_wals_queue": null,
        "minimum_redundancy": 0,
        "network_compression": false,
        "parallel_jobs": 1,
        "parallel_jobs_start_batch_period": 1,
        "parallel_jobs_start_batch_size": 10,
        "path_prefix": null,
        "post_archive_retry_script": null,
        "post_archive_script": null,
        "post_backup_retry_script": null,
        "post_backup_script": null,
        "post_delete_script": null,
        "post_delete_retry_script": null,
        "post_recovery_retry_script": null,
        "post_recovery_script": null,
        "post_wal_delete_script": null,
        "post_wal_delete_retry_script": null,
        "pre_archive_retry_script": null,
        "pre_archive_script": null,
        "pre_backup_retry_script": null,
        "pre_backup_script": null,
        "pre_delete_script": null,
        "pre_delete_retry_script": null,
        "pre_recovery_retry_script": null,
        "pre_recovery_script": null,
        "pre_wal_delete_script": null,
        "pre_wal_delete_retry_script": null,
        "primary_checkpoint_timeout": 0,
        "primary_conninfo": null,
        "primary_ssh_command": null,
        "recovery_options": "",
        "recovery_staging_path": null,
        "create_slot": "auto",
        "retention_policy": null,
        "retention_policy_mode": "auto",
        "reuse_backup": null,
        "slot_name": "barman",
        "snapshot_disks": null,
        "snapshot_gcp_project": null,
        "snapshot_instance": null,
        "snapshot_provider": null,
        "snapshot_zone": null,
        "ssh_command": null,
        "streaming_archiver": true,
        "streaming_archiver_batch_size": 0,
        "streaming_archiver_name": "barman_receive_wal",
        "streaming_backup_name": "barman_streaming_backup",
        "streaming_conninfo": "host=database user=streaming_barman password=barman port=5432",
        "streaming_wals_directory": "/var/lib/barman/pg/streaming",
        "tablespace_bandwidth_limit": null,
        "wal_conninfo": null,
        "wal_retention_policy": "main",
        "wal_streaming_conninfo": null,
        "wals_directory": "/var/lib/barman/pg/wals"
    }
}
```


## List backup
Command: `barman -f json list-backups pg`  
Output:
```json
{
  "pg": [
    {
      "backup_id": "20240622T084702",
      "end_time": "Sat Jun 22 08:47:10 2024",
      "end_time_timestamp": "1719046030",
      "retention_status": "-",
      "size": "52.7 MiB",
      "size_bytes": 55210398,
      "status": "DONE",
      "tablespaces": [],
      "wal_size": "32.0 MiB",
      "wal_size_bytes": 33554432
    }
  ]
}
```

## Show backup
Command: `barman -f json show-backup pg 20240622T084702`  
Output:
```json
{
  "pg": {
    "backup_id": "20240622T084702",
    "base_backup_information": {
      "analysis_time": "less than one second",
      "analysis_time_seconds": 0,
      "begin_lsn": "0/3000028",
      "begin_offset": 40,
      "begin_time": "2024-06-22 08:47:02.291562+00:00",
      "begin_time_timestamp": "1719046022",
      "begin_wal": "000000010000000000000003",
      "copy_time": "8 seconds",
      "copy_time_seconds": 8.239298,
      "disk_usage": "36.7 MiB",
      "disk_usage_bytes": 38433182,
      "disk_usage_with_wals": "52.7 MiB",
      "disk_usage_with_wals_bytes": 55210398,
      "end_lsn": "0/4000000",
      "end_offset": 0,
      "end_time": "2024-06-22 08:47:10.556880+00:00",
      "end_time_timestamp": "1719046030",
      "end_wal": "000000010000000000000003",
      "incremental_size": "36.7 MiB",
      "incremental_size_bytes": 38433182,
      "incremental_size_ratio": "-0.00%",
      "number_of_workers": 1,
      "throughput": "4.4 MiB/s",
      "throughput_bytes": 4664618.514829783,
      "timeline": 1
    },
    "catalog_information": {
      "next_backup": "- (this is the latest base backup)",
      "previous_backup": "- (this is the oldest base backup)",
      "retention_policy": "not enforced"
    },
    "pgdata_directory": "/var/lib/postgresql/data",
    "postgresql_version": 160003,
    "status": "DONE",
    "tablespaces": [],
    "wal_information": {
      "compression_ratio": 0,
      "disk_usage": "32.0 MiB",
      "disk_usage_bytes": 33554432,
      "last_available": "000000010000000000000005",
      "no_of_files": 2,
      "timelines": [],
      "wal_rate": "183.68/hour",
      "wal_rate_per_second": 0.05102330047234898
    }
  }
}
```

## Server status
Command: `barman -f json status all`  
Output:
```json
{
  "pg": {
    "active": {
      "description": "Active",
      "message": "True"
    },
    "backups_number": {
      "description": "No. of available backups",
      "message": "1"
    },
    "current_size": {
      "description": "Current data size",
      "message": "37.0 MiB"
    },
    "current_xlog": {
      "description": "Current WAL segment",
      "message": "000000010000000000000006"
    },
    "data_directory": {
      "description": "PostgreSQL Data directory",
      "message": "/var/lib/postgresql/data"
    },
    "disabled": {
      "description": "Disabled",
      "message": "False"
    },
    "first_backup": {
      "description": "First available backup",
      "message": "20240622T084702"
    },
    "is_in_recovery": {
      "description": "Cluster state",
      "message": "in production"
    },
    "last_backup": {
      "description": "Last available backup",
      "message": "20240622T084702"
    },
    "minimum_redundancy": {
      "description": "Minimum redundancy requirements",
      "message": "satisfied (1/0)"
    },
    "passive_node": {
      "description": "Passive node",
      "message": "False"
    },
    "pg_version": {
      "description": "PostgreSQL version",
      "message": "16.3"
    },
    "retention_policies": {
      "description": "Retention policies",
      "message": "not enforced"
    }
  }
}
```