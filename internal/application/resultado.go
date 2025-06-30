package application

type Enumerated[Item any] struct {
	Index   int
	Item    Item
	IsFirst bool
	IsLast  bool
}

type Resultado[Consulta, Item any] struct {
	Total    int       `json:"total"`
	Consulta *Consulta `json:"consulta"`
	Itens    *[]*Item  `json:"itens"`
}

func (r Resultado[Consulta, Item]) Enumerated() *[]Enumerated[Item] {
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
