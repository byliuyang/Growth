package main_test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"Growth/core/adapter/testadapter"
	"Growth/core/entity"
	"Growth/dep"
	"github.com/stretchr/testify/assert"
)

func TestGraphQLAPI(t *testing.T) {
	fakeAdapterStore := testadapter.FakeEventStore{
		Capacity: 10,
	}

	d := dep.Dep{
		EventStore: &fakeAdapterStore,
		SchemaPaths: []string{
			"gql/schema/schema.graphql",
			"gql/schema/query.graphql",
			"gql/schema/mutation.graphql",
			"gql/schema/types.graphql",
		},
	}

	handler := d.RelayHandler()
	assert.Nil(t, d.Err)

	ts := httptest.NewServer(handler)
	defer ts.Close()

	t.Run("query event", func(t2 *testing.T) {
		testCases := []struct {
			eventId  int
			response string
		}{
			{
				eventId:  0,
				response: `{"errors":[{"message":"event:0 not found","path":["event"]}],"data":{"event":null}}`,
			},
			{
				eventId:  1,
				response: `{"data":{"event":{"id":1}}}`,
			},
			{
				eventId:  10,
				response: `{"errors":[{"message":"event:10 not found","path":["event"]}],"data":{"event":null}}`,
			},
		}

		for _, tc := range testCases {
			query := fmt.Sprintf(`
			{
				"query": "query getEvent($id: Int!) {event(id: $id) {id}}",
				"variables": {"id": %d}
			}`,
				tc.eventId)

			fakeAdapterStore.Save(entity.Event{})

			res, err := http.Post(ts.URL, "application/json", strings.NewReader(query))
			assert.Nil(t, err)

			assert.Equal(t, http.StatusOK, res.StatusCode)

			b, err := ioutil.ReadAll(res.Body)
			assert.Nil(t, err)
			assert.Equal(t, tc.response, string(b))

			fakeAdapterStore.Clear()
		}
	})
}
