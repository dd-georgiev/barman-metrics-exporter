#!/bin/bash
barman cron
barman receive-wal --create-slot pg
barman switch-wal --force --archive --archive-timeout 60 pg
echo "Barman init done"
barman backup all
/opt/exporter -config /opt/exporter_config.yaml &
tail -f /dev/null