/**
 * @import {
 *    UUID,
 *    Float,
 *    Natureza,
 *    Meio,
 *    EventoLancamento,
 *    LancamentoApi,
 *    Lancamento,
 *    MeiosTransacao,
 *    Naturezas
 *  } from './types.d.ts'
 */

export default {
  name: "ConsultaLancamentos",
  component: () => ({
    /** @type {LancamentoApi[] | null} */
    _lancamentos: null,
    async lancamentos() {
      if (this._lancamentos === null) {
        this._lancamentos = await this._obtemLancamentosViaAPI();
      }
      return this._lancamentos;
    },
    /** @returns {Promise<LancamentoApi[]>} */
    async _obtemLancamentosViaAPI() {
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
  }),
};
