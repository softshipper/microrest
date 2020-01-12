package descriptor

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

/*
 * Parse valid JSON
 */
func TestJsonAndYamlFile(t *testing.T) {

	assert := assert.New(t)

	t.Run("Parse valid JSON", func(t *testing.T) {
		data, err := ioutil.ReadFile("./service_valid.json")
		assert.Nil(err)

		json, err := ParseJson(data)
		assert.Nil(err)
		assert.Equal(80, json.Port)
	})

	t.Run("Parse invalid JSON", func(t *testing.T) {
		data, err := ioutil.ReadFile("./invalid.json")
		assert.Nil(err)

		_, err = ParseJson(data)
		assert.NotNil(err)
	})

	t.Run("Parse valid YAML", func(t *testing.T) {
		data := []byte(`
name: foo-service
port: 80
description: Easy mocking
routes:
- uri: "/greet"
  methods:
  - verb: GET
    request:
      headers:
        accept: application/json
        if-Match: bfc13a64729c4290ef5b2c2730249c88ca92d82d
    response:
      code: 200
      headers:
        Content-Type: application/json
      body:
        text: Hello from foo service
- uri: "/greet"
  methods:
  - verb: GET
    request:
      headers:
        accept: application/json
        if-Match: bfc13a64729c4290ef5b2c2730249c88ca92d82d
    response:
      code: 200
      headers:
        Content-Type: application/json
      body:
        text: Hello from foo service`)

		s, err := ParseYaml(data)
		assert.Nil(err)
		assert.Equal(s.Port, 80)

	})
}
