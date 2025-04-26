/** @import { Lancamento } from './types.d.ts' */
import {
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
} from "./value.js";

import {
  LancamentoCriado,
  LancamentoPrevisto,
  LancamentoPago,
  LancamentoRecebido,
  LancamentoCancelado,
} from "./value.js";

export default {
  name: "Lancamento",
  data: (/** @type {Lancamento} */ lancamento) => ({
    async init() {
      this.natureza = {
        icone: this._iconeNatureza(),
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
    valores: {},
    natureza: {},
    /** @type {Lancamento} */
    value: Alpine.reactive({ ...lancamento }),
    _iconeNatureza() {
      switch (this.value.natureza) {
        case SALARIO:
          return "bi bi-cash-coin green";
        case BENEFICIOS:
          return "bi bi-cash green-light";
        case COMPRAS:
          return "bi bi-bag red-light";
        case MERCADO:
          return "bi bi-cart3 blue-light";
        case LUZ:
          return "bi bi-lightbulb-fill blue-light";
        case CONDOMINIO:
          return "bi bi-building-fill blue-light";
        case AGUA_E_GAS:
          return "bi bi-moisture blue-light";
        case TELEFONIA:
          return "bi bi-reception-4 blue-light";
        case NUVEM:
          return "bi bi-cloud-check-fill yellow-light";
        case INTERNET:
          return "bi bi-router-fill blue-light";
        case RECEITA_EXTRA:
          return "bi bi-cash green-light";
        case PETSHOP:
          return "bi bi-shop blue-light";
        case SAUDE:
          return "bi bi-bandaid-fill yellow-light";
        case INVESTIMENTO:
          return "bi bi-graph-up green-light";
        case OUTRO:
        default:
          return "bi bi-circle red-light";
      }
    },
  }),
};
