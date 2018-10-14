package dep

import (
	"Growth/core/adapter"
	"Growth/graphql"
	"Growth/graphql/resolver"
	"net/http"
)

type Dep struct {
	EventStore adapter.EventStore
}

func (d *Dep) RelayHandler(schema string) http.Handler {
	r := graphql.SchemaResolver{
		Query:resolver.Query{
			EventStore:d.EventStore,
		},
		Mutation: resolver.Mutation{
			EventStore:d.EventStore,
		},
	}
	return graphql.RelayHandler(schema, r)
}
