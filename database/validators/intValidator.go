package validators

import "errors"

type IntValidator struct {
}

func (*IntValidator) Validate(val interface{}) error {
	if val == nil {
		return nil
	}

	if _, ok := val.(int); ok {
		return nil
	}

	return errors.New("invalid value type")
}
