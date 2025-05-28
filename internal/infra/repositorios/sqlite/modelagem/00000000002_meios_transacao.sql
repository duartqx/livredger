-- sql: Cria tabela meios_financeiro
CREATE TABLE IF NOT EXISTS meios_financeiro (
    nome VARCHAR(128) PRIMARY KEY NOT NULL,
    CHECK(TRIM(nome) != '')
);

-- sql: Insere meios financeiros
INSERT INTO meios_financeiro (nome)
VALUES
    ('Cartão de Benefícios'),
    ('Cartão de Crédito'),
    ('Cartão de Débito'),
    ('Dinheiro'),
    ('PIX'),
    ('Poupança'),
    ('Transferência Bancária')
ON CONFLICT DO NOTHING;
