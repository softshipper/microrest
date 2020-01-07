package verifier

/*
 * When the route can not be found in the request context
 */

func NewRouteError(text string) error {
	return &RouteError{text}
}

type RouteError struct {
	msg string
}

func (r *RouteError) Error() string {
	return r.msg
}
