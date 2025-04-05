CREATE TABLE IF NOT EXISTS lancamentos (
    id INTEGER PRIMARY KEY AUTOINCREMENT,

    data_de_criacao DATETIME DEFAULT (datetime('now')),
    data_de_modificacao DATETIME DEFAULT NULL,

    valor REAL NOT NULL,
    descr VARCHAR(500) NOT NULL,
    data_de_pagamento DATETIME,
    data_de_vencimento DATETIME,
    tipo VARCHAR(24) NOT NULL
);
