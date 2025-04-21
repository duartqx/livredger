package types

import "errors"

var (
	BusinessLogicError = errors.New("BusinessLogicError")
	InternalError      = errors.New("InternInternalError")
	NotFoundError      = errors.New("NotFoundError")
)
