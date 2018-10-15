const assert = require('assert');
const axios = require('axios');

describe('GraphQL', function () {
    describe('event', function () {
        it('push event should succeed', async function () {
            let resp = await axios.post("http://localhost:8080/graphql", {
                query: `query {
                    event(id: 1) {
                        id
                    }
                }`
            });
            console.log(resp.data);
            assert.equal(resp.data, {
                errors: [
                    {
                        message: "event:1 not found",
                        path: [
                            "event"
                        ]
                    }
                ],
                data: {
                    event: null
                }
            })
        })
    });
});
