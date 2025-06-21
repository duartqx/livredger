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
  }

  connectedCallback() {
    this.innerHTML = "";

    this.input = document.createElement("input");
    this.input.setAttribute("readonly", true);

    this.after = document.createElement("span");
    this.after.classList.add("bi", "bi-calendar2-week", "after");

    this.intervalo = this.#montarIntervalo();

    this.append(
      this.after,
      this.input,
      this.intervalo.inicio,
      this.intervalo.final,
    );

    if (!window.AirDatepicker) {
      return;
    }

    this.picker = new window.AirDatepicker(this.input, {
      container: this,
      range: true,
      multipleDatesSeparator: " - ",
      locale: locale(),
      position: "bottom center",
      offset: -1,
      onSelect: ({ date }) => {
        this.intervalo.atualiza("inicio", date[0]);
        this.intervalo.atualiza("final", date[1]);

        this.dispatchEvent(
          new CustomEvent("change", {
            bubbles: true,
            composed: true,
            detail: { value: { inicio: date[0], final: date[1] } },
          }),
        );
      },
    });
  }

  disconnectedCallback() {
    this.picker.destroy();
    this.picker = null;
  }

  /** @returns {IntervaloRef} */
  #montarIntervalo() {
    const criarInnerInputs = (nome) => {
      const input = document.createElement("input");

      input.setAttribute("hidden", true);
      input.setAttribute("name", `${this.getAttribute("name")}[${nome}]`);
      input.setAttribute("autocomplete", "off");

      return input;
    };

    return {
      inicio: criarInnerInputs("inicio"),
      final: criarInnerInputs("final"),
      atualiza: function (chave, valor) {
        this[chave].setAttribute(
          "value",
          (valor && dayjs(valor).format("YYYY-MM-DD")) || "",
        );
      },
    };
  }
}

export class InputDatetime extends HTMLElement {
  constructor() {
    super();
  }

  connectedCallback() {
    this.innerHTML = "";

    this.input = document.createElement("input");
    this.input.setAttribute("name", this.getAttribute("name"));
    this.input.setAttribute("autocomplete", "off");

    if (this.getAttribute("required")) {
      this.input.setAttribute("required", true);
    }

    const initialValue = [this.getAttribute("value")];

    this.after = document.createElement("span");
    this.after.classList.add("bi", "bi-calendar2-week", "after");

    this.removeAttribute("name");

    this.append(this.after, this.input);

    if (!window.AirDatepicker) {
      return;
    }

    this.picker = new window.AirDatepicker(this.input, {
      container: this,
      range: false,
      timepicker: true,
      locale: locale(),
      position: "bottom center",
      offset: -1,
      selectedDates: initialValue
        .filter((v) => v)
        .map((v) => dayjs(v).toDate()),
      onSelect: ({ date }) => {
        this.setAttribute("value", dayjs(date).format());
        this.dispatchEvent(
          new CustomEvent("change", {
            bubbles: true,
            composed: true,
            detail: { value: date },
          }),
        );
      },
    });
  }

  disconnectedCallback() {
    this.picker.destroy();
    this.picker = null;
  }
}
