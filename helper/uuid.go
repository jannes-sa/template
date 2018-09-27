package helper

import (
	uuid "github.com/satori/go.uuid"
)

// GetUUID - Generate UUID V4
func GetUUID() string {
	return uuid.NewV4().String()
}
