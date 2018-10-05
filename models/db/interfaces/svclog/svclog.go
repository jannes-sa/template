package svclog

import (
	dbStruct "template/structs/db"

	"github.com/astaxie/beego/orm"
)

// ISvcLog - svcLog Logic Interface
type ISvcLog interface {
	GetAllServiceLog() ([]dbStruct.ServiceLog, error)
	GetOneByJobIDServiceLog(dbStruct.ServiceLog) (dbStruct.ServiceLog, error)
	InsertServiceLog(orm.Ormer, interface{}) (int64, error)
	UpdateByJobIDServiceLog(orm.Ormer, dbStruct.ServiceLog) error
	UpdateReturnByJobIDServiceLog(orm.Ormer, dbStruct.ServiceLog) ([]dbStruct.ServiceLog, error)
	DeleteByJobIDServiceLog(orm.Ormer, dbStruct.ServiceLog) error
}
