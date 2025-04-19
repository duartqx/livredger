import Lancamentos from "./lancamentos/lancamentos.js";

document.addEventListener("alpine:init", () => {
  Alpine.data(Lancamentos.name, Lancamentos.component);
});
