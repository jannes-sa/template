package svclog

import (
	"strings"
	"template/helper/constant"
	"template/helper/constant/tablename"
	db "template/models/db/pgsql2"
	dbStruct "template/structs/db"

	"github.com/jinzhu/gorm"
)

// SvcLog - Logic Struct DB
type SvcLog struct{}

var tblServiceLog = tablename.ServiceLog

// GetAllServiceLog - GetAllServiceLog  GetAll
func (d *SvcLog) GetAllServiceLog() (rows []dbStruct.ServiceLog, err error) {
	err = db.DB.Find(&rows).Error

	return
}

// GetOneByJobIDServiceLog - GetOneByJobIDServiceLog GetOne
func (d *SvcLog) GetOneByJobIDServiceLog(r dbStruct.ServiceLog) (row dbStruct.ServiceLog, err error) {
	err = db.DB.Where(&r).Find(&row).Error

	return
}

// InsertServiceLog - InsertServiceLog Insert
func (d *SvcLog) InsertServiceLog(tx *gorm.DB, v interface{}) (cnt int64, err error) {
	err = tx.Create(v).Error

	return cnt, nil
}

// UpdateByJobIDServiceLog - UpdateByJobIDServiceLog Update
func (d *SvcLog) UpdateByJobIDServiceLog(
	tx *gorm.DB,
	row dbStruct.ServiceLog,
) (err error) {

	err = tx.Model(&row).Where(constant.JobIDStr+" = ?", row.JobID).Update("req", row.Req).Error

	return
}

// UpdateReturnByJobIDServiceLog - UpdateReturnByJobIDServiceLog UpdateReturn
func (d *SvcLog) UpdateReturnByJobIDServiceLog(
	tx *gorm.DB,
	row dbStruct.ServiceLog,
) (rows []dbStruct.ServiceLog, err error) {
	q := []string{
		"UPDATE", tblServiceLog,
		"SET req = ?",
		"WHERE job_id = ?",
		"RETURNING type, job_id, req, res, errcode",
	}
	sql := strings.Join(q, " ")

	err = tx.Raw(sql, row.Req, row.JobID).Scan(&rows).Error

	return
}

// DeleteByJobIDServiceLog - DeleteByJobIDServiceLog Delete
func (d *SvcLog) DeleteByJobIDServiceLog(
	tx *gorm.DB,
	row dbStruct.ServiceLog,
) (err error) {
	err = tx.Where(constant.JobIDStr+" = ?", row.JobID).Delete(&row).Error

	return
}
