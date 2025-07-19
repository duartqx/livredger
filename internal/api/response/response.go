package response

import (
	"context"
	"encoding/json"
	"errors"
	"html/template"
	"net/http"

	"github.com/duartqx/livredger/internal/api/decoders"

	"github.com/duartqx/livredger/internal/application/services/visualizadores"

	"github.com/duartqx/livredger/internal/common/mimetypes"
	"github.com/duartqx/livredger/internal/common/types"
)

type Templates struct {
	ComBase *template.Template
	Partial *template.Template
	Error   *template.Template
}

func JsonResponse[C any](ctx context.Context, w http.ResponseWriter, response *visualizadores.Resposta[C]) {

	if ctx.Err() != nil {
		return
	}

	w.Header().Set("Content-Type", mimetypes.JSON)

	status := getStatusByError(response.Error)

	if status == http.StatusInternalServerError && response.Error != nil {
		panic(response.Error)
	}

	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		panic(err)
	}
}

func getStatusByError(err error) int {
	switch {
	case err == nil:
		return http.StatusOK
	case errors.Is(err, types.NotFoundError):
		return http.StatusNotFound
	case errors.Is(err, types.BusinessLogicError) ||
		errors.Is(err, types.RequestError) ||
		errors.Is(err, &json.UnmarshalTypeError{}) ||
		errors.Is(err, decoders.DecoderError):
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}

func JsonErrorResponse(w http.ResponseWriter, err error) {

	w.Header().Set("Content-Type", mimetypes.JSON)

	status := getStatusByError(err)

	w.WriteHeader(status)
	w.Write(*marshal(err))
}

func ErrorResponse(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", mimetypes.HTML)

	status := getStatusByError(err)

	w.WriteHeader(status)
	w.Write([]byte(err.Error()))
}

func marshal(err error) *[]byte {
	var res []byte

	switch {
	case errors.Is(err, decoders.DecoderError):
		res, err = json.Marshal(map[string]any{"error": decoders.ParseDecodeError(err)})
	default:
		res, err = json.Marshal(map[string]string{"error": err.Error()})
	}

	if err != nil {
		panic(err)
	}

	return &res
}
