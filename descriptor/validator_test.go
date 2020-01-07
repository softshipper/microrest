package descriptor

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

/*
 * Test if the JSON description correspond to HTTP norm
 */
func TestPath(t *testing.T) {
	assert := assert.New(t)

	uri := "/hello"

	err := ValidateUri(uri)
	assert.Nil(err)

}

func TestHttpVerb(t *testing.T) {
	assert := assert.New(t)

	valid := ValidateHttpVerb("GET")
	assert.Nil(valid)

	invalid := ValidateHttpVerb("TTT")
	assert.NotNil(invalid)
	assert.Equal("The given TTT method is not supported!", invalid.Error())

}

func TestHttpCode(t *testing.T) {
	assert := assert.New(t)

	valid := ValidateHttpCode(http.StatusAccepted)
	assert.Nil(valid)

	invalid := ValidateHttpCode(3432)
	assert.Error(invalid, "The http code 3432 does not exist!")

}
