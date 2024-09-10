package validators

type IValidator interface {
	Validate(val interface{}) error
}
