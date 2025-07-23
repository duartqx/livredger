package response

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/duartqx/livredger/internal/application/services/executores"
	"github.com/duartqx/livredger/internal/application/services/visualizadores"
	"github.com/duartqx/livredger/internal/common/mimetypes"
)

func ErrorResponse(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", string(mimetypes.HTML))

	status := GetStatusCodeFromError(err)

	w.WriteHeader(status)
	w.Write([]byte(err.Error()))
}

func ErrorJsonResponse(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", string(mimetypes.JSON))

	w.WriteHeader(GetStatusCodeFromError(err))

	json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
}

func CommandJsonResponse[C any](w http.ResponseWriter, response *executores.Response[C]) {

	w.Header().Set("Content-Type", string(mimetypes.JSON))

	w.WriteHeader(GetStatusCodeFromError(response.Error))

	if err := json.NewEncoder(w).Encode(response); err != nil {
		panic(err)
	}
}

func QueryJsonResponse[C any](ctx context.Context, w http.ResponseWriter, response *visualizadores.Response[C]) {

	if ctx.Err() != nil {
		return
	}

	w.Header().Set("Content-Type", string(mimetypes.JSON))

	w.WriteHeader(GetStatusCodeFromError(response.Error))

	if err := json.NewEncoder(w).Encode(response); err != nil {
		panic(err)
	}
}
