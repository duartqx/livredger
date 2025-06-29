package consultas

type ConsultaContas struct {
	Nome string `json:"nome"`
}

func ConsultaContasPadrao() *ConsultaContas {
	return &ConsultaContas{}
}
