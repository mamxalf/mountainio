package validation

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"mountainio/app/exception"
	"mountainio/domain/model"
)

func ValidateRegisterUser(params model.RegisterUser) {
	err := validation.ValidateStruct(&params,
		validation.Field(&params.Name, validation.Required),
		validation.Field(&params.Email, validation.Required),
		validation.Field(&params.Password, validation.Required),
	)

	if err != nil {
		panic(exception.ValidationError{
			Message: err.Error(),
		})
	}
}
