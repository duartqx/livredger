package entidade

import "time"

type Lancamento struct {
	Id int64 `json:"id"`

	DataDeCriacao     time.Time  `json:"data_de_criacao"`
	DataDeModificacao *time.Time `json:"data_de_modificacao"`

	Valor            float64    `json:"valor"`
	Descr            string     `json:"description"`
	DataDePagamento  *time.Time `json:"data_de_pagamento"`
	DataDeVencimento *time.Time `json:"data_de_vencimento"`
	Tipo             string     `json:"tipo"`
}

func NewLancamento() *Lancamento {
	return &Lancamento{
		DataDeCriacao: time.Now(),
	}
}
