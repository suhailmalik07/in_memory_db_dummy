package validators

import "errors"

type StringMaxLengthValidator struct {
	MaxLength int
}

func (i *StringMaxLengthValidator) Validate(val interface{}) error {
	if val == nil {
		return nil
	}

	v := val.(string)

	if len(v) > i.MaxLength {
		return errors.New("value length is greater than specified")
	}

	return nil
}
