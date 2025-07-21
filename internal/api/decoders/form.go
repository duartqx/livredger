package decoders

import (
	"errors"
	"net/url"

	"github.com/go-playground/form"
	"github.com/google/uuid"
)

var (
	FormDecoderError = errors.New("DecoderError")
)

type formDecoder struct {
	wrapped *form.Decoder
}

func (d formDecoder) Decode(v interface{}, values url.Values) error {
	err := d.wrapped.Decode(v, values)

	switch {
	case err == nil:
		return nil
	case errors.Is(err, &form.InvalidDecoderError{}):
		return err
	default:
		return errors.Join(FormDecoderError, err)

	}
}

func NewFormDecoder() *formDecoder {
	dec := &formDecoder{wrapped: form.NewDecoder()}

	dec.wrapped.RegisterCustomTypeFunc(func(vals []string) (any, error) {
		if len(vals) == 0 || vals[0] == "" {
			return uuid.Nil, nil
		}
		return uuid.Parse(vals[0])
	}, uuid.UUID{})

	return dec
}

func ParseFormDecodeError(err error) map[string]string {
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
