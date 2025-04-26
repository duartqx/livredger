/**
 * @import {
 *  Float,
 *  Natureza,
 *  EventoLancamento,
 *  Lancamento,
 *  LancamentoComando,
 *  MeioFinanceiro,
 LancamentoApi
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

import { LancamentoComandoFalho, LancamentoComandoSucesso } from "./eventos.js";
import { MEIOS_FINANCEIRO, NATUREZAS } from "./value.js";

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
  data: () => ({
    criando: false,
    /** @type {LancamentoModelo} */
    modelo: modeloPadrao(),
    naturezas: NATUREZAS,
    meios_financeiros: MEIOS_FINANCEIRO,
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

      /** @type {{ resultado: LancamentoApi } | { error: string }} */
      const body = await response.json();

      if (body.error) {
        window.dispatchEvent(
          new CustomEvent(LancamentoComandoFalho, {
            detail: { motivo: body.error },
          }),
        );
      } else if (body.resultado) {

        /** @type { LancamentoApi } */
        const resultado = body.resultado;

        /** @type { Lancamento } */
        const lancamento = {
          ...resultado,
          timestamp: dayjs(resultado.timestamp).toDate(),
          vencimento: dayjs(resultado.vencimento).toDate(),
        }

        window.dispatchEvent(
          new CustomEvent(LancamentoComandoSucesso, {
            detail: { lancamento: lancamento },
          }),
        );

        this.modelo = modeloPadrao();
      }

      this.criando = !criando;
    },
  }),
};
