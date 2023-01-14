package errpkg

import "net/http"

var ErrNotFound = &DomainError{
	ErrorCode:      "E4-0001",
	HTTPStatusCode: http.StatusNotFound,
	Message:        "not found",
}
