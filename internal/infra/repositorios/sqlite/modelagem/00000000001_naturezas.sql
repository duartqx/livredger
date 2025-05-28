-- sql: Cria tabela de naturezas
CREATE TABLE IF NOT EXISTS naturezas (
    nome VARCHAR(128) PRIMARY KEY NOT NULL,
    CHECK(TRIM(nome) != '')
);

-- sql: Insere naturezas conhecidas
INSERT INTO naturezas (nome)
VALUES
    ('Água e Gás'),
    ('Benefícios'),
    ('Compras'),
    ('Condomínio'),
    ('Internet'),
    ('Investimento'),
    ('Juros'),
    ('Luz'),
    ('Mercado'),
    ('Nuvem'),
    ('Outro'),
    ('Petshop'),
    ('Receita Extra'),
    ('Salário'),
    ('Saúde'),
    ('Telefonia')
ON CONFLICT DO NOTHING;
