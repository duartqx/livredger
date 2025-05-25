package decoders

import (
	"github.com/go-playground/form"
	"github.com/google/uuid"
)

var decoder *form.Decoder

func Decoder() *form.Decoder {
	if decoder != nil {
		return decoder
	}

	decoder = form.NewDecoder()

	decoder.RegisterCustomTypeFunc(func(vals []string) (interface{}, error) {
		if len(vals) == 0 || vals[0] == "" {
			return uuid.Nil, nil
		}
		return uuid.Parse(vals[0])
	}, uuid.UUID{})

	return decoder
}
