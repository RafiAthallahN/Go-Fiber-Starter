/*
This where everything about validation lies (include the user validation)
*/

package utils

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func ParseAndValidate[T any](c *fiber.Ctx, data []byte) (T, *Response) {
	// Create a new validator instance
	validate := validator.New()

	var request T

	if errParse := json.Unmarshal(data, &request); errParse != nil {
		return request, &Response{
			StatusCode: fiber.StatusBadRequest,
			Message:    "Invalid data sent. Please Check your data and try again",
			Error:      "Failed to parse data to struct",
		}
	}

	// Validate the request struct
	if errValidate := validate.Struct(request); errValidate != nil {
		// Handle validation errors
		var errors []string
		if _, ok := errValidate.(*validator.InvalidValidationError); ok {
			return request, &Response{
				StatusCode: fiber.StatusBadRequest,
				Message:    "Invalid data sent. Please check your data and try again",
				Error:      "Validation Error: " + errValidate.Error(),
			}
		}

		for _, err := range errValidate.(validator.ValidationErrors) {
			errors = append(errors, fmt.Sprintf("%s is %s", err.Field(), err.Tag()))
		}

		return request, &Response{
			StatusCode: fiber.StatusBadRequest,
			Message:    "Fill the required fields",
			Error:      strings.Join(errors, ", "),
		}
	}

	return request, nil
}
