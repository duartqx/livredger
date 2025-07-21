package visualizadores

import (
	"encoding/json"

	ce "github.com/duartqx/livredger/internal/common/errors"
)

type Enumerated[Item any] struct {
	Index   int
	Item    Item
	IsFirst bool
	IsLast  bool
}

type Result[Item any] struct {
	Total int      `json:"total"`
	Itens *[]*Item `json:"itens"`
}

func (r Result[Item]) Enumerated() *[]Enumerated[Item] {
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

type Query[C any] struct {
	Parsed *C  `json:"parsed"`
	Raw    any `json:"raw"`
}

type Response[C any] struct {
	Query  *Query[C] `json:"query"`
	Result any       `json:"result"`
	Error  error     `json:"error"`
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

func (r *Response[C]) WithError(err error) *Response[C] {
	r.Error = err
	return r
}

func (r *Response[C]) WithQuery(query *Query[C]) *Response[C] {
	r.Query = query
	return r
}

func (r *Response[C]) WithResult(result any) *Response[C] {
	r.Result = result
	return r
}
