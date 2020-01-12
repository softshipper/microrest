package descriptor

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"strings"
	"testing"
)

/*
 * Test if the JSON description correspond to HTTP norm
 */


func TestInvalidDescription1(t *testing.T) {

	assert := assert.New(t)

	file, err := ioutil.ReadFile("./service_invalid_1.json")
	assert.Nil(err)

	service, err := ParseJson(file)
	assert.Nil(err)

	msg := ValidateDescription(&service).Error()

	assert.True(strings.Contains("✖ Invalid service name", msg))
	assert.True(strings.Contains("✖ Invalid uri format", msg))
	assert.True(strings.Contains("✖ Method MMM is not valid", msg))

}


func TestInvalidDescription2(t *testing.T) {

	assert := assert.New(t)

	file, err := ioutil.ReadFile("./service_invalid_2.json")
	assert.Nil(err)

	service, err := ParseJson(file)
	assert.Nil(err)

	msg := ValidateDescription(&service).Error()
	assert.True(strings.Contains("✖ The request body is missing", msg))
	assert.True(strings.Contains("✖ The URI to the new location is missing", msg))

}


func TestInvalidDescription3(t *testing.T) {


	assert := assert.New(t)
	file, err := ioutil.ReadFile("./service_invalid_3.json")
	assert.Nil(err)

	service, err := ParseJson(file)
	assert.Nil(err)
	msg := ValidateDescription(&service).Error()
	assert.True(strings.Contains("✖ The response code is missing", msg))
	assert.True(strings.Contains("✖ The response body does not contain deleted resource", msg))

}


func TestInvalidDescription4(t *testing.T) {

	assert := assert.New(t)

	file, err := ioutil.ReadFile("./service_invalid_4.json")
	assert.Nil(err)

	service, err := ParseJson(file)
	assert.Nil(err)
	msg := ValidateDescription(&service).Error()

	assert.True(strings.Contains("✖ The uri /users exists twice", msg))
}
