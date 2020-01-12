package verifier

import (
	"github.com/stretchr/testify/assert"
	"microrest/descriptor"
	"net/http"
	"strings"
	"testing"
)

func TestSimpleRequest(t *testing.T) {

	req, err := http.NewRequest("GET", "/greet", strings.NewReader(""))
	assert.Nil(t, err)
	req.Header.Add("accept", "text/html")

	route := []descriptor.Route{
		{
			Uri: "/greet",
			Methods: []descriptor.Method{
				{
					Verb: "GET",
					Request: descriptor.Request{
						Headers: descriptor.Headers{
							"accept": "text/html",
						},
					},
				},
			},
		},
	}

  err = VerifyRequest(req, &route)

	assert.Nil(t, err)

}

func TestPostRequestWithBody(t *testing.T) {

	body1 := `key1=value1&key2=value2`
	route1 := []descriptor.Route{
		{
			Uri: "/greet",
			Methods: []descriptor.Method{
				{
					Verb: "GET",
					Request: descriptor.Request{
						Headers: descriptor.Headers{
							"accept": "text/html",
						},
						Body: "",
					},

				},
			},
		},
	}
	req1, err1 := http.NewRequest("POST", "/greet", strings.NewReader(``))
	req1.Header.Add("content-type", "application/x-www-form-urlencoded")
	assert.Nil(t, err1)

	


	req2, err2 := http.NewRequest("POST", "/greet", strings.NewReader(``))
	req2.Header.Add("content-type", "multipart/form-data")
	assert.Nil(t, err2)

	req3, err3 := http.NewRequest("POST", "/greet", strings.NewReader(``))
	req3.Header.Add("content-type", "text/plain")
	assert.Nil(t, err3)

	req4, err4 := http.NewRequest("POST", "/greet", strings.NewReader(`{ "form1": "data1", "form2": "data2" }`))
	req4.Header.Add("accept", "application/json")
	assert.Nil(t, err4)

}

func TestPostRequestWithoutBody(t *testing.T) {

	req, err := http.NewRequest("POST", "/greet", strings.NewReader(``))
	req.Header.Add("accept", "application/json")
	assert.Nil(t, err)
}

func TestPutRequest(t *testing.T) {

	req1, err1 := http.NewRequest("PUT", "/greet/1", strings.NewReader(``))
	req1.Header.Add("content-type", "application/x-www-form-urlencoded")
	assert.Nil(t, err1)

	req2, err2 := http.NewRequest("PUT", "/greet/2", strings.NewReader(`{ "form1": "data1", "form2": "data2" }`))
	req2.Header.Add("accept", "application/json")
	assert.Nil(t, err2)
}

func TestDeleteRequest(t *testing.T) {

	_, err := http.NewRequest("DELETE", "/delete/1", strings.NewReader(``))
	assert.Nil(t, err)
}
