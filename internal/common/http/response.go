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

func JsonErrorResponse(w http.ResponseWriter, err error, statusCode int) {
	res, errMarshal := json.Marshal(map[string]string{"error": err.Error()})

	if errMarshal != nil {
		http.Error(w, errMarshal.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if errors.Is(err, t.NotFoundError) {
		w.WriteHeader(http.StatusNotFound)
		w.Write(res)
		return
	}

	w.WriteHeader(statusCode)

	w.Write(res)
}
