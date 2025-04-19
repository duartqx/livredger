/**
 * @import {
 *    UUID,
 *    Float,
 *    Natureza,
 *    Meio,
 *    EventoLancamento,
 *    LancamentoApi,
 *    Lancamento
 *  } from './types.d.ts'
 */

export default {
  name: "lancamentos",
  component: () => ({
    async lancamentos() {
      if (this._lancamentos === null) {
        this._lancamentos = await this._obtemLancamentosViaAPI();
      }
      return this._lancamentos;
    },
    /** @type {LancamentoApi[] | null} */
    _lancamentos: null,
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

      console.log("GET", corpo.itens)

      return corpo.itens;
    },
  }),
};
