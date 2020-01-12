package verifier

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"microrest/descriptor"
	"net/http"
	"strings"
)

/*
 * Verify if the incoming request is corresponding to the description
 */

// The route is passed as pointer to avoid copy
func VerifyRequest(req *http.Request, s *[]descriptor.Route) error {

	route, err := filterPathContext(req.URL.Path, s)
	if err != nil {
		return err
	}

	method, err := filterReqMethod(req.Method, &route.Methods)
	if err != nil {
		return err
	}

	err = verifyHeader(&req.Header, &method.Request.Headers)
	if err != nil {
		return err
	}

	if req.Method == "POST" || req.Method == "PUT" {

	}

	return nil
}

// Get current route context correspond to path
func filterPathContext(uri string, r *[]descriptor.Route) (*descriptor.Route, error) {

	for _, e := range *r {
		if e.Uri == uri {
			return &e, nil
		}
	}

	msg := fmt.Sprintf("The route %s can not be found.", uri)
	return nil, NewRouteError(msg)
}

// Filter out current request method
func filterReqMethod(method string, m *[]descriptor.Method) (*descriptor.Method, error) {

	for _, e := range *m {
		if e.Verb == method {
			return &e, nil
		}
	}

	return nil, NewMethodNotFoundError()

}

// Verify if the incoming request header is equal to desired header
func verifyHeader(hActual *http.Header, hMustBe *descriptor.Headers) error {

	for k, v := range *hMustBe {

		var msg string

		found := hActual.Get(k)

		if len(strings.TrimSpace(found)) == 0 {
			msg = fmt.Sprintf("The key %s is missing in the request header.", k)
			return errors.New(msg)
		}

		if found != v {
			msg = fmt.Sprintf("The value %s of key %s is not equal to desired value %s.", k, found, v)
			return errors.New(msg)
		}

	}

	return nil

}

// Verify body content only on method POST and PUT
func verifyBodyContent(reader io.Reader, body interface{}) error {

	b, err := ioutil.ReadAll(reader)
	if err != nil {
		return err
	}

	fmt.Println(string(b))

	return nil

}
