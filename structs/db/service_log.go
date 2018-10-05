package db

import (
	"github.com/astaxie/beego/orm"
)

// ServiceLog - ServiceLogStruct
type ServiceLog struct {
	Type    string `orm:"column(type)"`
	Req     string `orm:"column(req)"`
	Res     string `orm:"column(res)"`
	Errcode string `orm:"column(errcode)"`
	JobID   string `orm:"column(job_id);pk"`
}

func init() {
	orm.RegisterModel(new(ServiceLog))
}
