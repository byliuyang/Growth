package main

import (
	"Growth/core/adapter/testadapter"
	"Growth/dep"
	"Growth/server"
	"io/ioutil"
)

func mustReadFile(filename string) string {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func main() {
	d := dep.Dep{
		EventStore: &testadapter.FakeEventStore{
			Capacity:10,
		},
	}

	schema := mustReadFile("graphql/schema/schema.graphqls")

	handler := d.RelayHandler(schema)
	server.Listen(":8080", handler)
}
