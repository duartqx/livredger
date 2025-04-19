package sql

import (
	"bytes"
	"database/sql"
	"strings"
	"text/template"

	"github.com/google/uuid"
)

type Condicoes struct {
	op    string
	conds []string
	args  []any
}

func (w *Condicoes) Add(condicao string, args ...sql.NamedArg) {

	w.conds = append(w.conds, condicao)

	for _, arg := range args {
		w.args = append(w.args, arg)
	}
}

func (w Condicoes) Condicoes() string {
	return "WHERE " + strings.Join(w.conds, w.op)
}

func (w Condicoes) Argumentos() *[]any {
	return &w.args
}

type Valores map[string]string

type Builder struct {
	tmpl *template.Template
}

func NewBuilder(tmpl string) *Builder {
	return &Builder{
		tmpl: template.Must(
			template.
				New(uuid.New().String()).
				Parse(tmpl),
		),
	}
}

func (b Builder) Render(vals *Valores) string {
	var buf bytes.Buffer

	err := b.tmpl.Execute(&buf, vals)

	if err != nil {
		panic(err)
	}

	return buf.String()
}
