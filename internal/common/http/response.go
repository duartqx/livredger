package http

import (
	"encoding/json"
	"errors"
	"net/http"

	t "github.com/duartqx/livredger/internal/common/types"
)

type Resultado[T any] struct {
	Itens *[]*T `json:"itens"`
	Total int   `json:"total"`
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
