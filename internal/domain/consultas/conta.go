package consultas

type ConsultaContas struct {
	Nome string `json:"nome" form:"nome"`
}

func ConsultaContasPadrao() *ConsultaContas {
	return &ConsultaContas{}
}
