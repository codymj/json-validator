package jsonvalidator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestValidateArraySuccess
func TestValidateArraySuccess(t *testing.T) {
	// stub
	schema := createArraySchema()
	payload := createSuccessfulArrayPayload()

	// invocation
	errors, err := Validate(schema, payload)

	// assertion
	assert.Nil(t, err)
	assert.Empty(t, errors)
}

/*
 * helper functions
 */

// createArraySchema creates a json schema for validating array payloads
func createArraySchema() []byte {
	return []byte(`{"$id":"TestArraySchema","type":"array","items":{"type":"object","$ref":"#/$defs/veggie"},"$defs":{"veggie":{"type":"object","required":["veggieName","veggieLike"],"properties":{"veggieName":{"type":"string"},"veggieLike":{"type":"boolean"}}}}}`)
}

// createSuccessfulArrayPayload creates a json payload that successfully passes validation
func createSuccessfulArrayPayload() []byte {
	return []byte(`[{"veggieName": "name","veggieLike": false}]`)
}
