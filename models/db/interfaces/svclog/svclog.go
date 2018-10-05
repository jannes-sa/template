package svclog

import (
	dbStruct "template/structs/db"

	"github.com/astaxie/beego/orm"
)

// ISvcLog - svcLog Logic Interface
type ISvcLog interface {
	GetAllServiceLog() ([]dbStruct.ServiceLog, error)
	GetOneServiceLog(cond string) (dbStruct.ServiceLog, error)
	InsertServiceLog(orm.Ormer, interface{}) (int64, error)
	UpdateServiceLog(orm.Ormer, []string, string) error
	UpdateReturnServiceLog(orm.Ormer, []string, string, []string) ([]dbStruct.ServiceLog, error)
	DeleteServiceLog(orm.Ormer, string) error
}
