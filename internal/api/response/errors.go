package response

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/duartqx/livredger/internal/api/decoders"
	"github.com/duartqx/livredger/internal/common/types"
)

func GetStatusCodeFromError(err error) int {
	switch {
	case err == nil:
		return http.StatusOK
	case errors.Is(err, types.NotFoundError):
		return http.StatusNotFound
	case errors.Is(err, types.BusinessLogicError),
		errors.Is(err, types.RequestError),
		errors.Is(err, &json.UnmarshalTypeError{}),
		errors.Is(err, decoders.DecoderError):
		return http.StatusBadRequest
	default:
		panic(err)
	}
}
