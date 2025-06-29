export type UUID = string;

export type Float = number;

export type Natureza =
  | "Benefícios"
  | "Compras"
  | "Condomínio"
  | "Internet"
  | "Investimento"
  | "Luz"
  | "Mercado"
  | "Nuvem"
  | "Outro"
  | "Petshop"
  | "Receita Extra"
  | "Salário"
  | "Saúde"
  | "Telefonia"
  | "Água e Gás";

export type MeioFinanceiro =
  | "Cartão de Benefícios"
  | "Cartão de Crédito"
  | "Cartão de Débito"
  | "Dinheiro"
  | "PIX"
  | "Transferência Bancária";

export type EventoLancamento =
  | "LancamentoPrevisto"
  | "LancamentoPago"
  | "LancamentoRecebido"
  | "LancamentoTransferido"
  | "LancamentoCorrigido"
  | "LancamentoCancelado";

export type LancamentoApi = {
  id: string;
  evento: EventoLancamento;
  timestamp: string;
  chave: string;
  versao: number;
  valores: number;
  totais: number;
  natureza: string;
  meio_financeiro: string;
  vencimento: string;
  descricao: string;
};

export type Lancamento = {
  id: UUID;
  evento: EventoLancamento;
  timestamp: Date;
  chave: UUID;
  versao: Number;
  valores: Float;
  totais: Float;
  natureza: Natureza;
  meio_financeiro: MeioFinanceiro;
  vencimento: Date;
  descricao: string;
};

export type LancamentoComando = {
  evento: EventoLancamento;
  timestamp: Date;
  chave: UUID;
  versao: Number;
  valores: Float;
  natureza: Natureza;
  meio_financeiro: MeioFinanceiro;
  vencimento: Date;
  descricao: string;
};

export type Opcao = {
  label: string;
  value: string;
};

export type EventosLancamento = EventoLancamento[];

export type Naturezas = Natureza[];

export type MeiosFinanceiros = MeioFinanceiro[];
