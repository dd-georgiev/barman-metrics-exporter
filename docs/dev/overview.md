# Overview
The project provides a set of metrics (defined in `docs/spec`) in prometheus format. The application its self is written in Go. The integration tests are written in NodeJS with [TestContainers](https://node.testcontainers.org/) and [SuperTest](https://www.npmjs.com/package/supertest).

This document provides information about external and internal `dependencies` and `how-to`.

# Dependencies and conventions
## Dependencies
### Config
[https://github.com/gookit/config](https://github.com/gookit/config) is used to handle configuration data. 
The setup is done in `cmd/main.go`, the main function calls `initConfig` which takes care of setting up the library. Currently the idea is that the user will pass `-config` flag to the application pointing to configuration file. The method then setups the config based on this file.

### Metrics
The [otel/metric](https://pkg.go.dev/go.opentelemetry.io/otel/metric) package is used for handling the low-level details around the metric. In `cmd/main.go` the initial setup is done(the `initMetrics`) method. Then individual metrics implementing the `BarmanMetric` interfaces consume the APIs.


[Prometheus client](https://pkg.go.dev/github.com/prometheus/client_golang/prometheus/promhttp) is used to expose the metrics via HTTP in standard prometheus format.

# Internal libraries
## Integration
The integration package is integrating the go application with `barman`. Currently the only integration avaialbe is `shell`, some day we may add `ssh` integration which is executing shell commands via `ssh` connection.
There is `mock_executor.go` which returns hard-coded json used for testing.

## Outputs
The outputs package is a set of `structs` which represent data from `barman`. In addition to the structs, each of the files provides a method for converting `json` as outputted by `barman` to the struct. 
Most of the tests for unmarshalling the `json` output by `barman` are in this package.
## Metrics
The metric package is wrapper around the `otel/metrics` libary. There is single interface `BarmanMetric` which must be implemented by every custom metric the application provides.
The `InitAllMetrics` and `UpdateAllMetrics` are functions which provide means to work with all metrics. The `InitAllMetrics` method is resposnible for creating all custom metrics and returning `slice` of them. The` UpdateAllMetrics` takes this `slice` and calls the `Update` method on each of those metrics. No other operations on all metrics are supported.
# Notes 
1. The json parser may not parse the objects in same order. For example if we have json for two servers:
```yaml
{
    'pg-0': { ... }
    'pg-1': { ... }
}
```

and then we unmarshall into two structs and form an slice from those two structs, it is not guranteed that `pg-0` will be before `pg-1`. Thats why in the `outputs` test cases we have:
```go
if struct_eg[0].Server != "pg-0" && struct_eg[0].Server != "pg-1" { // the server may be pg-0 or pg-1, can't be something else
    // Error - the server name is neither pg-0 nor pg-1
}
if struct_eg[1].Server != "pg-0" && struct_eg[1].Server != "pg-1" { // the server may be pg-0 or pg-1, can't be something else
    // Error - the server name is neither pg-0 nor pg-1
}
if struct_eg[0].Server == struct_eg[1].Server { // make sure that the servers are different
    // Error the two server names are the same. It may be legit case!
}
```

# How-To
## Run locally
1. Navigate to `exporter/cmd/`
2. Run `go run main.go -config config.yaml`

This will start the exporter with `mock` integration and metrics on `localhost:2222/metrics`
## Add config option
Add it in `config.yaml` and then use config.TYPE("OPTION_NAME") `[e.g. config.String("port")]` to retrieve it. Doing so will require importing `"github.com/gookit/config/v2"` in the file.
## Add new metrics
1. Add info about the metric in `docs/spec/metrics.md`
2. If new barman command/output will be used:
- Add info about the barman output in `docs/spec/barman_output_examples.md`
- Add the type(s) in `internal/integration/outputs/TYPE_NAME.go`
- Add method for retrieving the data in `CommandExecutor`(`integration.go`) and implement it in `shell_executor.go` and `mock_executor.go`. 
- Add method for retrieving the data for the `BarmanIntegration` type together with caching
- Add tests for the new output type and integration methods in `integration_test.go`, add example test json in `integration_test.go`
3. Add new file in `internal/metrics/` with the name of the metric
4. Add struct in the file created in `3.` implementing the `BarmanMetric` in for the new metric
5. Modify the `InitAllMetrics`(in `metrics.go`) to include the new metric
6. Modify the integration tests to include the new metric.
You can use `internal/metrics/barman_total_backups.go` as example. It is the simplest metric.
For new output types, the `internal/outputs/server_backups.go` is the simplest example, while `internal/outputs/server_status.go` showcases parsing complex JSON, though with few fields being extracted.
## Run local dev env
The `dev_env` directory contains environment, meant to be used locally for testing. The specification of the environment is as follows:
1. Two Postgres 16 containers (with config from `dev_env/pg_config|pg_hba`)
2. Barman v3.10.0 in Debian container (with config from `dev_env/barman.conf`)
3. Init/Provisioning container which executes the script located in `dev_env/pg_setup/init_pg.sh`
4. Barman exporter running inside the Barman container with config from `dev_env/exporter_setup/config.yaml`, exported on local port `2220`. The config sets up debug log level and shell integration 
5. Container with barman, barman exporter and `systemd` to test OS signal handling and changes to the systemd unit file.
The exporter is being built with `exporter/build.Dockerfile` container. So before starting the environment you must have this container built locally. To run the build:
1. Go to `exporter` directory(where the Go code and the build.Dockerfile are located)
2. Run `docker build --progress=plain . -f build.Dockerfile -t exporter`

To start the environment navigate into the `dev_env` directory and run:
```bash
docker-compose up -d # NOTE: It usually takes 60-90 seconds to fully provision the postgres instances, start barman and take backup. Remove '-d' to see all logs.
# If you need to update the exporter code run:
docker-compose up -d --build
# If you suspect that the code isn't updated[shouldn't be the case generally speaking]:
docker-compose build --no-cache && docker-compose up -d
```

To get shell to the barman container(where both barman and the exporter are running)
```bash
docker-compose exec barman /bin/bash
```

To stop the environment:
```bash
docker-compose down
```

## Run integration tests
Visit the README.md in the `integration_tests` directory for more information about running the integration tests locally. 