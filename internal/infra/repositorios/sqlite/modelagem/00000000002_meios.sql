-- sql: Cria tabela meios
CREATE TABLE IF NOT EXISTS meios (
    nome VARCHAR(128) PRIMARY KEY NOT NULL,
    CHECK(TRIM(nome) != '')
);

-- sql: Insere meios de transação
INSERT INTO meios (nome)
VALUES
    ('Transferência Bancária'),
    ('PIX'),
    ('Cartão de Crédito'),
    ('Cartão de Débito'),
    ('Dinheiro'),
    ('Cartão de Benefícios')
ON CONFLICT DO NOTHING;
