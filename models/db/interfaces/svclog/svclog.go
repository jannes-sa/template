package svclog

import (
	dbStruct "template/structs/db"

	"github.com/jinzhu/gorm"
)

// ISvcLog - svcLog Logic Interface
type ISvcLog interface {
	GetAllServiceLog() ([]dbStruct.ServiceLog, error)
	GetOneByJobIDServiceLog(dbStruct.ServiceLog) (dbStruct.ServiceLog, error)
	InsertServiceLog(*gorm.DB, interface{}) (int64, error)
	UpdateByJobIDServiceLog(*gorm.DB, dbStruct.ServiceLog) error
	UpdateReturnByJobIDServiceLog(*gorm.DB, dbStruct.ServiceLog) ([]dbStruct.ServiceLog, error)
	DeleteByJobIDServiceLog(*gorm.DB, dbStruct.ServiceLog) error
}
