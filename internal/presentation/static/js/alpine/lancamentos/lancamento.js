/** @import { Lancamento } from './types.d.ts' */
import {
  AGUA_E_GAS,
  BENEFICIOS,
  COMPRAS,
  CONDOMINIO,
  INTERNET,
  INVESTIMENTO,
  LANCAMENTOS_MAPEADOS_PARA_EXIBICAO,
  LUZ,
  MERCADO,
  NUVEM,
  OUTRO,
  PETSHOP,
  RECEITA_EXTRA,
  SALARIO,
  SAUDE,
  TELEFONIA,
} from "./value.js";

export default {
  name: "Lancamento",
  data: (/** @type {Lancamento} */ lancamento) => ({
    async init() {
      this.natureza = {
        ...this._iconeECorNatureza(),
        exibicao: this.value.natureza,
      };

      this.valores = {
        cor: this.value.valores <= 0 ? "normal" : "green",
        exibicao:
          this.value.valores < 0
            ? `-R$ ${Math.abs(this.value.valores)}`
            : `R$ ${this.value.valores}`,
      };
    },
    natureza: {},
    valores: {},
    /** @type {Lancamento} */
    value: Alpine.reactive({ ...lancamento }),
    evento() {
      return LANCAMENTOS_MAPEADOS_PARA_EXIBICAO[this.value.evento];
    },
    /** @returns {{ icone: string, cor: string }} */
    _iconeECorNatureza() {
      switch (this.value.natureza) {
        case SALARIO:
          return { icone: "bi bi-cash-coin", cor: "green" };
        case BENEFICIOS:
          return { icone: "bi bi-cash", cor: "green-light" };
        case COMPRAS:
          return { icone: "bi bi-bag", cor: "red-light" };
        case MERCADO:
          return { icone: "bi bi-cart3", cor: "blue-light" };
        case LUZ:
          return { icone: "bi bi-lightbulb-fill", cor: "blue-light" };
        case CONDOMINIO:
          return { icone: "bi bi-building-fill", cor: "blue-light" };
        case AGUA_E_GAS:
          return { icone: "bi bi-moisture", cor: "blue-light" };
        case TELEFONIA:
          return { icone: "bi bi-reception-4", cor: "blue-light" };
        case NUVEM:
          return { icone: "bi bi-cloud-check-fill", cor: "yellow-light" };
        case INTERNET:
          return { icone: "bi bi-router-fill", cor: "blue-light" };
        case RECEITA_EXTRA:
          return { icone: "bi bi-cash", cor: "green-light" };
        case PETSHOP:
          return { icone: "bi bi-shop", cor: "blue-light" };
        case SAUDE:
          return { icone: "bi bi-bandaid-fill", cor: "yellow-light" };
        case INVESTIMENTO:
          return { icone: "bi bi-graph-up", cor: "green-light" };
        case OUTRO:
        default:
          return { icone: "bi bi-circle", cor: "red-light" };
      }
    },
  }),
};
