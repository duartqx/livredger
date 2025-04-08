package types

import (
	"fmt"
	"time"
)

type Intervalo struct {
	Inicio time.Time `json:"inicio"`
	Final  time.Time `json:"final"`
}

func (i Intervalo) IsZero() bool {
	return i.Inicio.IsZero() && i.Final.IsZero()
}

func ParseIntervalo(inicio, final string) (*Intervalo, error) {
	var intervalo Intervalo

	if inicio != "" {
		if t, err := time.Parse(time.RFC3339, inicio); err == nil {
			intervalo.Inicio = t
		} else {
			return nil, fmt.Errorf("Inicio Inválido para Intervalo")
		}
	}

	if final != "" {
		if t, err := time.Parse(time.RFC3339, final); err == nil {
			intervalo.Final = t
		} else {
			return nil, fmt.Errorf("Final Inválido para Intervalo")
		}
	}

	return &intervalo, nil
}
