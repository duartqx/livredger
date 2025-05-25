package types

type Resultado[C, T any] struct {
	Total    int   `json:"total"`
	Consulta *C    `json:"consulta"`
	Itens    *[]*T `json:"itens"`
}
