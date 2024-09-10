package validators

import "errors"

type IntMinMaxValidator struct {
	Min int
	Max int
}

func (i *IntMinMaxValidator) Validate(val interface{}) error {
	if val == nil {
		return nil
	}

	v := val.(int)

	if v < i.Min {
		return errors.New("value is less then specefied")
	}

	if v > i.Max {
		return errors.New("value is greater then specefied")
	}

	return nil
}
