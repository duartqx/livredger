package types

import (
	"errors"
)

var (
	BusinessLogicError = errors.New("BusinessLogicError")
	InternalError      = errors.New("InternalError")
	NotFoundError      = errors.New("NotFoundError")
	TimeOut            = errors.New("TimeOut")
)
