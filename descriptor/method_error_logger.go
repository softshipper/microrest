package descriptor

import "fmt"

type methodsErrorLogger []methodErrorLogger

type methodErrorLogger struct {
	method string
	errors []string
}

func newMethodErrorLogger(method string) methodErrorLogger {
	return methodErrorLogger{method, []string{}}
}

func (m *methodErrorLogger) add(msg string) {
	text := fmt.Sprintf("✖ %s", msg)
	m.errors = append(m.errors, text)
}

func (m *methodErrorLogger) any() bool {
	return len(m.errors) > 0
}
