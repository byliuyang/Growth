package dep

import (
	"net/http"

	"Growth/core/adapter"
	"Growth/gql"
	"Growth/server"

	"github.com/pkg/errors"
)

type Dep struct {
	EventStore  adapter.EventStore
	SchemaPaths []string
	Err         error
}

func (d *Dep) RelayHandler() http.Handler {
	schema, err := gql.ReadSchemas(d.SchemaPaths...)
	if err != nil {
		d.Err = errors.WithStack(err)
		return nil
	}

	r := gql.RootResolver{
		Query: gql.Query{
			EventStore: d.EventStore,
		},
		Mutation: gql.Mutation{
			EventStore: d.EventStore,
		},
	}

	h, err := server.RelayHandler(schema, r)
	d.Err = errors.WithStack(err)
	return h
}
