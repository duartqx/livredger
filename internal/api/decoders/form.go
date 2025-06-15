package decoders

import (
	"errors"
	"fmt"

	"github.com/duartqx/livredger/internal/common/types"
	"github.com/go-playground/form"
	"github.com/google/uuid"
)

var (
	DecoderError = fmt.Errorf("%w: DecoderError", types.BusinessLogicError)
	decoder      *form.Decoder
)

func Decoder() *form.Decoder {
	if decoder != nil {
		return decoder
	}

	decoder = form.NewDecoder()

	decoder.RegisterCustomTypeFunc(func(vals []string) (any, error) {
		if len(vals) == 0 || vals[0] == "" {
			return uuid.Nil, nil
		}
		return uuid.Parse(vals[0])
	}, uuid.UUID{})

	return decoder
}

func ParseDecodeError(err error) map[string]string {
	out := make(map[string]string)

	if err == nil {
		return out
	}

	var errs form.DecodeErrors

	if !errors.As(err, &errs) {
		return out
	}

	for key, err := range errs {
		out[key] = err.Error()
	}

	return out
}
