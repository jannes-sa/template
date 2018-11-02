package svclog

import (
	"template/helper/constant/tablename"
	dbStruct "template/structs/db"

	"github.com/jinzhu/gorm/dialects/postgres"

	"github.com/jinzhu/gorm"
)

// SvcLog - Logic Struct DB
type SvcLog struct{}

var tblServiceLog = tablename.ServiceLog

// GetAllServiceLog - GetAllServiceLog GetAll
func (d *SvcLog) GetAllServiceLog() (rows []dbStruct.ServiceLog, err error) {
	row := dbStruct.ServiceLog{
		JobID:   "job1",
		Req:     postgres.Jsonb{[]byte("{}")},
		Res:     postgres.Jsonb{[]byte("{}")},
		Errcode: "errcode",
		Type:    "http",
	}
	rows = append(rows, row)
	return
}

// GetOneByJobIDServiceLog - GetOneByJobIDServiceLog GetOne
func (d *SvcLog) GetOneByJobIDServiceLog(r dbStruct.ServiceLog) (row dbStruct.ServiceLog, err error) {
	row = dbStruct.ServiceLog{
		JobID:   "job1",
		Req:     postgres.Jsonb{[]byte("{}")},
		Res:     postgres.Jsonb{[]byte("{}")},
		Errcode: "errcode",
		Type:    "http",
	}
	return
}

// InsertServiceLog - InsertServiceLog Insert
func (d *SvcLog) InsertServiceLog(o *gorm.DB, v interface{}) (cnt int64, err error) {

	return
}

// UpdateByJobIDServiceLog - UpdateByJobIDServiceLog Update
func (d *SvcLog) UpdateByJobIDServiceLog(
	o *gorm.DB,
	row dbStruct.ServiceLog,
) (err error) {

	return
}

// UpdateReturnByJobIDServiceLog - UpdateReturnByJobIDServiceLog UpdateReturn
func (d *SvcLog) UpdateReturnByJobIDServiceLog(
	o *gorm.DB,
	row dbStruct.ServiceLog,
) (rows []dbStruct.ServiceLog, err error) {
	row = dbStruct.ServiceLog{
		JobID:   "job1",
		Req:     postgres.Jsonb{[]byte("{}")},
		Res:     postgres.Jsonb{[]byte("{}")},
		Errcode: "errcode",
		Type:    "http",
	}
	rows = append(rows, row)

	return
}

// DeleteByJobIDServiceLog - DeleteByJobIDServiceLog Delete
func (d *SvcLog) DeleteByJobIDServiceLog(
	o *gorm.DB,
	row dbStruct.ServiceLog,
) (err error) {

	return
}
