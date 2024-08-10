#!/bin/bash
barman cron
echo "Barman init done"
/opt/exporter -config /opt/exporter_config.yaml &
tail -f /dev/null