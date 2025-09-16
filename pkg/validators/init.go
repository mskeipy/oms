package validators

import (
	"context"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func SetupValidator() {
	var validatePassword validator.FuncCtx = func(ctx context.Context, fl validator.FieldLevel) bool {
		return ValidatePassword(fl)
	}

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidationCtx("secure_password", validatePassword, true)
	}
}
