package errors

import (
	"errors"
)

var (
	BusinessLogicError = errors.New("BusinessLogicError")
	InternalError      = errors.New("InternalError")
	NotFoundError      = errors.New("NotFoundError")
	TimeOutError       = errors.New("TimeOutError")
	RequestError       = errors.New("RequestError")
)
