package types

type Comando interface {
	Validar() error
}
