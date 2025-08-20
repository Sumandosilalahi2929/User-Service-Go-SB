package apperror

import (
	"errors"
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

type ValidationResponse struct {
	Field   string `json:"field,omitempty"`
	Message string `json:"message,omitempty"`
}

var ErrValidator = map[string]string{}

func ErrValidationResponse(err error) []ValidationResponse {
	var fieldErrors validator.ValidationErrors
	var validationResponse []ValidationResponse

	if errors.As(err, &fieldErrors) {
		for _, e := range fieldErrors {
			switch e.Tag() {
			case "required":
				validationResponse = append(validationResponse, ValidationResponse{
					Field:   e.Field(),
					Message: "this field is required",
				})
			case "email":
				validationResponse = append(validationResponse, ValidationResponse{
					Field:   e.Field(),
					Message: "invalid email format",
				})
			default:
				if msgTemplate, ok := ErrValidator[e.Tag()]; ok {
					count := strings.Count(msgTemplate, "%s")
					if count == 1 {
						validationResponse = append(validationResponse, ValidationResponse{
							Field:   e.Field(),
							Message: fmt.Sprintf(msgTemplate, e.Field()),
						})
					} else {
						validationResponse = append(validationResponse, ValidationResponse{
							Field:   e.Field(),
							Message: fmt.Sprintf(msgTemplate, e.Field(), e.Param()),
						})
					}
				} else {
					validationResponse = append(validationResponse, ValidationResponse{
						Field:   e.Field(),
						Message: fmt.Sprintf("something wrong on %s, %s", e.Field(), e.Tag()),
					})
				}
			}
		}
	}
	return validationResponse
}

func WrapError(err error) error {
	logrus.Errorf("error: %v", err)
	return err
}
