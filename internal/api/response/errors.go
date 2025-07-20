package response

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/duartqx/livredger/internal/api/decoders"
	ce "github.com/duartqx/livredger/internal/common/errors"
)

func GetStatusCodeFromError(err error) int {
	if err == nil {
		return http.StatusOK
	}

	switch {
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
	badRequest []error = []error{ce.BusinessLogicError, ce.RequestError, &json.UnmarshalTypeError{}, decoders.DecoderError}
	notFound   []error = []error{ce.NotFoundError}
)
