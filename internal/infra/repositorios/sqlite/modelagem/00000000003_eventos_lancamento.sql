-- sql: Cria tabela eventos_lancamento
CREATE TABLE IF NOT EXISTS eventos_lancamento (
    nome VARCHAR(128) PRIMARY KEY NOT NULL,
    CHECK(TRIM(nome) != '')
);

-- sql: Insere eventos lancamento
INSERT INTO eventos_lancamento (nome)
VALUES
    ('LancamentoPrevisto'),
    ('LancamentoPago'),
    ('LancamentoRecebido'),
    ('LancamentoTransferido'),
    ('LancamentoCorrigido'),
    ('LancamentoCancelado')
ON CONFLICT DO NOTHING;
