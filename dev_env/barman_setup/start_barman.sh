#!/bin/bash
barman cron
barman receive-wal --create-slot pg-0
barman switch-wal --force --archive --archive-timeout 60 pg-0
barman receive-wal --create-slot pg-1
barman switch-wal --force --archive --archive-timeout 60 pg-1
echo "Barman init done"
barman backup all
/opt/exporter -config /opt/exporter_config.yaml &
tail -f /dev/null