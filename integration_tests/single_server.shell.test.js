const { DockerComposeEnvironment, Wait } = require("testcontainers");
const assert = require('./tests_utils/assertions')
const composeFilePath = "integration_test_env/single_server"
const composeFile = "docker-compose.yaml";
const request = require('supertest');

const SECONDS = 1000
const MINUTES = 60*SECONDS
describe("Barman exporter with single postgres servers", () => { 
    let environment
    let res
    jest.setTimeout(10 * MINUTES)
    beforeAll(async () => { 
        environment = await new DockerComposeEnvironment(composeFilePath, composeFile)
        .withWaitStrategy("barman", Wait.forLogMessage("serving metrics atserving metrics at localhost:2222/metrics"))
        .withStartupTimeout(5 * MINUTES)
        .up();

        // let the metrics collect and be exposed
        const delay = ms => new Promise(resolve => setTimeout(resolve, ms))
        await delay(30 * SECONDS)


        const req = request("http://localhost:2223")
        res = await req.get('/metrics')

    })

    afterAll(async () => { 
        await environment.down()
    })

    it("Must return response with status 200 OK", () => { 
        expect(res.status).toBe(200)
    })
    it("Must contain all exposed metrics", () => {
        assert.AllMetricsArePresentedForServer(expect, res.text, "pg")
    })
    it("Must contain valid values for barman checks", async () => { 
        assert.AllBarmanChecksAreCorrect(expect, res.text, "pg")
    })
})