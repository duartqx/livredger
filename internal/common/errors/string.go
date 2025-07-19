package errors

func Stringer(err error) *string {
	if err == nil {
		return nil
	}

	s := err.Error()

	return &s
}
