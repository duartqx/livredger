-- sql: Cria tabela eventos_lancamento
CREATE TABLE IF NOT EXISTS eventos (
    nome VARCHAR(128) PRIMARY KEY NOT NULL,
    CHECK(TRIM(nome) != '')
);

-- sql: Insere eventos lancamento
INSERT INTO eventos (nome)
VALUES
    ('ContaAberta'),
    ('ContaFechada'),
    ('ContaPausada'),
    ('LancamentoPrevisto'),
    ('LancamentoPago'),
    ('LancamentoRecebido'),
    ('LancamentoTransferido'),
    ('LancamentoCorrigido'),
    ('LancamentoCancelado')
ON CONFLICT DO NOTHING;
