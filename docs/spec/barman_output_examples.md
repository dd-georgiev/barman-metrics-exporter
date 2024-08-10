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
Command: `barman -f json replication-status all`  
Output:
```json
{
  "pg-0": {
    "server_lsn": "0/580000C8",
    "streaming_clients": [
      {
        "application_name": "barman_receive_wal",
        "communication": "TCP/IP",
        "current_state": "streaming",
        "current_sync_state": "async",
        "flush_lsn": "0/58000000",
        "flush_lsn_diff": "-200 B",
        "flush_lsn_diff_bytes": -200,
        "host": null,
        "ip_address": "172.19.0.4",
        "name": "Async WAL streamer",
        "port": 47956,
        "replication_slot": "barman",
        "sent_lsn": "0/580000C8",
        "sent_lsn_diff": "0 B",
        "sent_lsn_diff_bytes": 0,
        "started_at": "2024-08-09 19:18:37.386277+00:00",
        "sync_stage": "3/3 Remote write",
        "user_name": "streaming_barman",
        "wal_sender_pid": 63,
        "write_lsn": "0/580000C8",
        "write_lsn_diff": "0 B",
        "write_lsn_diff_bytes": 0
      }
    ]
  },
  "pg-1": {
    "server_lsn": "0/5D0000C8",
    "streaming_clients": [
      {
        "application_name": "barman_receive_wal",
        "communication": "TCP/IP",
        "current_state": "streaming",
        "current_sync_state": "async",
        "flush_lsn": "0/5D000000",
        "flush_lsn_diff": "-200 B",
        "flush_lsn_diff_bytes": -200,
        "host": null,
        "ip_address": "172.19.0.4",
        "name": "Async WAL streamer",
        "port": 60194,
        "replication_slot": "barman",
        "sent_lsn": "0/5D0000C8",
        "sent_lsn_diff": "0 B",
        "sent_lsn_diff_bytes": 0,
        "started_at": "2024-08-09 19:18:37.448315+00:00",
        "sync_stage": "3/3 Remote write",
        "user_name": "streaming_barman",
        "wal_sender_pid": 65,
        "write_lsn": "0/5D0000C8",
        "write_lsn_diff": "0 B",
        "write_lsn_diff_bytes": 0
      }
    ]
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