import { locale } from "../utils/locale.js";

/**
 * @typedef {Object} IntervaloRef
 * @property {HTMLInputElement} inicio
 * @property {HTMLInputElement} final
 * @property {(chave: string, valor: string | Date | undefined) => void} atualiza
 */

export class IntervaloPicker extends HTMLInputElement {
  constructor() {
    super();

    this.setAttribute("readonly", "true");

    /** @type {IntervaloRef} */
    this.intervalo = {
      inicio: this.#criarInnerInputDoIntervalo("inicio"),
      final: this.#criarInnerInputDoIntervalo("final"),
      atualiza: function(chave, valor) {
        this[chave].setAttribute(
          "value",
          (valor && dayjs(valor).format("YYYY-MM-DD")) || "",
        );
      },
    };

    this.removeAttribute("name");

    if (window.AirDatepicker) {
      this.picker = new window.AirDatepicker(this, {
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
  /** @returns {HTMLInputElement} */
  #criarInnerInputDoIntervalo(nome) {
    const membro = document.createElement("input");

    membro.setAttribute("hidden", true);
    membro.setAttribute("name", `${this.getAttribute("name")}[${nome}]`);

    this.append(membro);

    return membro;
  }
}
