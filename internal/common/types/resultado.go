package types

type Enumerated[T any] struct {
	Index   int
	Item    T
	IsFirst bool
	IsLast  bool
}

type Resultado[C, T any] struct {
	Total    int   `json:"total"`
	Consulta *C    `json:"consulta"`
	Itens    *[]*T `json:"itens"`
}

func (r Resultado[C, T]) Enumerated() *[]Enumerated[T] {
	enumerated := []Enumerated[T]{}

	for index, item := range *r.Itens {
		enumerated = append(
			enumerated,
			Enumerated[T]{Index: index, Item: *item, IsFirst: index == 0, IsLast: index == r.Total-1},
		)
	}

	return &enumerated
}
