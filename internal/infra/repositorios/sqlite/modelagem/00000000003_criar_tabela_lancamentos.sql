-- sql: Cria a tabela de lançamentos
CREATE TABLE IF NOT EXISTS lancamentos (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    evento VARCHAR(128) NOT NULL CHECK(
        evento IN (
            'LancamentoCriado',
            'LancamentoAtualizado',
            'LancamentoPago',
            'LancamentoRecebido',
            'LancamentoCancelado'
        )
    ),
    timestamp DATETIME DEFAULT (datetime('now')),

    chave VARCHAR(36) NOT NULL,
    versao INTEGER NOT NULL CHECK(versao > 0),

    valores REAL NOT NULL,
    natureza VARCHAR(128) NOT NULL REFERENCES naturezas(nome),
    meio VARCHAR(128) NOT NULL REFERENCES meios(nome),
    vencimento DATETIME,

    descr VARCHAR(500) NOT NULL CHECK(TRIM(descr) != ''),

    UNIQUE (chave, versao)
);

-- sql: Index em lancamentos.evento
CREATE INDEX IF NOT EXISTS idx_lancamentos_evento ON lancamentos (evento);

-- sql: Index em lancamentos.timestamp
CREATE INDEX IF NOT EXISTS idx_lancamentos_timestamp ON lancamentos (timestamp);

-- sql: Trigger que garante que versao != 1 deve já existir lançamentos com a mesma chave
CREATE TRIGGER IF NOT EXISTS tgr_lancamentos_bf_ins_versao_nao_igual_a_1_deve_existir_chave
BEFORE INSERT ON lancamentos
WHEN NEW.versao != 1 AND NOT EXISTS (SELECT 1 FROM lancamentos WHERE chave = NEW.chave)
BEGIN
    SELECT RAISE(ABORT, 'Não é permitido criar lançamentos com nova chave sem ser versao = 1');
END;

-- sql: Trigger que garante que versao = 1 deve ter nova chave
CREATE TRIGGER IF NOT EXISTS tgr_lancamentos_bf_ins_versao_1_deve_ter_nova_chave
BEFORE INSERT ON lancamentos
WHEN NEW.versao = 1 AND EXISTS (SELECT 1 FROM lancamentos WHERE chave = NEW.chave)
BEGIN
    SELECT RAISE(ABORT, 'Não é permitido criar lançamentos com versão = 1 se já existem lançamentos com a mesma chave');
END;

-- sql: Trigger que garante que versao > 1 já deve existir lançamentos com a chave e com versao = NEW.versao - 1
CREATE TRIGGER IF NOT EXISTS tgr_lancamentos_bf_ins_versao_maior_que_1_chave_deve_existir
BEFORE INSERT ON lancamentos
WHEN NEW.versao > 1 AND NOT EXISTS (SELECT 1 FROM lancamentos WHERE chave = NEW.chave AND versao = NEW.versao - 1)
BEGIN
    SELECT RAISE(ABORT, 'Não é permitido criar lançamentos com versão fora de ordem');
END;
