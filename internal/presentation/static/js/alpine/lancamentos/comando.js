/**
 * @import {
 *  Float,
 *  Natureza,
 *  EventoLancamento,
 *  LancamentoComando,
 *  MeioTransacao
 * } from './types.d.ts'
 *
 * @typedef {{
 *  evento: EventoLancamento
 *  natureza: Natureza | ""
 *  meioTransacao: MeioTransacao | ""
 *  valores: Float
 *  vencimento: string
 *  descricao: string
 *  versao: number
 *  comando: () => LancamentoComando
 * }} LancamentoModelo
 */

import { MEIOS_TRANSACAO, NATUREZAS } from "./value.js";

export default {
  name: "CriarLancamento",
  component: () => ({
    /** @type {LancamentoModelo} */
    modelo: {
      evento: "LancamentoPrevisto",
      natureza: "",
      meioTransacao: "",
      valores: 0,
      vencimento: "",
      descricao: "",
      versao: 1,
      comando() {
        return {
          evento: this.evento,
          timestamp: window.dayjs(),
          chave: crypto.randomUUID(),
          versao: this.versao,
          valores: Number(this.valores),
          natureza: this.natureza,
          meio_transacao: this.meioTransacao,
          vencimento: window.dayjs(this.vencimento),
          descricao: this.descricao,
        };
      },
    },
    naturezas: NATUREZAS,
    meios_transacao: MEIOS_TRANSACAO,
    async submit() {
      console.log(this.modelo.comando());
    },
  }),
};
