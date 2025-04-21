/**
 * @import {
 *  Float,
 *  Natureza,
 *  EventoLancamento,
 *  LancamentoComando,
 *  MeioFinanceiro
 * } from './types.d.ts'
 *
 * @typedef {{
 *  evento: EventoLancamento
 *  natureza: Natureza | ""
 *  meioFinanceiro: MeioFinanceiro | ""
 *  valores: Float
 *  vencimento: string
 *  descricao: string
 *  versao: number
 *  comando: () => LancamentoComando
 * }} LancamentoModelo
 */

import { MEIOS_FINANCEIRO as MEIOS_FINANCEIROS, NATUREZAS } from "./value.js";

/** @returns {LancamentoModelo} */
function modeloPadrao() {
  if (window.location.host.startsWith("localhost")) {
    return {
      evento: "LancamentoPrevisto",
      natureza: "Água e Gás",
      meioFinanceiro: "Cartão de Benefícios",
      valores: 10.0,
      vencimento: dayjs().add(3, "hours").toISOString().slice(0, 16),
      descricao: "afksjdfl",
      versao: 1,
    };
  }
  return {
    evento: "LancamentoCriado",
    natureza: "",
    meioFinanceiro: "",
    valores: 0,
    vencimento: "",
    descricao: "",
    versao: 1,
  };
}

export default {
  name: "CriarLancamento",
  component: () => ({
    criando: false,
    /** @type {LancamentoModelo} */
    modelo: modeloPadrao(),
    naturezas: NATUREZAS,
    meios_financeiros: MEIOS_FINANCEIROS,
    /** @returns {LancamentoComando} */
    comando() {
      return {
        evento: this.modelo.evento,
        timestamp: window.dayjs(),
        chave: crypto.randomUUID(),
        versao: this.modelo.versao,
        valores: Number(this.modelo.valores),
        natureza: this.modelo.natureza,
        meio_financeiro: this.modelo.meioFinanceiro,
        vencimento: window.dayjs(this.modelo.vencimento),
        descricao: this.modelo.descricao,
      };
    },
    async submit() {
      const criando = !this.criando;

      this.criando = criando;

      const comando = this.comando();

      const response = await fetch("/api/lancamentos", {
        method: "POST",
        credentials: "include",
        body: JSON.stringify(comando),
      });

      /** @type {{ resultado: Lancamento } | { error: string }} */
      const body = await response.json();

      if (body.resultado && !body.error) {
        window.dispatchEvent(
          new CustomEvent("LancamentoComandoSucesso", {
            detail: { lancamento: body.resultado },
          }),
        );

        this.modelo = modeloPadrao();
      } else {
        window.dispatchEvent(
          new CustomEvent("LancamentoComandoFalhou", {
            detail: { motivo: body.error },
          }),
        );
      }

      this.criando = !criando;
    },
  }),
};
