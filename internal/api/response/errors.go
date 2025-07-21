package response

import (
	"errors"
	"net/http"

	ce "github.com/duartqx/livredger/internal/common/errors"
)

func GetStatusCodeFromError(err error) int {
	switch {
	case err == nil:
		return http.StatusOK
	case is(err, notFound...):
		return http.StatusNotFound
	case is(err, badRequest...):
		return http.StatusBadRequest
	default:
		panic(err)
	}
}

func is(err error, errTypes ...error) bool {
	for _, errType := range errTypes {
		if errors.Is(err, errType) {
			return true
		}
	}
	return false
}

var (
	badRequest []error = []error{ce.BusinessLogicError, ce.RequestError}
	notFound   []error = []error{ce.NotFoundError}
)
