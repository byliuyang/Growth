package main

import (
	"fmt"
	"log"

	"Growth/core/adapter/testadapter"
	"Growth/dep"
	"Growth/server"
)

func main() {
	d := dep.Dep{
		EventStore: &testadapter.FakeEventStore{
			Capacity: 10,
		},
		SchemaPaths: []string{"gql/schema/schema.graphql"},
	}

	handler := d.RelayHandler()
	if d.Err != nil {
		panic(fmt.Sprintf("%+v\n", d.Err))
	}

	server.Start("localhost:8080", logFlag, handler)
}

const logFlag = log.Llongfile | log.LUTC | log.LstdFlags
