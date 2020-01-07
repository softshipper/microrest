package descriptor

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

/*
 * Parse valid JSON
 */
func TestParseValidJson(t *testing.T) {
	assert := assert.New(t)

	data := []byte(`
{
  "name": "foo-service",
  "port": 80,
  "description": "Easy mocking",
  "routes": [
    {
      "uri": "/greet",
      "methods": [
        {
          "verb": "GET",
          "request": {
            "headers": {
              "accept": "application/json",
              "if-Match": "bfc13a64729c4290ef5b2c2730249c88ca92d82d"
            }
          },
          "response": {
            "code": 200,
            "headers": {
              "Content-Type": "application/json"
            },
            "body": {
              "text": "Hello from foo service"
            }
          }
        }
      ]
    },
    {
      "uri": "/greet",
      "methods": [
        {
          "verb": "GET",
          "request": {
            "headers": {
              "accept": "application/json",
              "if-Match": "bfc13a64729c4290ef5b2c2730249c88ca92d82d"
            }
          },
          "response": {
            "code": 200,
            "headers": {
              "Content-Type": "application/json"
            },
            "body": {
              "text": "Hello from foo service"
            }
          }
        }
      ]
    }
  ]
}`)

	json, err := ParseJson(data)
	assert.Nil(err)
	assert.Equal(80, json.Port)
}

/*
 * Parse invalid JSON
 */
func TestParseInvalidJson(t *testing.T) {
	assert := assert.New(t)

	data := []byte(`
  "description": "Easy mocking",
  "routes": [
    {
      "route": {
        "uri": "/greet",
        "methods": [
          {
            "verb": "GET",
            "request": {
              "headers": {
                "accept": "application/json",
                "if-Match": "bfc13a64729c4290ef5b2c2730249c88ca92d82d"
              }
            },
            "response": {
              "code": 200,
              "headers": {
                "Content-Type": "application/json"
              },
              "body": {
                "text": "Hello from foo service"
              }
            }
          }
        ]
      }
    }
  ]
}`)

	_, err := ParseJson(data)
	assert.NotNil(err)

}

/*
 * Parse valid YAML
 */
func TestParseValidYaml(t *testing.T) {

	assert := assert.New(t)

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

}
