package db

import "github.com/astaxie/beego/orm"

// ServiceLog - ServiceLogStruct
type ServiceLog struct {
	Type    string `orm:"column(type);null"`
	Req     string `orm:"column(req);type(json);null"`
	Res     string `orm:"column(res);type(json);null"`
	Errcode string `orm:"column(errcode);type(json);null"`
	JobID   string `orm:"column(job_id);pk"`
}

func init() {
	orm.RegisterModel(new(ServiceLog))
}
