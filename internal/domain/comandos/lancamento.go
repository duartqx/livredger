package comandos

import (
	"fmt"
	"time"
)

type CriarLancamento struct {
	Valor            float64   `json:"valor"`
	Descricao        string    `json:"descricao"`
	DataDePagamento  time.Time `json:"data_de_pagamento"`
	DataDeVencimento time.Time `json:"data_de_vencimento"`
	Tipo             string    `json:"tipo"`
}

func (c CriarLancamento) Validar() error {
	if c.Descricao == "" {
		return fmt.Errorf("Descrição é obrigatória")
	}

	if c.Tipo == "" {
		return fmt.Errorf("Tipo inválido")
	}

	return nil
}
