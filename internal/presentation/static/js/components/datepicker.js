import { locale } from "../utils/locale.js";

/**
 * @typedef {Object} IntervaloRef
 * @property {HTMLInputElement} inicio
 * @property {HTMLInputElement} final
 * @property {(chave: string, valor: string | Date | undefined) => void} atualiza
 */

export class InputIntervalo extends HTMLElement {
  constructor() {
    super();

    /** @type {HTMLInputElement} */
    this.input = document.createElement("input");
    this.input.setAttribute("readonly", true);

    /** @type {IntervaloRef} */
    this.intervalo = this.#montarIntervalo();

    this.append(this.input, this.intervalo.inicio, this.intervalo.final);

    this.removeAttribute("name");

    if (window.AirDatepicker) {
      this.picker = new window.AirDatepicker(this.input, {
        range: true,
        multipleDatesSeparator: " - ",
        locale: locale(),
        position: "bottom center",
        offset: -1,
        onSelect: ({ date, formattedDate, datepicker }) => {
          this.intervalo.atualiza("inicio", date[0]);
          this.intervalo.atualiza("final", date[1]);
        },
      });
    }
  }
  /** @returns {IntervaloRef} */
  #montarIntervalo() {
    const criarInnerInputs = (nome) => {
      const input = document.createElement("input");

      input.setAttribute("hidden", true);
      input.setAttribute("name", `${this.getAttribute("name")}[${nome}]`);

      return input;
    };

    return {
      inicio: criarInnerInputs("inicio"),
      final: criarInnerInputs("final"),
      atualiza: function(chave, valor) {
        this[chave].setAttribute(
          "value",
          (valor && dayjs(valor).format("YYYY-MM-DD")) || "",
        );
      },
    };
  }
}
