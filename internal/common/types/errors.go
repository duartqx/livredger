package types

import "errors"

var (
	NotFoundError      error = errors.New("NotFoundError")
	BusinessLogicError       = errors.New("BusinessLogicError")
)
