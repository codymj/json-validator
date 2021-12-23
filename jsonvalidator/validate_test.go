package jsonvalidator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestValidateArraySuccess tests a successful json array payload
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

// TestValidateArrayFailed tests a failed json array payload
func TestValidateArrayFailed(t *testing.T) {
	// stub
	schema := createArraySchema()
	payload := createFailedArrayPayload()

	// invocation
	errors, err := Validate(schema, payload)

	// assertion
	assert.Nil(t, err)
	assert.Equal(t, 1, len(errors))
}

// TestValidateArrayMalformed tests a malformed json array payload
func TestValidateArrayMalformed(t *testing.T) {
	// stub
	schema := createArraySchema()
	payload := createMalformedArrayPayload()

	// invocation
	errors, err := Validate(schema, payload)

	// assertion
	assert.NotNil(t, err)
	assert.Empty(t, errors)
}

// TestValidateObjectSuccess tests a successful json object payload
func TestValidateObjectSuccess(t *testing.T) {
	// stub
	schema := createObjectSchema()
	payload := createSuccessfulObjectPayload()

	// invocation
	errors, err := Validate(schema, payload)

	// assertion
	assert.Nil(t, err)
	assert.Empty(t, errors)
}

// TestValidateObjectFailed tests a failed json object payload
func TestValidateObjectFailed(t *testing.T) {
	// stub
	schema := createObjectSchema()
	payload := createFailedObjectPayload()

	// invocation
	errors, err := Validate(schema, payload)

	// assertion
	assert.Nil(t, err)
	assert.Equal(t, 1, len(errors))
}

// TestValidateObjectMalformed tests a malformed json object payload
func TestValidateObjectMalformed(t *testing.T) {
	// stub
	schema := createArraySchema()
	payload := createMalformedObjectPayload()

	// invocation
	errors, err := Validate(schema, payload)

	// assertion
	assert.NotNil(t, err)
	assert.Empty(t, errors)
}

/*
 * array schemas
 */

// createArraySchema creates a json schema for validating array payloads
func createArraySchema() []byte {
	return []byte(`{"$id":"TestArraySchema","type":"array","items":{"type":"object","$ref":"#/$defs/veggie"},"$defs":{"veggie":{"type":"object","required":["veggieName","veggieLike"],"properties":{"veggieName":{"type":"string"},"veggieLike":{"type":"boolean"}}}}}`)
}

// createSuccessfulArrayPayload creates a json payload that successfully passes validation
func createSuccessfulArrayPayload() []byte {
	return []byte(`[{"veggieName": "name","veggieLike": false}]`)
}

// createFailedArrayPayload creates a json payload that does not pass validation
func createFailedArrayPayload() []byte {
	return []byte(`[{"veggieLike": false}]`)
}

// createMalformedArrayPayload creates an invalid json array payload
func createMalformedArrayPayload() []byte {
	return []byte(`["veggieName": "name","veggieLike": false}]`)
}

/*
 * object schemas
 */

// createObjectSchema creates a json schema for validating object payloads
func createObjectSchema() []byte {
	return []byte(`{"$id":"TestArraySchema","type":"object","required":["veggieName","veggieLike"],"properties":{"veggieName":{"type":"string"},"veggieLike":{"type":"boolean"}}}`)
}

// createSuccessfulObjectPayload creates a json payload that successfully passes validation
func createSuccessfulObjectPayload() []byte {
	return []byte(`{"veggieName": "name","veggieLike": false}`)
}

// createFailedObjectPayload creates a json payload that does not pass validation
func createFailedObjectPayload() []byte {
	return []byte(`{"veggieName": "name"}`)
}

// createMalformedObjectPayload creates an invalid json payload
func createMalformedObjectPayload() []byte {
	return []byte(`{"veggieName": "name"`)
}
