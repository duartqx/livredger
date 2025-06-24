/**
 * @import {
 * Naturezas,
 * MeiosTransacao,
 * EventoLancamento,
 * EventosLancamento
 * } from './types.d.ts'
 */

export const AGUA_E_GAS = "Água e Gás";
export const BARBEIRO = "Barbeiro";
export const BENEFICIOS = "Benefícios";
export const COMPRAS = "Compras";
export const CONDOMINIO = "Condomínio";
export const EDUCACAO = "Educação";
export const ENTRETENIMENTO = "Entretenimento";
export const IMPOSTO = "Imposto";
export const INTERNET = "Internet";
export const INVESTIMENTO = "Investimento";
export const JUROS = "Juros";
export const LUZ = "Luz";
export const MERCADO = "Mercado";
export const NUVEM = "Nuvem";
export const OUTRO = "Outro";
export const PETSHOP = "Petshop";
export const RECEITA_EXTRA = "Receita Extra";
export const SALARIO = "Salário";
export const SAUDE = "Saúde";
export const TELEFONIA = "Telefonia";
export const TRABALHO = "Trabalho";
export const VAQUINHA = "Vaquinha";
export const VIAGEM = "Viagem";

/** @type {Naturezas} */
export const NATUREZAS = [
  AGUA_E_GAS,
  BENEFICIOS,
  COMPRAS,
  CONDOMINIO,
  EDUCACAO,
  ENTRETENIMENTO,
  IMPOSTO,
  INTERNET,
  INVESTIMENTO,
  JUROS,
  LUZ,
  MERCADO,
  NUVEM,
  OUTRO,
  PETSHOP,
  RECEITA_EXTRA,
  SALARIO,
  SAUDE,
  TELEFONIA,
  TRABALHO,
  VAQUINHA,
  VIAGEM,
];

export const BOLETO_BANCARIO = "Boleto Bancário"
export const CARTAO_DE_BENEFICIOS = "Cartão de Benefícios";
export const CARTAO_DE_CREDITO = "Cartão de Crédito";
export const CARTAO_DE_DEBITO = "Cartão de Débito";
export const DINHEIRO = "Dinheiro";
export const POUPANCA = "Poupança";
export const PIX = "PIX";
export const TRANFERENCIA_BANCARIA = "Transferência Bancária";
export const OUTRO_MEIO = "Outro"

/** @type {MeiosTransacao} */
export const MEIOS_FINANCEIRO = [
  BOLETO_BANCARIO,
  CARTAO_DE_BENEFICIOS,
  CARTAO_DE_CREDITO,
  CARTAO_DE_DEBITO,
  DINHEIRO,
  PIX,
  POUPANCA,
  TRANFERENCIA_BANCARIA,
  OUTRO_MEIO,
];

export const LancamentoPrevisto = "LancamentoPrevisto";
export const LancamentoPago = "LancamentoPago";
export const LancamentoRecebido = "LancamentoRecebido";
export const LancamentoCancelado = "LancamentoCancelado";
export const LancamentoCorrigido = "LancamentoCorrigido";
export const LancamentoTransferido = "LancamentoTransferido";

/** @type {EventosLancamento} */
export const LANCAMENTOS = [
  LancamentoPrevisto,
  LancamentoPago,
  LancamentoRecebido,
  LancamentoCancelado,
  LancamentoCorrigido,
  LancamentoTransferido,
];

/** @type { [ key: EventoLancamento ]: string } */
export const LANCAMENTOS_MAPEADOS_PARA_EXIBICAO = {
  LancamentoPrevisto: "Previsto",
  LancamentoPago: "Pago",
  LancamentoRecebido: "Recebido",
  LancamentoCancelado: "Cancelado",
  LancamentoCorrigido: "Corrigido",
  LancamentoTransferido: "Transferido",
};

/** @type {{ label: string, value: EventoLancamento }[]} */
export const LANCAMENTOS_MAPEADOS_PARA_OPCOES = [
  {
    label: LANCAMENTOS_MAPEADOS_PARA_EXIBICAO[LancamentoPrevisto],
    value: LancamentoPrevisto,
  },
  {
    label: LANCAMENTOS_MAPEADOS_PARA_EXIBICAO[LancamentoPago],
    value: LancamentoPago,
  },
  {
    label: LANCAMENTOS_MAPEADOS_PARA_EXIBICAO[LancamentoRecebido],
    value: LancamentoRecebido,
  },
  {
    label: LANCAMENTOS_MAPEADOS_PARA_EXIBICAO[LancamentoCancelado],
    value: LancamentoCancelado,
  },
  {
    label: LANCAMENTOS_MAPEADOS_PARA_EXIBICAO[LancamentoCorrigido],
    value: LancamentoCorrigido,
  },
  {
    label: LANCAMENTOS_MAPEADOS_PARA_EXIBICAO[LancamentoTransferido],
    value: LancamentoTransferido,
  },
];
