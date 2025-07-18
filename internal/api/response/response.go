package response

import (
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

func JsonResponse[C any](w http.ResponseWriter, response *visualizadores.Resposta[C]) {

	w.Header().Set("Content-Type", mimetypes.JSON)

	writeHeaderStatus(w, response.Error)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		panic(err)
	}
}

func writeHeaderStatus(w http.ResponseWriter, err error) {
	switch {
	case err == nil:
		return
	case errors.Is(err, types.NotFoundError):
		w.WriteHeader(http.StatusNotFound)
	case errors.Is(err, types.BusinessLogicError) || errors.Is(err, &json.UnmarshalTypeError{}) || errors.Is(err, decoders.DecoderError):
		w.WriteHeader(http.StatusBadRequest)
	default:
		panic(err)
	}
}

func JsonErrorResponse(w http.ResponseWriter, err error) {

	w.Header().Set("Content-Type", mimetypes.JSON)

	writeHeaderStatus(w, err)

	w.Write(*marshal(err))
}

func ErrorResponse(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", mimetypes.HTML)

	writeHeaderStatus(w, err)

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
