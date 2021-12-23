package jsonvalidator

import (
	"fmt"

	"github.com/xeipuuv/gojsonschema"
)

// Validate will validate a JSON payload against a schema
func Validate(schema, payload []byte) (errors []string, err error) {
	// init errors
	errors = make([]string, 0)

	// validate
	ls := gojsonschema.NewBytesLoader(schema)
	ld := gojsonschema.NewBytesLoader(payload)
	result, err := gojsonschema.Validate(ls, ld)
	if err != nil {
		return errors, err
	}

	// check for validation errors
	if !result.Valid() {
		for _, e := range result.Errors() {
			errors = append(errors, fmt.Sprintf("%s", e.String()))
		}
	}

	return errors, nil
}
