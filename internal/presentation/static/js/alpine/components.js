import LancamentosConsulta from "./lancamentos/consulta.js";
import LancamentosComando from "./lancamentos/comando.js";
import Lancamento from "./lancamentos/lancamento.js";

document.addEventListener("alpine:init", () => {
  for (const component of [LancamentosConsulta, LancamentosComando, Lancamento]) {
    Alpine.data(component.name, component.data);
  }
});
