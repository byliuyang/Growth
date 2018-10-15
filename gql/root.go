package gql

import (
	"io/ioutil"
	"strings"

	"github.com/pkg/errors"
)

type RootResolver struct {
	Query
	Mutation
}

// ReadSchemas reads multiple files and concatenate their content into one string.
func ReadSchemas(files ...string) (string, error) {
	builder := strings.Builder{}
	for _, file := range files {
		b, err := ioutil.ReadFile(file)
		if err != nil {
			return "", errors.WithStack(err)
		}
		_, err = builder.Write(b)
		if err != nil {
			return "", errors.WithStack(err)
		}
	}
	return builder.String(), nil
}
