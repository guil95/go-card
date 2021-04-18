package handlers

import (
	"log"

	"github.com/go-playground/validator"
)

func isValidRequest(payload interface{}) bool {
	var validate *validator.Validate = validator.New()

	err := validate.Struct(payload)

	if err != nil {
		log.Println(err.Error())
		return false
	}

	return true
}
