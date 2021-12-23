# jsonvalidator
A wrapper for validating JSON payloads

## Dependencies

* [github.com/xeipuuv/gojsonschema](https://github.com/xeipuuv/gojsonschema)

## Functions

### `Validate(schema, payload []byte) (errors []string, err error)`

| parameter | description                                        |
|-----------|----------------------------------------------------|
| schema    | json schema to be validated against                |
| payload   | json payload to be validated                       |
| errors    | schema validation error messages                   |
| err       | non-schema validation errors (like malformed JSON) |

**Example**

```go
schema := []byte(`{"$id":"TestObjectSchema","type":"object","required":["veggieName","veggieLike"],"properties":{"veggieName":{"type":"string"},"veggieLike":{"type":"boolean"}}}`)
payload := []byte(`{"veggieName": "name"}`)

errors, err := jsonvalidator.Validate(schema, payload)
if err != nil {
  // handle error
}

if len(errors) > 0 {
  // validation errors have occurred
  // errors[0] = "(root): veggieLike is required"
}
```
