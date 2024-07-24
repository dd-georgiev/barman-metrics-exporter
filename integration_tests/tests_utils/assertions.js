// This module provides assertion used by Jest/Supertest test cases.
// Those assertions are specific for the output of the barman exporter. They are mostly set of RegExp for matching specific metric/info about metic.
// The convention is to name exports asserions starting with Capital later and not prefix.
// Internal functions|assertions are prefixed with assertNAME. 
// This is the case because, in theory the package will be imported with require and named assert. Here is snipper for example usage:
/*
const assert = require('assertions')
assert.AllMetricsArePresentedForServer(expect, res.text, SRV_NAME)
*/

// The public assertions must be placed at the top, while the internal one are at the bottom.
// Preferably all regexes will be put in internal assertion(s) with desccriptive name.

const utils = require('./utils')
// Public assertions
/*
I. Convention
    1. All public assertions must take `expect` as first argument. Expect is the jest function.
    2. All public assertions must have JSDoc description for the function as well as for all the parameters it takes.
    3. None of the public assertions must return value
*/
 /**
 * Asserts that all supported metrics are presented. To ensure that the barman_up metric is properly presented, all checks labels are asserted as well
 * @param {function} expect - The set of Jest Matchers. https://jestjs.io/docs/expect
 * @param {string} body - the response body from the exporter
 * @param {string} server - name of the server, used for searching the `server` label(e.g. server="pg-0") 
 */
function AllMetricsArePresentedForServer(expect, body, server) {
    assertAllMetricsArePresentedForServer(expect, body, server)
    assertAllBarmanChecksArePresentedForServer(expect, body, server)
}

/**
 * Asserts that all checks from the barman_up command are either equal to 0 or 1. Other values for this metric are invalid.
 */
function AllBarmanChecksAreCorrect(expect, body, server) {
    assertAllBarmanChecksAreEqualToOneOrZero(expect, body, server)
}
// Internal assertions
/*
I. Convention
    1. All internal assertions must start with lowercase "assert"
    2. If assertiong  needs more than one regex, preferably split it into multiple smaller assertions and combine them in the public assertion.
        NOTE: Creating internal assertion invoking multiple internal assertions is also fine, but preferably avoided without good reason. 
              If there is good reason for creating such abstraction, you must document it with comment and the method must have JSDoc
*/
function assertAllMetricsArePresentedForServer(expect, body, server) {
    for (const metric of utils.ALL_METRICS_NAMES) {
        expect(body).toMatch(new RegExp(`${metric}.*server="${server}".*`))
    }
}
function assertAllBarmanChecksArePresentedForServer(expect, body, server) {
    for (const check of utils.BARMAN_UP_LABELS) {
        expect(body).toMatch(new RegExp(`barman_up{check="${check}".*server="${server}"}.*`))
    }
}
function assertAllBarmanChecksAreEqualToOneOrZero(expect, body, server) {
    const REGEX_MATCH_ZERO_OR_ONE = '[01]'
    for (const check of utils.BARMAN_UP_LABELS) {
        // note: \s is for any whitespace charecter.
        expect(body).toMatch(new RegExp(`barman_up{check="${check}".*server="${server}"}\\s${REGEX_MATCH_ZERO_OR_ONE}`))
    }
}

module.exports = { 
    AllMetricsArePresentedForServer,
    AllBarmanChecksAreCorrect
}