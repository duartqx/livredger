-- sql: Cria tabela de naturezas
CREATE TABLE IF NOT EXISTS naturezas (
    nome VARCHAR(128) PRIMARY KEY NOT NULL,
    CHECK(TRIM(nome) != '')
);

-- sql: Insere naturezas conhecidas
INSERT INTO naturezas (nome)
VALUES
    ('Salário'),
    ('Benefícios'),
    ('Compras'),
    ('Mercado'),
    ('Luz'),
    ('Condomínio'),
    ('Água e Gás'),
    ('Telefonia'),
    ('Nuvem'),
    ('Internet'),
    ('Receita Extra'),
    ('Petshop'),
    ('Saúde'),
    ('Investimento'),
    ('Outro'),
ON CONFLICT DO NOTHING;
