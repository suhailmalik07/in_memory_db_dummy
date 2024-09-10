package validators

import (
	"errors"
)

type RequiredValidator struct {
}

func (*RequiredValidator) Validate(val interface{}) error {
	if val == nil {
		return errors.New("value is required")
	}

	return nil
}
