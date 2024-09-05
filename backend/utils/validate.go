/*
This where everything about validation lies (include the user validation)
*/

package utils

import (
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func ParseAndValidate(c *fiber.Ctx, data any) error {

	validate := validator.New()

	if errParse := c.BodyParser(&data); errParse != nil {
		return c.Status(fiber.StatusBadRequest).JSON(Response{
			StatusCode: fiber.StatusBadRequest,
			Message:    "Invalid data sent. Please Check your data and try again",
			Error:      "Failed to parse data to struct",
		})
	}

	if errValidate := validate.Struct(data); errValidate != nil {
		var errors []string

		if _, ok := errValidate.(*validator.InvalidValidationError); ok {
			return c.Status(fiber.StatusBadRequest).JSON(Response{
				StatusCode: fiber.StatusBadRequest,
				Message:    "Invalid data sent. Please Check your data and try again",
				Error:      "Validation Error : " + strings.Join([]string{errValidate.Error()}, ","),
			})
		}

		for _, err := range errValidate.(validator.ValidationErrors) {
			errors = append(errors, err.Field())
		}

		return c.Status(fiber.StatusBadRequest).JSON(Response{
			StatusCode: fiber.StatusBadRequest,
			Message:    "Fill The Required Fields",
			Error:      errors,
		})
	}

	return nil
}
