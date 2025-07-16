-- sql: Cria a tabela de demonstrativos_mensais
CREATE TABLE IF NOT EXISTS demonstrativos_mensais (
    id VARCHAR(36) PRIMARY KEY NOT NULL,
    chave VARCHAR(36) NOT NULL REFERENCES contas(chave),
    timestamp DATETIME DEFAULT (datetime('now')),

    despesa REAL NOT NULL,
    receita REAL NOT NULL,
    saldo REAL NOT NULL,

    mes VARCHAR(7) NOT NULL CHECK(
        LENGTH(mes) = 7
        AND strftime('%Y-%m', mes || '-01') = mes
    )
);
