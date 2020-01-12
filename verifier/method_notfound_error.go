package verifier

/*
 * When the request context does not match to any http method
 */

func NewMethodNotFoundError() error {
	return &RouteError{}
}

type MethodNotFoundError struct {
}

func (m *MethodNotFoundError) Error() string {
	return ""
}
