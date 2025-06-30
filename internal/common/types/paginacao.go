package types

type Paginacao struct {
	Pagina    int       `json:"pagina"`
	Ordenacao Ordenacao `json:"ordenacao"`
}

type Ordenacao struct {
	Campo   string `json:"campo"`
	Direcao string `json:"direcao" valid:"ASC|DESC"`
}
