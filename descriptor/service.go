package descriptor

type Service struct {
	Name        string
	Port        int
	Description string
	Routes      []Route
}

type Route struct {
	Uri     string
	Methods []Method
}

type Method struct {
	Verb     string
	Request  Request
	Response Response
}

type Headers map[string]string

type Request struct {
	Headers Headers
	Body    interface{}
}

type Response struct {
	Code    int
	Headers Headers
	Body    interface{}
}
