-- sql: Cria a tabela de contas
CREATE TABLE IF NOT EXISTS contas (
    id VARCHAR(36) PRIMARY KEY NOT NULL,
    nome VARCHAR(128) NOT NULL CHECK(TRIM(nome) != '')
)
