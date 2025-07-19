package executores

import (
	"encoding/json"

	"github.com/duartqx/livredger/internal/common/errors"
)

type Response[C any] struct {
	Comando   *C    `json:"comando"`
	Resultado any   `json:"resultado"`
	Error     error `json:"error"`
}

func (r Response[C]) MarshalJSON() ([]byte, error) {
	type Alias Response[C]
	return json.Marshal(&struct {
		Alias
		Error *string `json:"error"`
	}{
		Alias: Alias(r),
		Error: errors.Stringer(r.Error),
	})
}

func (r *Response[C]) WithResultado(resultado any) *Response[C] {
	r.Resultado = resultado
	return r
}

func (r *Response[C]) WithError(err error) *Response[C] {
	r.Error = err
	return r
}
