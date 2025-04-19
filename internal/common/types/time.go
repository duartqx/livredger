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

func IntervaloDesseMes() Intervalo {
	now := time.Now()

	return Intervalo{
		Inicio: time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location()),
		Final:  now,
	}
}

func ParseIntervalo(inicio, final string) (*Intervalo, error) {
	var intervalo Intervalo

	if inicio != "" {
		if t, err := time.Parse(time.RFC3339, inicio); err == nil {
			intervalo.Inicio = t
		} else {
			return nil, fmt.Errorf("Inicio Inválido para Intervalo: %w", err)
		}
	}

	if final != "" {
		if t, err := time.Parse(time.RFC3339, final); err == nil {
			intervalo.Final = t
		} else {
			return nil, fmt.Errorf("Final Inválido para Intervalo: %w", err)
		}
	}

	if intervalo.Final.Before(intervalo.Inicio) {
		return nil, fmt.Errorf("Intervalo de tempo inválido (Início maior que final)")
	}

	return &intervalo, nil
}
