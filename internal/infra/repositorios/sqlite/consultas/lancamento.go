package consultas

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/duartqx/livredger/internal/common"
	e "github.com/duartqx/livredger/internal/domain/entidade"
)

type RepositorioDeConsultaLancamentos struct{}

func NewRepositorioDeConsultaLancamentos() *RepositorioDeConsultaLancamentos {
	return &RepositorioDeConsultaLancamentos{}
}

func (r RepositorioDeConsultaLancamentos) BuscarPorId(db *sql.DB, id int) (*e.Lancamento, error) {
	row := db.QueryRow(
		`
		SELECT
			id,
			data_de_criacao,
			data_de_modificacao,
			valor,
			descr,
			data_de_pagamento,
			data_de_vencimento,
			tipo
		FROM lancamentos
		WHERE id = :id
		LIMIT 1
		`,
		sql.Named("id", id),
	)

	var lancamento e.Lancamento

	err := row.Scan(
		&lancamento.Id,
		&lancamento.DataDeCriacao,
		&lancamento.DataDeModificacao,
		&lancamento.Valor,
		&lancamento.Descr,
		&lancamento.DataDePagamento,
		&lancamento.DataDeVencimento,
		&lancamento.Tipo,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("%w: Lançamento não encontrado com id {%d}", common.NotFoundError, id)
		}
		return nil, fmt.Errorf("Não foi possível mapear lançamento: %w", err)
	}

	return &lancamento, nil
}

func (r RepositorioDeConsultaLancamentos) Buscar(db *sql.DB) (*[]e.Lancamento, error) {
	return nil, nil
}
