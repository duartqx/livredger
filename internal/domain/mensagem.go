package domain

type Evento string

type Mensagem string

type Comando interface {
	Validar() error
}
