package gql_test

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
	fakeAdapterStore := testadapter.FakeEventStore{
		Capacity: 10,
	}

	d := dep.Dep{
		EventStore: &fakeAdapterStore,
		SchemaPaths: []string{
			"schema/schema.graphql",
			"schema/query.graphql",
			"schema/mutation.graphql",
			"schema/types.graphql",
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
				eventId:  1,
				response: `{"data":{"event":{"id":1}}}`,
			},
			{
				eventId:  10,
				response: `{"errors":[{"message":"event:10 not found","path":["event"]}],"data":{"event":null}}`,
			},
			{
				eventId:  2414143321,
				response: `{"errors":[{"message":"could not unmarshal 2.414143321e+09 (float64) into int32: not a 32-bit integer"}],"data":{}}`,
			},
		}

		for _, tc := range testCases {
			t2.Run(fmt.Sprintf("with id=%d", tc.eventId), func(t3 *testing.T) {
				query := fmt.Sprintf(`
			{
				"query": "query getEvent($id: Int!) {event(id: $id) {id}}",
				"variables": {"id": %d}
			}`,
					tc.eventId)

				fakeAdapterStore.Save(entity.Event{})

				res, err := http.Post(ts.URL, "", strings.NewReader(query))
				require.Nil(t, err)

				require.Equal(t, http.StatusOK, res.StatusCode)

				b, err := ioutil.ReadAll(res.Body)
				require.Nil(t, err)
				require.Equal(t, tc.response, string(b))

				fakeAdapterStore.Clear()
			})
		}
	})
}
