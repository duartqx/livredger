/** @import { LancamentoApi } from './types.d.ts' */

import { LancamentoComandoSucesso } from './eventos.js';

export default {
  name: "ConsultaLancamentos",
  data: () => ({
    async init() {
      this.lancamentos = await this._consultarLancamentos();

      /** Reatividade */
      window.addEventListener(
        LancamentoComandoSucesso,
        this._adicionaLancamentoRecemCriado.bind(this),
      );
    },
    /** @type {LancamentoApi[]} */
    lancamentos: [],
    /** @returns {Promise<LancamentoApi[]>} */
    async _consultarLancamentos() {
      const response = await fetch(
        `/api/lancamentos?q={"somente_versao_mais_recente": true}`,
        { method: "GET", credentials: "include" },
      );

      if (!response.ok) {
        throw new Error(
          `Não foi possível obter lista de lançamentos: ${response.statusText}`,
        );
      }

      /** @type {{total: number, itens: LancamentoApi[]}} */
      const corpo = await response.json();

      console.log("GET", corpo.itens);

      return corpo.itens;
    },
    _adicionaLancamentoRecemCriado(event) {
      this.lancamentos.unshift(event.detail.lancamento);
    },
  }),
};
