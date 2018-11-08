package gqltest_test

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
	"github.com/stretchr/testify/require"
)

func TestGraphQLAPI(t *testing.T) {
	fakeAdapterStore := testadapter.FakeEventStore{}

	d := dep.Dep{
		EventStore: &fakeAdapterStore,
		SchemaPaths: []string{
			"../schema/schema.graphql",
			"../schema/query.graphql",
			"../schema/mutation.graphql",
			"../schema/types.graphql",
		},
	}

	handler := d.RelayHandler()
	assert.Nil(t, d.Err)

	ts := httptest.NewServer(handler)
	defer ts.Close()

	t.Run("query", func(t2 *testing.T) {
		t2.Run("event", func(t3 *testing.T) {
			fakeAdapterStore.Capacity = 12

			testCases := []struct {
				eventId          int
				expectedResponse string
			}{
				{
					eventId:          1,
					expectedResponse: `{"data":{"event":{"id":1}}}`,
				},
				{
					eventId:          10,
					expectedResponse: `{"errors":[{"message":"event:10 not found","path":["event"]}],"data":{"event":null}}`,
				},
				{
					eventId:          2414143321,
					expectedResponse: `{"errors":[{"message":"could not unmarshal 2.414143321e+09 (float64) into int32: not a 32-bit integer"}],"data":{}}`,
				},
			}

			for _, tc := range testCases {
				t3.Run(fmt.Sprintf("with id=%d", tc.eventId), func(t4 *testing.T) {
					query := fmt.Sprintf(`
			{
				"query": "query getEvent($id: Int!) {event(id: $id) {id}}",
				"variables": {"id": %d}
			}`,
						tc.eventId)

					fakeAdapterStore.Save(entity.Event{})

					res, err := http.Post(ts.URL, "", strings.NewReader(query))
					require.Nil(t4, err)

					require.Equal(t4, http.StatusOK, res.StatusCode)

					b, err := ioutil.ReadAll(res.Body)
					require.Nil(t4, err)
					require.Equal(t4, tc.expectedResponse, string(b))

					fakeAdapterStore.Clear()
				})
			}
		})
	})

	t.Run("mutation", func(t2 *testing.T) {
		t2.Run("newEvent", func(t3 *testing.T) {
			fakeAdapterStore.Capacity = 2

			testCases := []struct {
				existingEventsCount int
				expectedResponse    string
			}{
				{
					existingEventsCount: 0,
					expectedResponse:    `{"data":{"newEvent":{"id":1}}}`,
				},
				{
					existingEventsCount: 2,
					expectedResponse:    `{"errors":[{"message":"event store out of capacity, max: 2","path":["newEvent"]}],"data":{"newEvent":null}}`,
				},
			}

			for _, tc := range testCases {
				t3.Run(fmt.Sprintf("with eventCount=%d", tc.existingEventsCount), func(t4 *testing.T) {
					mutation := `
			{
				"query": "mutation { newEvent {id} }"
			}`
					for i := 0; i < tc.existingEventsCount; i++ {
						fakeAdapterStore.Save(entity.Event{})
					}

					res, err := http.Post(ts.URL, "", strings.NewReader(mutation))
					require.Nil(t4, err)

					require.Equal(t4, http.StatusOK, res.StatusCode)

					b, err := ioutil.ReadAll(res.Body)
					require.Nil(t4, err)
					require.Equal(t4, tc.expectedResponse, string(b))

					fakeAdapterStore.Clear()
				})
			}
		})
	})
}
