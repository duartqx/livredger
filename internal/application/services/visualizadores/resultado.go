package visualizadores

import "encoding/json"

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

type Resposta[C any] struct {
	Consulta  *Consulta[C] `json:"consulta"`
	Resultado any          `json:"resultado"`
	Error     error        `json:"error"`
}

func (r Resposta[C]) MarshalJSON() ([]byte, error) {
	type Alias Resposta[C]

	var errStr *string
	if r.Error != nil {
		s := r.Error.Error()
		errStr = &s
	}

	return json.Marshal(&struct {
		Alias
		Error *string `json:"error"`
	}{
		Alias: Alias(r),
		Error: errStr,
	})
}
