var ApolloClient = require('apollo-boost');
var gql = require('graphql-tag');

const client = new ApolloClient.ApolloClient({
    uri: 'localhost:8080/graphql'
});

var assert = require('assert');
describe('GraphQL', function() {
    describe('event', function() {
        it('push event should succeed', function() {

            client.query({
                query: gql`
                query {
                    event(id: 1) {
                        id
                    }
                }
                `,
            })
                .then(data => console.log(data))
                .catch(error => console.error(error));

        });

        it('push event should succeed', function() {
            assert.equal([1,2,3].indexOf(4), -1);
        });
    });
});
