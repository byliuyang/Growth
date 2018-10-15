package main

import (
	"io/ioutil"
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
	}

	schema := mustReadFile("gql/schema/schema.graphql")
	server.Start("localhost:8080", logFlag, d.RelayHandler(schema))
}

const logFlag = log.Llongfile | log.LUTC | log.LstdFlags

func mustReadFile(filename string) string {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return string(b)
}
