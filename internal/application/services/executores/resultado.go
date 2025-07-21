package executores

import (
	"encoding/json"

	ce "github.com/duartqx/livredger/internal/common/errors"
)

type Response[C any] struct {
	Command *C    `json:"command"`
	Result  any   `json:"result"`
	Error   error `json:"error"`
}

func (r Response[C]) MarshalJSON() ([]byte, error) {
	type Alias Response[C]
	return json.Marshal(&struct {
		Alias
		Error *string `json:"error"`
	}{
		Alias: Alias(r),
		Error: ce.Stringer(r.Error),
	})
}

func (r *Response[C]) WithResult(result any) *Response[C] {
	r.Result = result
	return r
}

func (r *Response[C]) WithError(err error) *Response[C] {
	r.Error = err
	return r
}
