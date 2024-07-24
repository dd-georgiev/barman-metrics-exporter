# Overview
This folder contains the code for the integration tests. The structure is as follows:
```
/integration_tests
   /integration_test_env            # Contains docker containers and docker-compose files
        /single_server              # Barman exporter setup with single postgres server
        /multi_server               # Barman exporter setup with multiple postgres servers
    /tests_utils                    # Contains utilities for tests and custom assertions
    multi_server.shell.test.js      # Contains test cases for multi server with shell integration
    single_server.shell.test.js     # Contains test cases for single server with shell integration
```

# Running
You will need to setup [supported container runtime by TestContainers](https://node.testcontainers.org/supported-container-runtimes/) (I use Docker) and Node v22. 

Make sure that your use can start containers 
To run the tests run:
```
npm install
npm run test
```
# How-to
## Add new metric to test cases
If the output must contain new metric **for each server**
1. Open the `tests_utils/utils.js` file
2. Add the metric name in `ALL_METRICS_NAMES` array. The metric name must be searchable with the following regex:
`${metric}.*server="${server}".*`

If the metric **doesn't contain server label**
1. Open `tests_utils/assertions.js`
2. Add function called `assert#METRIC_NAME#IsPresented`
3. Call the function in `AllMetricsArePresentedForServer` function in the same file.

## Add new test case
Test cases can be added to each indivitual test scenario (a test scenario is for example `multi_server.shell.test.js`). Simply add the desired test in `it` block, as the other test cases. If the assertions contain specifics related to the exporter output, create assertions methods as described in `tests_utils/assertions.js`. Avoid adding too much domain-specific code in the top-level test cases.

The `tests_utils/assertions.js` file contains documentation about adding new assertions.

## Add new integration test scenario
In this project, a scenario is usually related to the integration and specific barman setup. This is the case because the export, must in theory provide the same output no matter what is the underlying system. 

With that in mind, to add sceario:
1. Add the environment specific configuration in new folder inside the `integration_test_env` directory. If the name is not descriptive enough, modify this README file and include comment in the `folder structure` section at the top.

2. Add new file inside the `integration_test` directory following this convention: `SCENARIO_NAME.INTEGRATION.test.js`. Inside this file setup the environment using [TestContainers](https://node.testcontainers.org/features/compose/)(You can use the other environments as example).

3. Add any new assertions in `tests_utils/assertions.js`, all assertions are shared across all scenarios. If the file will grow too big, create separated file/s with the new assertions inside the `test_utils` directory(or subdirectory) and modify the module_exports variable of the `tests_utils/assertions.js` to include the new assertions as well. This can be accomplished using the [spread syntax](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Operators/Spread_syntax)
