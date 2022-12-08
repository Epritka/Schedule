package validator

import (
	"fmt"

	pgvalidator "github.com/go-playground/validator"
	"github.com/iancoleman/strcase"
)

type Validator interface {
	Struct(s interface{}) error
	GetErrors(err error) (errors map[string]string)
	AddCustomValidators(name string, fn pgvalidator.Func) error
}

type validator struct {
	validate *pgvalidator.Validate
}

func New() Validator {
	v := &validator{
		validate: pgvalidator.New(),
	}
	return v
}

func (v *validator) AddCustomValidators(name string, fn pgvalidator.Func) error {
	err := v.validate.RegisterValidation(name, fn)
	return err
}

func (v *validator) Struct(s interface{}) error {
	return v.validate.Struct(s)
}

func (v *validator) GetErrors(err error) (errors map[string]string) {
	var message string
	var field string
	errors = map[string]string{}

	for _, validErr := range err.(pgvalidator.ValidationErrors) {
		field = strcase.ToLowerCamel(validErr.Field())

		switch validErr.Tag() {
		default:
			message = fmt.Sprint(validErr)
		}

		if _, ok := errors[field]; ok {
			errors[field] += fmt.Sprintf(
				" %s", message,
			)
		}

		errors[field] = message
	}

	return errors
}
