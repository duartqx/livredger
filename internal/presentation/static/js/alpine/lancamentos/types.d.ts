export type UUID = string;

export type Float = number;

export type Natureza =
  | "Salário"
  | "Benefícios"
  | "Compras"
  | "Mercado"
  | "Luz"
  | "Condomínio"
  | "Água e Gás"
  | "Telefonia"
  | "Nuvem"
  | "Internet"
  | "Receita Extra"
  | "Petshop"
  | "Saúde"
  | "Investimento"
  | "Outro";

export type MeioFinanceiro =
  | "Transferência Bancária"
  | "PIX"
  | "Cartão de Crédito"
  | "Cartão de Débito"
  | "Dinheiro"
  | "Cartão de Benefícios";

export type EventoLancamento =
  | "LancamentoCriado"
  | "LancamentoPrevisto"
  | "LancamentoPago"
  | "LancamentoRecebido"
  | "LancamentoRecebido"
  | "LancamentoCancelado";

export type LancamentoApi = {
  id: number;
  evento: EventoLancamento;
  timestamp: string;
  chave: string;
  versao: number;
  valores: number;
  natureza: string;
  meio_financeiro: string;
  vencimento: string;
  descricao: string;
};

export type Lancamento = {
  id: number;
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
