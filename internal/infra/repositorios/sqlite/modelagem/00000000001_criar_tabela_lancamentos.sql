CREATE TABLE IF NOT EXISTS lancamentos (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    evento VARCHAR(128) NOT NULL,
    timestamp DATETIME DEFAULT (datetime('now')),

    chave VARCHAR(36) NOT NULL,
    versao INTEGER NOT NULL,

    valores REAL NOT NULL,
    vencimento DATETIME,

    descr VARCHAR(500) NOT NULL,

    UNIQUE (chave, versao)
);

CREATE INDEX idx_lancamentos_evento ON lancamentos (evento);

CREATE INDEX idx_lancamentos_timestamp ON lancamentos (timestamp);
