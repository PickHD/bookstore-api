package utils

import "github.com/twinj/uuid"

func GenerateUuid() string {
	newUuid := uuid.NewV4()

	return newUuid.String()
}
