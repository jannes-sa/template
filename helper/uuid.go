package helper

import (
	uuid "github.com/satori/go.uuid"
)

// GetUUID ...
func GetUUID() string {
	UUID, errUUID := uuid.NewV4()
	CheckErr("error uuid helper", errUUID)
	return UUID.String()
}
