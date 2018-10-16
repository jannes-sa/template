package db

import (
	"github.com/astaxie/beego/orm"
)

// ServiceLog - ServiceLogStruct
type ServiceLog struct {
	Type    string `orm:"column(type)" bson:"type"`
	Req     string `orm:"column(req)" bson:"req"`
	Res     string `orm:"column(res)" bson:"res"`
	Errcode string `orm:"column(errcode)" bson:"errcode"`
	JobID   string `orm:"column(job_id);pk" bson:"job_id"`
}

func init() {
	orm.RegisterModel(new(ServiceLog))
}
