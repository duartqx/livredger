-- sql: Cria a tabela de lançamentos
CREATE TABLE IF NOT EXISTS lancamentos (
    id VARCHAR(36) PRIMARY KEY NOT NULL,
    evento VARCHAR(128) NOT NULL REFERENCES eventos_lancamento(nome),
    timestamp DATETIME DEFAULT (datetime('now')),

    chave VARCHAR(36) NOT NULL REFERENCES contas(chave),
    versao INTEGER NOT NULL CHECK(versao > 0),

    valores REAL NOT NULL, -- Valores desse lançamento específico
    totais REAL NOT NULL DEFAULT 0, -- Valores Pré Computados entre todos os lançamentos de mesma chave

    moeda VARCHAR(3) NOT NULL DEFAULT 'BRL' CHECK(moeda IN ('BRL', 'USD', 'EUR')),
    natureza VARCHAR(128) NOT NULL REFERENCES naturezas(nome),
    meio_financeiro VARCHAR(128) NOT NULL REFERENCES meios_financeiro(nome),
    vencimento DATETIME,

    descricao VARCHAR(500) NOT NULL CHECK(TRIM(descricao) != ''),

    UNIQUE (chave, versao)
);

-- sql: Index em lancamentos.evento
CREATE INDEX IF NOT EXISTS idx_lancamentos_evento ON lancamentos (evento);

-- sql: Index em lancamentos.timestamp
CREATE INDEX IF NOT EXISTS idx_lancamentos_timestamp ON lancamentos (timestamp);

-- sql: Index em lancamentos.chave
CREATE INDEX IF NOT EXISTS idx_lancamentos_chave ON lancamentos (chave);

-- sql: Index em lançamentos positivos
CREATE INDEX IF NOT EXISTS idx_lancamentos_valores_positivos
ON lancamentos (valores) WHERE valores >= 0;

-- sql: Index em lançamentos negativos
CREATE INDEX IF NOT EXISTS idx_lancamentos_valores_negativos
ON lancamentos (valores) WHERE valores < 0;

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
