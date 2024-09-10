package database

import "db/database/validators"

type ColumnType int

var ColumnTypeEnum = struct {
	Int    ColumnType
	String ColumnType
}{
	Int:    0,
	String: 1,
}

type Column struct {
	name       string
	colType    ColumnType
	validators []validators.IValidator
}

func (c Column) AddValidator(validator validators.IValidator) Column {
	c.validators = append(c.validators, validator)

	return c
}

func (c *Column) Validate(val interface{}) error {
	for _, validator := range c.validators {
		if err := validator.Validate(val); err != nil {
			return err
		}
	}
	return nil
}

func NewIntColumn(name string) Column {
	return Column{
		name:       name,
		colType:    ColumnTypeEnum.Int,
		validators: []validators.IValidator{&validators.IntValidator{}},
	}
}

func NewStringColumn(name string) Column {
	return Column{
		name:       name,
		colType:    ColumnTypeEnum.String,
		validators: []validators.IValidator{&validators.StringValidator{}},
	}
}
