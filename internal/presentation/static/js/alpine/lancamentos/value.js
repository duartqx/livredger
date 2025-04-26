/**
 * @import {
 * Naturezas,
 * MeiosTransacao,
 EventosLancamento
 * } from './types.d.ts'
 */

export const AGUA_E_GAS = "Água e Gás";
export const BENEFICIOS = "Benefícios";
export const COMPRAS = "Compras";
export const CONDOMINIO = "Condomínio";
export const INTERNET = "Internet";
export const INVESTIMENTO = "Investimento";
export const LUZ = "Luz";
export const MERCADO = "Mercado";
export const NUVEM = "Nuvem";
export const OUTRO = "Outro";
export const PETSHOP = "Petshop";
export const RECEITA_EXTRA = "Receita Extra";
export const SALARIO = "Salário";
export const SAUDE = "Saúde";
export const TELEFONIA = "Telefonia";

/** @type {Naturezas} */
export const NATUREZAS = [
  AGUA_E_GAS,
  BENEFICIOS,
  COMPRAS,
  CONDOMINIO,
  INTERNET,
  INVESTIMENTO,
  LUZ,
  MERCADO,
  NUVEM,
  OUTRO,
  PETSHOP,
  RECEITA_EXTRA,
  SALARIO,
  SAUDE,
  TELEFONIA,
];

export const CARTAO_DE_BENEFICIOS = "Cartão de Benefícios";
export const CARTAO_DE_CREDITO = "Cartão de Crédito";
export const CARTAO_DE_DEBITO = "Cartão de Débito";
export const DINHEIRO = "Dinheiro";
export const PIX = "PIX";
export const TRANFERENCIA_BANCARIA = "Transferência Bancária";

/** @type {MeiosTransacao} */
export const MEIOS_FINANCEIRO = [
  CARTAO_DE_BENEFICIOS,
  CARTAO_DE_CREDITO,
  CARTAO_DE_DEBITO,
  DINHEIRO,
  PIX,
  TRANFERENCIA_BANCARIA,
];

export const LancamentoCriado = "LancamentoCriado";
export const LancamentoPrevisto = "LancamentoPrevisto";
export const LancamentoPago = "LancamentoPago";
export const LancamentoRecebido = "LancamentoRecebido";
export const LancamentoCancelado = "LancamentoCancelado";

/** @type {EventosLancamento} */
export const LANCAMENTOS = [
  LancamentoCriado,
  LancamentoPrevisto,
  LancamentoPago,
  LancamentoRecebido,
  LancamentoRecebido,
  LancamentoCancelado,
];
