package validator

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/qri-io/jsonschema"
)

// Validate will validate a json payload utilizing a schema
func Validate(schema, payload []byte) (valid bool, errs string, err error) {
	ctx := context.Background()
	rs := &jsonschema.Schema{}

	// unmarshal schema
	err = json.Unmarshal(schema, &rs)
	if err != nil {
		return false, "", err
	}

	// validate schema
	errors, err := rs.ValidateBytes(ctx, payload)
	if err != nil {
		return false, "", err
	}

	// check for errors and append their messages
	valid = true
	strs := make([]string, len(errs))
	if len(errs) > 0 {
		valid = false
		for i := range errs {
			strs[i] = errors[i].Message
		}
	}

	return valid, strings.Join(strs, ", "), nil
}
