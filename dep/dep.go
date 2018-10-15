package dep

import (
	"net/http"

	"Growth/core/adapter"
	"Growth/gql"
	"Growth/server"
)

type Dep struct {
	EventStore adapter.EventStore
}

func (d *Dep) RelayHandler(schema string) http.Handler {
	r := gql.RootResolver{
		Query: gql.Query{
			EventStore: d.EventStore,
		},
		Mutation: gql.Mutation{
			EventStore: d.EventStore,
		},
	}
	return server.RelayHandler(schema, r)
}
