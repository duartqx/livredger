/** @import { LancamentoApi } from './types.d.ts' */

import { LancamentoComandoSucesso } from "./eventos.js";
import { LANCAMENTOS_MAPEADOS_PARA_OPCOES } from "./value.js";

export default {
  name: "ListarLancamentos",
  data: () => ({
    async init() {
      this.lancamentos = await this._consultarLancamentos();

      /** Reatividade */
      window.addEventListener(
        LancamentoComandoSucesso,
        this._adicionaLancamentoRecemCriado.bind(this),
      );
    },
    filtros: {
      somente_versao_mais_recente: true,
      evento: "",
      intervalo: "",
      descricao: "",
      chave: "",
    },
    /** @type {LancamentoApi[]} */
    lancamentos: [],
    eventosLancamento: LANCAMENTOS_MAPEADOS_PARA_OPCOES,
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

      return corpo.itens.map((lancamento) => ({
        ...lancamento,
        timestamp: dayjs(lancamento.timestamp).toDate(),
        vencimento: dayjs(lancamento.vencimento).toDate(),
      }));
    },
    _adicionaLancamentoRecemCriado(event) {
      if (!event.detail.index) {
        return this.lancamentos.unshift(event.detail.lancamento);
      }

      this.lancamentos[event.detail.index] = event.detail.lancamento;

      return this.lancamentos.length;
    },
  }),
}
