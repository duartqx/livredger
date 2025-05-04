import { locale } from "../utils/locale.js";

export class IntervaloPicker extends HTMLInputElement {
  constructor() {
    super();

    if (window.AirDatepicker) {
      this.picker = new window.AirDatepicker(this, {
        range: true,
        multipleDatesSeparator: " - ",
        locale: locale(),
        onSelect: ({ date, formattedDate, datepicker }) => {
          console.log(formattedDate)
          this.intervalo = {
            inicio: date[0] && dayjs(date[0]).format("YYYY-MM-DD") || "",
            final: date[1] && dayjs(date[1]).format("YYYY-MM-DD") || "",
          };
          console.log("Intervalo", this.intervalo)
          this.dispatchEvent(new Event("change", { bubbles: true }));
        },
      });
    }
  }
}
