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
 *  meio_financeiro: MeioFinanceiro | ""
 *  valores: Float
 *  vencimento: string
 *  descricao: string
 *  versao: number
 *  comando: () => LancamentoComando
 * }} LancamentoModelo
 */

import { LancamentoComandoFalho, LancamentoComandoSucesso } from "./eventos.js";
import {
  MEIOS_FINANCEIRO,
  NATUREZAS,
  LANCAMENTOS_MAPEADOS_PARA_OPCOES,
  LancamentoCancelado,
} from "./value.js";

/** @returns {LancamentoModelo} */
function modeloPadrao() {
  if (window.location.host.startsWith("localhost")) {
    return {
      evento: "LancamentoPrevisto",
      natureza: "Água e Gás",
      meio_financeiro: "Cartão de Benefícios",
      valores: 10.0,
      vencimento: dayjs().add(3, "hours").toISOString().slice(0, 16),
      descricao: "afksjdfl",
      versao: 0,
    };
  }
  return {
    evento: "LancamentoPrevisto",
    natureza: "",
    meio_financeiro: "",
    valores: 0,
    vencimento: "",
    descricao: "",
    versao: 0,
  };
}

export default {
  name: "CriarLancamento",
  data: (/** @type {LancamentoApi} */ lancamentoOriginal) => ({
    criando: false,
    /** @type {LancamentoModelo} */
    modelo: (() => {
      if (lancamentoOriginal) {
        console.log("lancamentoOriginal", lancamentoOriginal);
        return {
          ...lancamentoOriginal,
          vencimento: dayjs(lancamentoOriginal.vencimento)
            .add(3, "hours")
            .toISOString()
            .slice(0, 16),
        };
      }
      return {
        ...modeloPadrao(),
        versao: 1,
      };
    })(),
    naturezas: NATUREZAS,
    meios_financeiros: MEIOS_FINANCEIRO,
    eventosLancamento: LANCAMENTOS_MAPEADOS_PARA_OPCOES,
    reset(/** @type {LancamentoApi} */ lancamentoCriado) {
      const lancamento =
        lancamentoCriado || lancamentoOriginal || modeloPadrao();

      return {
        ...lancamento,
        vencimento: dayjs(lancamento.vencimento)
          .add(3, "hours")
          .toISOString()
          .slice(0, 16),
        versao: lancamento.versao + 1,
      };
    },
    /** @returns {LancamentoComando} */
    comando() {
      return {
        evento: this.modelo.evento,
        timestamp: window.dayjs(),
        chave: this.modelo.chave || crypto.randomUUID(),
        versao: this.modelo.versao,
        valores: Number(this.modelo.valores),
        natureza: this.modelo.natureza,
        meio_financeiro: this.modelo.meio_financeiro,
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
        };

        window.dispatchEvent(
          new CustomEvent(LancamentoComandoSucesso, {
            detail: { lancamento: lancamento },
          }),
        );

        this.modelo = this.reset(lancamento);
      }

      this.criando = !criando;
    },
  }),
};
