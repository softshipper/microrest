package verifier

import (
	"errors"
	"fmt"
	"microrest/descriptor"
	"net/http"
	"strings"
)

/*
 * Verify if the incoming request is corresponding to the description
 */

// The route is passed as pointer to avoid copy
func verifyRequest(req *http.Request, s *[]descriptor.Route) error {
}

// Get current route context correspond to uri
func filterUriContext(uri string, r []descriptor.Route) (descriptor.Route, error) {

	for _, e := range r {
		if e.Uri == uri {
			return e, nil
		}
	}

	msg := fmt.Sprintf("The route %s can not be found.", uri)
	var empty descriptor.Route
	return empty, NewRouteError(msg)
}

// Verify if the incoming request header is equal to desired header
func verifyHeader(hActual http.Header, hMustBe descriptor.Headers) error {

	for k, v := range hMustBe {

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
