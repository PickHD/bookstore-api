package utils

import (
	"github.com/twinj/uuid"
)

type ErrorValidateResponse struct {
	FailedField string
	Tag         string
	Value       string
}

func ValidateUuid(data string) bool {
	_, err := uuid.Parse(data)

	return err == nil
}
