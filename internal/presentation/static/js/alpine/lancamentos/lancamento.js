/** @import { Lancamento } from './types.d.ts' */
import {
  AGUA_E_GAS,
  BENEFICIOS,
  COMPRAS,
  CONDOMINIO,
  EDUCACAO,
  ENTRETENIMENTO,
  IMPOSTO,
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
  TRABALHO,
  VAQUINHA,
  VIAGEM,
} from "./value.js";

/**
 * @param {number} valor
 * @returns {{cor: string, exibicao: string}}
 */
function exibicaoValores(valor) {
  return {
    cor: valor <= 0 ? "normal" : "green",
    exibicao:
      valor < 0
        ? `-R$ ${Math.abs(valor).toFixed(2)}`
        : `R$ ${valor.toFixed(2)}`,
  };
}

export default {
  name: "Lancamento",
  data: (/** @type {Lancamento} */ lancamento) => ({
    /** @type {Lancamento} */
    lancamento: Alpine.reactive({ ...lancamento }),
    get natureza() {
      return {
        ...this._iconeECorNatureza(),
        exibicao: this.lancamento.natureza,
      };
    },
    get valores() {
      return exibicaoValores(this.lancamento.valores);
    },
    get totais() {
      return exibicaoValores(this.lancamento.totais);
    },
    get timestamp() {
      return dayjs
        .utc(this.lancamento.timestamp)
        .tz("America/Sao_Paulo")
        .format("LLLL");
    },
    get evento() {
      return LANCAMENTOS_MAPEADOS_PARA_EXIBICAO[this.lancamento.evento];
    },
    /** @returns {{ icone: string, cor: string }} */
    _iconeECorNatureza() {
      switch (this.lancamento.natureza) {
        case SALARIO:
          return { icone: "bi bi-cash-coin", cor: "green" };
        case EDUCACAO:
          return { icone: "bi bi-book", cor: "green-light" };
        case ENTRETENIMENTO:
          return { icone: "bi bi-camera-reels-fill", cor: "green-light" };
        case IMPOSTO:
          return { icone: "bi bi-bank2", cor: "red-light" };
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
        case TRABALHO:
          return { icone: "bi bi-pc-display-horizontal", cor: "blue-light" };
        case VIAGEM:
          return { icone: "bi bi-luggage-fill", cor: "blue-light" };
        case VAQUINHA:
        case OUTRO:
        default:
          return { icone: "bi bi-circle", cor: "red-light" };
      }
    },
  }),
};
