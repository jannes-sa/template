package db

import (
	"template/helper/constant/tablename"

	"github.com/jinzhu/gorm/dialects/postgres"
)

// ServiceLog - ServiceLogStruct
type ServiceLog struct {
	JobID   string         `gorm:"primary_key;column:job_id;not null`
	Type    string         `gorm:"column:type;not null`
	Req     postgres.Jsonb `gorm:"column:req;not null`
	Res     postgres.Jsonb `gorm:"column:res;not null`
	Errcode string         `gorm:"column:errcode;not null`
}

// TableName - TableName
func (ServiceLog) TableName() string {
	return tablename.ServiceLog
}
