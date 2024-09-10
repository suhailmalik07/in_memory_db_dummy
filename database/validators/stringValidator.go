package validators

import "errors"

type StringValidator struct{}

func (*StringValidator) Validate(val interface{}) error {
	if val == nil {
		return nil
	}

	if _, ok := val.(string); ok {
		return nil
	}

	return errors.New("invalid value type")
}
