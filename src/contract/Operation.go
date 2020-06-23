package contract

import (
	"net/http"
)

// Operation ...
type Operation interface {
	ID() string
	Name() string
	Description() string
	Method() string
	Path() string
	CreateRequest() *http.Request

	Data() *OperationData
}

// OperationData is an interface to access data from various sources
// (spec path, spec op, cli input, test output) needed in order to build
// an http.Request instance.
type OperationData struct {
	URL     Parameters
	Query   Parameters
	Headers Parameters
}
