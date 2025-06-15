package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"net/http"

	"github.com/duartqx/livredger/internal/api/decoders"
	t "github.com/duartqx/livredger/internal/common/types"
)

const (
	JSON              = "application/json"
	HTML              = "text/html; charset=utf-8"
	HTMX              = "text/htmx"
	PlainText         = "text/plain; charset=utf-8"
	XML               = "application/xml"
	FormURLEncoded    = "application/x-www-form-urlencoded"
	MultipartFormData = "multipart/form-data"
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

	if res.Template == nil || res.Request.Header.Get("Accept") == "application/json" {
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

func JsonErrorResponse(w http.ResponseWriter, err error) {

	w.Header().Set("Content-Type", "application/json")

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

	w.Write(*marshal(err))
}

func ErrorResponse(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "text/html")

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
