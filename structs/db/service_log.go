package db

import "template/helper/constant/tablename"

// ServiceLog - ServiceLogStruct
type ServiceLog struct {
	JobID   string `gorm:"primary_key;column:job_id;not null`
	Type    string `gorm:"column:type;not null`
	Req     string `gorm:"column:req;not null`
	Res     string `gorm:"column:res;not null`
	Errcode string `gorm:"column:errcode;not null`
}

// TableName - TableName
func (ServiceLog) TableName() string {
	return tablename.ServiceLog
}
