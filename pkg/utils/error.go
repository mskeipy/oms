package utils

import (
	"dropx/internal/api/dto/response"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

func ValidationErrorResponse(err error, statusCode int) response.CommonResponse {
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		var fieldErrors []response.FieldError
		for _, e := range validationErrors {
			fieldErrors = append(fieldErrors, response.FieldError{
				Field:   e.Field(),
				Message: e.Tag(),
			})
		}
		return response.CommonResponse{
			Error: &response.ErrorResponse{
				Code:   http.StatusUnprocessableEntity,
				Errors: fieldErrors,
			},
		}
	}

	return response.CommonResponse{
		Error: &response.ErrorResponse{
			Code:   statusCode,
			Errors: []response.FieldError{{Message: err.Error()}}, // Lưu lỗi dạng list
		},
	}
}

func WriteErrorResponse(c *gin.Context, err error, statusCode int) {
	resp := ValidationErrorResponse(err, statusCode)
	c.JSON(resp.Error.Code, resp)
	c.Abort()
}
