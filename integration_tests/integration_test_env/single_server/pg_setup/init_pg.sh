export PGHOST=database
export PGUSER=tester
export PGPASSWORD=tester
export PGDATABASE=postgres
psql -c "create user streaming_barman with replication password 'barman';"

psql -c "CREATE DATABASE test_db;"
export PGDATABASE=test_db
psql -c "CREATE TABLE test(id int);"

# Insert useless data
for i in {1..5}; do psql -c "INSERT INTO test(id) VALUES($i);"; done
# Generate some WAL files
for i in {1..3}; do psql -c "SELECT pg_switch_wal();"; done