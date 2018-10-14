package graphql

import (
	"Growth/graphql/resolver"
	"net/http"

	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)

type SchemaResolver struct {
	resolver.Query
	resolver.Mutation
}



func RelayHandler(s string, resolver SchemaResolver) http.Handler {
	schema := graphql.MustParseSchema(s, &resolver)
	return &relay.Handler{Schema: schema}
}
