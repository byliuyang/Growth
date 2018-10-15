var __makeTemplateObject = (this && this.__makeTemplateObject) || function (cooked, raw) {
    if (Object.defineProperty) { Object.defineProperty(cooked, "raw", { value: raw }); } else { cooked.raw = raw; }
    return cooked;
};
var ApolloClient = require('apollo-boost');
var gql = require('graphql-tag');
var client = new ApolloClient.ApolloClient({
    uri: 'localhost:8080/graphql'
});
var assert = require('assert');
describe('GraphQL', function () {
    describe('event', function () {
        it('push event should succeed', function () {
            client.query({
                query: gql(__makeTemplateObject(["\n                query {\n                    event(id: 1) {\n                        id\n                    }\n                }\n                "], ["\n                query {\n                    event(id: 1) {\n                        id\n                    }\n                }\n                "])),
            })
                .then(function (data) { return console.log(data); })
                .catch(function (error) { return console.error(error); });
        });
        it('push event should succeed', function () {
            assert.equal([1, 2, 3].indexOf(4), -1);
        });
    });
});
