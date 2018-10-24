package svclog

import (
	"strings"
	"template/helper"
	"template/helper/constant"
	"template/helper/constant/tablename"
	dbStruct "template/structs/db"

	"github.com/astaxie/beego/orm"
)

// SvcLog - Logic Struct DB
type SvcLog struct{}

var tblServiceLog = tablename.ServiceLog

// GetAllServiceLog - GetAllServiceLog GetAll
func (d *SvcLog) GetAllServiceLog() (rows []dbStruct.ServiceLog, err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable(tblServiceLog).All(&rows)
	return
}

// GetOneByJobIDServiceLog - GetOneByJobIDServiceLog GetOne
func (d *SvcLog) GetOneByJobIDServiceLog(r dbStruct.ServiceLog) (row dbStruct.ServiceLog, err error) {
	o := orm.NewOrm()
	err = o.QueryTable(tblServiceLog).Filter(constant.JobIDStr, r.JobID).One(&row)
	return
}

// InsertServiceLog - InsertServiceLog Insert
func (d *SvcLog) InsertServiceLog(o orm.Ormer, v interface{}) (cnt int64, err error) {
	cnt, err = o.Insert(v)

	if err.Error() != constant.ExceptionLastInsertID {
		helper.CheckErr("Failed Inserted", err)
		return
	}

	return cnt, nil
}

// UpdateByJobIDServiceLog - UpdateByJobIDServiceLog Update
func (d *SvcLog) UpdateByJobIDServiceLog(
	o orm.Ormer,
	row dbStruct.ServiceLog,
) (err error) {

	_, err = o.QueryTable(tblServiceLog).Filter(constant.JobIDStr, row.JobID).Update(orm.Params{
		"req": row.Req,
	})

	return
}

// UpdateReturnByJobIDServiceLog - UpdateReturnByJobIDServiceLog UpdateReturn
func (d *SvcLog) UpdateReturnByJobIDServiceLog(
	o orm.Ormer,
	row dbStruct.ServiceLog,
) (rows []dbStruct.ServiceLog, err error) {
	q := []string{
		"UPDATE", tblServiceLog,
		"SET req = ?",
		"WHERE job_id = ?",
		"RETURNING type, job_id, req, res, errcode",
	}
	sql := strings.Join(q, " ")
	_, err = o.Raw(sql, row.Req, row.JobID).QueryRows(&rows)

	return
}

// DeleteByJobIDServiceLog - DeleteByJobIDServiceLog Delete
func (d *SvcLog) DeleteByJobIDServiceLog(
	o orm.Ormer,
	row dbStruct.ServiceLog,
) (err error) {
	_, err = o.QueryTable(tblServiceLog).Filter(constant.JobIDStr, row.JobID).Delete()
	return
}
