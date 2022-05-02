package validation

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"mountainio/app/exception"
	"mountainio/domain/model"
)

func ValidateRegisterUser(params model.RegisterUser) {
	err := validation.ValidateStruct(&params,
		validation.Field(&params.Name, validation.Required, validation.Length(4, 150)),
		validation.Field(&params.Email, validation.Required, is.Email),
		validation.Field(&params.Password, validation.Required),
	)

	if err != nil {
		panic(exception.ValidationError{
			Message: err.Error(),
		})
	}
}
