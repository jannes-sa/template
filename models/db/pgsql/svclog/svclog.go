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

	q := "select * from ?"
	_, err = o.Raw(q, tblServiceLog).QueryRows(&rows)
	return
}

// GetOneServiceLog - GetOneServiceLog GetOne
func (d *SvcLog) GetOneServiceLog(cond string) (row dbStruct.ServiceLog, err error) {
	o := orm.NewOrm()

	q := "select * from ? where ?"
	err = o.Raw(q, tblServiceLog, cond).QueryRow(&row)

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

// UpdateServiceLog - UpdateServiceLog Update
func (d *SvcLog) UpdateServiceLog(
	o orm.Ormer,
	params []string,
	cond string,
) (err error) {

	q := []string{
		"UPDATE ?",
		"SET" + strings.Join(params, ","),
		"WHERE ?",
	}
	sql := strings.Join(q, " ")
	_, err = o.Raw(sql, tblServiceLog, cond).Exec()

	return
}

// UpdateReturnServiceLog - UpdateReturnServiceLog UpdateReturn
func (d *SvcLog) UpdateReturnServiceLog(
	o orm.Ormer,
	params []string,
	cond string,
	columnReturn []string,
) (rows []dbStruct.ServiceLog, err error) {
	q := []string{
		"UPDATE ?",
		"SET", strings.Join(params, ","),
		"WHERE ?",
		"RETURNING" + strings.Join(columnReturn, ","),
	}
	sql := strings.Join(q, " ")
	_, err = o.Raw(sql, tblServiceLog, cond).QueryRows(&rows)

	return
}

// DeleteServiceLog - DeleteServiceLog Delete
func (d *SvcLog) DeleteServiceLog(o orm.Ormer, cond string) (err error) {
	q := []string{
		"DELETE FROM ?",
		"WHERE ?",
	}
	sql := strings.Join(q, " ")
	_, err = o.Raw(sql, tblServiceLog, cond).Exec()

	return
}
