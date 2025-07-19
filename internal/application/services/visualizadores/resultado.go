package visualizadores

import (
	"encoding/json"

	"github.com/duartqx/livredger/internal/common/errors"
)

type Enumerated[Item any] struct {
	Index   int
	Item    Item
	IsFirst bool
	IsLast  bool
}

type Resultado[Item any] struct {
	Total int      `json:"total"`
	Itens *[]*Item `json:"itens"`
}

func (r Resultado[Item]) Enumerated() *[]Enumerated[Item] {
	enumerated := []Enumerated[Item]{}

	for index, item := range *r.Itens {
		enumerated = append(
			enumerated,
			Enumerated[Item]{
				Index:   index,
				Item:    *item,
				IsFirst: index == 0,
				IsLast:  index == r.Total-1,
			},
		)
	}

	return &enumerated
}

type Consulta[C any] struct {
	Parsed *C  `json:"parsed"`
	Raw    any `json:"raw"`
}

type Response[C any] struct {
	Consulta  *Consulta[C] `json:"consulta"`
	Resultado any          `json:"resultado"`
	Error     error        `json:"error"`
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

func (r *Response[C]) WithError(err error) *Response[C] {
	r.Error = err
	return r
}

func (r *Response[C]) WithConsulta(consulta *Consulta[C]) *Response[C] {
	r.Consulta = consulta
	return r
}

func (r *Response[C]) WithResultado(resultado any) *Response[C] {
	r.Resultado = resultado
	return r
}
