package validation

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"mountainio/app/exception"
	"mountainio/domain/model"
)

func ValidateLoginAuth(params model.LoginInput) {
	err := validation.ValidateStruct(&params,
		validation.Field(&params.Email, validation.Required, is.Email),
		validation.Field(&params.Password, validation.Required),
	)

	if err != nil {
		panic(exception.ValidationError{
			Message: err.Error(),
		})
	}
}
