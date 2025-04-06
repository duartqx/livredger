package types

import (
	"fmt"
	"time"
)

type Intervalo struct {
	Inicio time.Time `json:"inicio"`
	Final  time.Time `json:"final"`
}

func ParseIntervalo(inicio, final string) (*Intervalo, error) {
	var intervalo Intervalo

	if inicio != "" {
		if t, err := time.Parse("2006-01-02", inicio); err == nil {
			intervalo.Inicio = t
		} else {
			return nil, fmt.Errorf("Inicio Inválido para Intervalo")
		}
	}

	if final != "" {
		if t, err := time.Parse("2006-01-02", final); err == nil {
			intervalo.Final = t
		} else {
			return nil, fmt.Errorf("Final Inválido para Intervalo")
		}
	}

	return &intervalo, nil
}
