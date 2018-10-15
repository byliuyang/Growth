package main

import (
	"log"

	"Growth/core/adapter/testadapter"
	"Growth/dep"
	"Growth/gql"
	"Growth/server"
)

func main() {
	d := dep.Dep{
		EventStore: &testadapter.FakeEventStore{
			Capacity: 10,
		},
	}

	schema, err := gql.ReadSchemas("gql/schema/schema.graphql")
	if err != nil {
		panic(err)
	}

	server.Start("localhost:8080", logFlag, d.RelayHandler(schema))
}

const logFlag = log.Llongfile | log.LUTC | log.LstdFlags
