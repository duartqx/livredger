package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"net/http"

	"github.com/duartqx/livredger/internal/api/decoders"
	"github.com/duartqx/livredger/internal/common/mimetypes"
	t "github.com/duartqx/livredger/internal/common/types"
)

type Templates struct {
	ComBase *template.Template
	Partial *template.Template
	Error   *template.Template
}

type Response[C, T any] struct {
	Writer    http.ResponseWriter
	Request   *http.Request
	Resultado *t.Resultado[C, T]
	Template  *template.Template
}

func HandleResponse[C, T any](res *Response[C, T]) {

	if res.Template == nil || res.Request.Header.Get("Accept") == mimetypes.JSON {
		res.Writer.Header().Set("Content-Type", "application/json")

		if err := json.NewEncoder(res.Writer).Encode(res.Resultado); err != nil {
			JsonErrorResponse(res.Writer, fmt.Errorf("%w: %w", t.InternalError, err))
			return
		}

		return
	}

	if err := res.Template.ExecuteTemplate(res.Writer, "resultados", res.Resultado); err != nil {
		panic(err)
	}
}

func writeHeaderStatus(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, t.NotFoundError):
		w.WriteHeader(http.StatusNotFound)
	case
		errors.Is(err, t.BusinessLogicError) ||
			errors.Is(err, &json.UnmarshalTypeError{}) ||
			errors.Is(err, decoders.DecoderError):
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

	if errors.Is(err, decoders.DecoderError) {
		res, err = json.Marshal(map[string]any{"error": decoders.ParseDecodeError(err)})
	} else {
		res, err = json.Marshal(map[string]string{"error": err.Error()})
	}

	if err != nil {
		panic(err)
	}

	return &res
}
