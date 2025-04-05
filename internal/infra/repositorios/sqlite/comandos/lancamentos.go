package comandos

import (
	"database/sql"
	"fmt"

	"github.com/duartqx/livredger/internal/domain/comandos"
)

type RepositorioDeComandoLancamentos struct{}

func NewRepositorioDeComandoLancamentos() *RepositorioDeComandoLancamentos {
	return &RepositorioDeComandoLancamentos{}
}

// TODO: REFATORAR PARA EVENTOS DE DOMINIO

func (r RepositorioDeComandoLancamentos) Criar(tx *sql.Tx, comando *comandos.CriarLancamento) error {

	dataDePagamento := &comando.DataDePagamento
	if dataDePagamento.IsZero() {
		dataDePagamento = nil
	}

	dataDeVencimento := &comando.DataDeVencimento
	if dataDeVencimento.IsZero() {
		dataDeVencimento = nil
	}

	_, err := tx.Exec(
		`
		INSERT INTO lancamentos (
			valor,
			descr,
			data_de_pagamento,
			data_de_vencimento,
			tipo
		) VALUES (
			:valor,
			:descr,
			:data_de_pagamento,
			:data_de_vencimento,
			:tipo
		)
		`,
		sql.Named("valor", comando.Valor),
		sql.Named("descr", comando.Descricao),
		sql.Named("data_de_pagamento", dataDePagamento),
		sql.Named("data_de_vencimento", dataDeVencimento),
		sql.Named("tipo", comando.Tipo),
	)

	if err != nil {
		return fmt.Errorf("%w: Não foi possível inserir novo lançamento", err)
	}

	return nil
}
