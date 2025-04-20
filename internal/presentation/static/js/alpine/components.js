import LancamentosConsulta from "./lancamentos/consulta.js";
import LancamentosComando from "./lancamentos/comando.js";

document.addEventListener("alpine:init", () => {
  Alpine.data(LancamentosConsulta.name, LancamentosConsulta.component);
  Alpine.data(LancamentosComando.name, LancamentosComando.component);
});
