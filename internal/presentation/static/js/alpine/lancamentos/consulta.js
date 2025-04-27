import { LANCAMENTOS_MAPEADOS_PARA_OPCOES } from "./value.js";

export default {
  name: "ConsultaLancamentos",
  data: () => ({
    eventosLancamento: LANCAMENTOS_MAPEADOS_PARA_OPCOES,
  }),
}
