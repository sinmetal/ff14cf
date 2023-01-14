package errpkg

import "fmt"

type errorsIs interface {
	Is(error) bool
}

type errorsAs interface {
	As(interface{}) bool
}

var _ errorsIs = &DomainError{}
var _ errorsAs = &DomainError{}

type DomainError struct {
	ErrorCode      string `json:"errorCode"`
	HTTPStatusCode int    `json:"-"`
	Message        string `json:"message"`
}

func (err *DomainError) Error() string {
	return fmt.Sprintf("%s: %s", err.ErrorCode, err.Message)
}

func (err *DomainError) Is(other error) bool {
	derr, ok := other.(*DomainError)
	if !ok {
		return false
	}
	return err.ErrorCode == derr.ErrorCode
}

func (err *DomainError) As(target interface{}) bool {
	switch target := target.(type) {
	case **DomainError:
		*target = &DomainError{
			ErrorCode: err.ErrorCode,
			Message:   err.Message,
		}
		return true
	default:
		return false
	}
}
