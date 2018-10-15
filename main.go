package main

import (
	"io/ioutil"

	"Growth/core/adapter/testadapter"
	"Growth/dep"
	"Growth/server"
)

func main() {
	d := dep.Dep{
		EventStore: &testadapter.FakeEventStore{
			Capacity: 10,
		},
	}

	schema := mustReadFile("gql/schema/schema.graphql")

	handler := d.RelayHandler(schema)
	server.Start("localhost:8080", handler)
}

func mustReadFile(filename string) string {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return string(b)
}
