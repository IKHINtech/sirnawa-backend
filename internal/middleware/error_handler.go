package middleware

import (
	"encoding/json"
	"reflect"

	"github.com/IKHINtech/sirnawa-backend/pkg/utils"
	"github.com/IKHINtech/sirnawa-backend/pkg/validators"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func ValidateRequest(reqBody any) fiber.Handler {
	return func(c *fiber.Ctx) error {
		errHandler := &utils.ResponseHandler{}

		// Decode request body into a map
		var bodyMap map[string]any
		if err := json.Unmarshal(c.Body(), &bodyMap); err != nil {
			return errHandler.BadRequest(c, []string{"Invalid request body"})
		}

		typ := reflect.TypeOf(reqBody).Elem()
		v := reflect.New(typ).Interface()

		if err := c.BodyParser(v); err != nil {
			return errHandler.BadRequest(c, []string{"Cannot parse JSON"})
		}

		// Create a set of allowed fields from the DTO
		allowedFields := make(map[string]struct{})
		value := reflect.ValueOf(reqBody).Elem()
		for i := range value.NumField() {
			allowedFields[value.Type().Field(i).Tag.Get("json")] = struct{}{}
		}

		// Check for unexpected fields
		for field := range bodyMap {
			if _, found := allowedFields[field]; !found {
				return errHandler.BadRequest(c, []string{"Unexpected field: " + field})
			}
		}

		if err := validators.ValidateStruct(v); err != nil {
			var errors []string
			for _, err := range err.(validator.ValidationErrors) {
				errors = append(errors, err.Field()+": "+err.Tag())
			}
			return errHandler.BadRequest(c, errors)
		}

		c.Locals("validatedReqBody", v)
		return c.Next()
	}
}
