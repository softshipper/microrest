package descriptor

import (
	"bytes"
	"fmt"
)

/*
 *
 */

func newRouteErrorLogger(uri string) routeErrorLogger {
	return routeErrorLogger{uri, methodsErrorLogger{}}

}

type routeErrorLogger struct {
	uri     string
	methods methodsErrorLogger
}

func (r *routeErrorLogger) Error() string {

	if len(r.methods) == 0 {
		return ""
	}

	text := fmt.Sprintf("URI: %s\n", s.uri)
	buf := bytes.NewBufferString(text)

	for _, m := range r.methods {

		buf.WriteString(fmt.Sprintf(" - Method: %s\n", m))

		for _, e := range m.errors {
			buf.WriteString(fmt.Sprintf("   %s\n", e))
		}

	}

	return buf.String()

}
