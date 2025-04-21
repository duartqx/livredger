package http

import (
	"encoding/json"
	"errors"
	"net/http"

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

type Resultado[T any] struct {
	Total int   `json:"total"`
	Itens *[]*T `json:"itens"`
}

func JsonErrorReponse(w http.ResponseWriter, err error) {

	w.Header().Set("Content-Type", "application/json")

	switch {
	case errors.Is(err, t.NotFoundError):
		w.WriteHeader(http.StatusNotFound)
	case errors.Is(err, t.BusinessLogicError) || errors.Is(err, &json.UnmarshalTypeError{}):
		w.WriteHeader(http.StatusBadRequest)
	default:
		panic(err)
	}

	w.Write(*marshal(err))
}

func marshal(err error) *[]byte {
	res, err := json.Marshal(map[string]string{"error": err.Error()})

	if err != nil {
		panic(err)
	}

	return &res
}
