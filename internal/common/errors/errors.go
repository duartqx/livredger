package errors

import (
	"errors"
	"fmt"
)

var (
	BusinessLogicError = errors.New("BusinessLogicError")
	InternalError      = errors.New("InternalError")
	NotFoundError      = errors.New("NotFoundError")
	TimeOutError       = fmt.Errorf("%w TimeOutError", InternalError)
	RequestError       = errors.New("RequestError")
)
