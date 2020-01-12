package descriptor

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestRouteError(t *testing.T) {

	bug := &RouteError{
		uri:    "/greet",
		method: "POST",
	}

	bug.Add("Missing the content of the body")

	res := strings.Contains(bug.Error(), "Missing the content of the body")
	assert.True(t, res)

}
