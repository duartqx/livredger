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

export type Meio =
  | "Transferência Bancária"
  | "PIX"
  | "Cartão de Crédito"
  | "Cartão de Débito"
  | "Dinheiro"
  | "Cartão de Benefícios";

export type EventoLancamento =
  | "LancamentoCriado"
  | "LancamentoPrevisto"
  | "LancamentoAtualizado"
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
  meio: string;
  vencimento: string;
  description: string;
};

export type Lancamento = {
  id: number;
  evento: EventoLancamento;
  timestamp: Date;
  chave: UUID;
  versao: Number;
  valores: Float;
  natureza: Natureza;
  meio: Meio;
  vencimento: Date;
  description: string;
};
