package svclog

import (
	"template/helper"
	dbBase "template/models/db/pgsql"
	"template/structs"
	dbStruct "template/structs/db"
	logicStruct "template/structs/logic"
)

// GetAllServiceLog - GetAllServiceLog
func GetAllServiceLog(errCode *[]structs.TypeError) (rows []dbStruct.ServiceLog) {
	rows, err := DBSvcLog.GetAllServiceLog()
	if err != nil {
		structs.ErrorCode.DatabaseError.String(errCode, err.Error(), "GetAllServiceLog ", logicName)
	}

	return
}

// GetOneByJobIDServiceLog - GetOneByJobIDServiceLog
func GetOneByJobIDServiceLog(
	contextStruct logicStruct.ContextStruct,
	errCode *[]structs.TypeError,
) (row dbStruct.ServiceLog) {
	row.JobID = contextStruct.JobID
	row, err := DBSvcLog.GetOneByJobIDServiceLog(row)
	if err != nil {
		structs.ErrorCode.DatabaseError.String(errCode, err.Error(), "GetOneByJobIDServiceLog", logicName)
	}

	return

}

// InsertServiceLog - InsertServiceLog
func InsertServiceLog(
	contextStruct logicStruct.ContextStruct,
	errCode *[]structs.TypeError,
) {
	var (
		nmFunc = "InsertServiceLog"
		row    dbStruct.ServiceLog
	)

	row.JobID = contextStruct.JobID
	row.Req = "req"
	row.Res = "res"
	row.Errcode = "ERRCODE"
	row.Type = "http"

	db := dbBase.Session()
	err := db.Begin()
	if err != nil {
		structs.ErrorCode.DatabaseError.String(errCode, err.Error(), nmFunc, logicName)
		return
	}

	_, err = DBSvcLog.InsertServiceLog(db, &row)
	if err != nil {
		structs.ErrorCode.DatabaseError.String(errCode, err.Error(), nmFunc, logicName)
		err = db.Rollback()
		helper.CheckErr(nmFunc+" "+logicName, err)

		return
	}

	err = db.Commit()
	if err != nil {
		structs.ErrorCode.DatabaseError.String(errCode, err.Error(), nmFunc, logicName)
		err = db.Rollback()
		helper.CheckErr(nmFunc+" "+logicName, err)

		return
	}
}

// UpdateByJobIDServiceLog - UpdateByJobIDServiceLog
func UpdateByJobIDServiceLog(
	contextStruct logicStruct.ContextStruct,
	errCode *[]structs.TypeError,
) {
	var (
		nmFunc = "UpdateByJobIDServiceLog"
		row    dbStruct.ServiceLog
	)
	row.JobID = contextStruct.JobID
	row.Req = "REQ UPDATE DATA"

	db := dbBase.Session()
	err := db.Begin()
	if err != nil {
		structs.ErrorCode.DatabaseError.String(errCode, err.Error(), nmFunc, logicName)
		return
	}

	err = DBSvcLog.UpdateByJobIDServiceLog(db, row)
	if err != nil {
		structs.ErrorCode.DatabaseError.String(errCode, err.Error(), nmFunc, logicName)
		err = db.Rollback()
		helper.CheckErr(nmFunc+" "+logicName, err)

		return
	}

	err = db.Commit()
	if err != nil {
		structs.ErrorCode.DatabaseError.String(errCode, err.Error(), nmFunc, logicName)
		err = db.Rollback()
		helper.CheckErr(nmFunc+" "+logicName, err)

		return
	}
}

// UpdateReturnByJobIDServiceLog - UpdateReturnByJobIDServiceLog
func UpdateReturnByJobIDServiceLog(
	contextStruct logicStruct.ContextStruct,
	errCode *[]structs.TypeError,
) (rows []dbStruct.ServiceLog) {
	var (
		nmFunc = "UpdateReturnByJobIDServiceLog"
		row    dbStruct.ServiceLog
	)
	row.JobID = contextStruct.JobID
	row.Req = "REQ UPDATE DATA RETURN"

	db := dbBase.Session()
	err := db.Begin()
	if err != nil {
		structs.ErrorCode.DatabaseError.String(errCode, err.Error(), nmFunc, logicName)
		return
	}

	rows, err = DBSvcLog.UpdateReturnByJobIDServiceLog(db, row)
	if err != nil {
		structs.ErrorCode.DatabaseError.String(errCode, err.Error(), nmFunc, logicName)
		err = db.Rollback()
		helper.CheckErr(nmFunc+" "+logicName, err)

		return
	}

	err = db.Commit()
	if err != nil {
		structs.ErrorCode.DatabaseError.String(errCode, err.Error(), nmFunc, logicName)
		err = db.Rollback()
		helper.CheckErr(nmFunc+" "+logicName, err)

		return
	}

	return
}

// DeleteByJobIDServiceLog - DeleteByJobIDServiceLog
func DeleteByJobIDServiceLog(
	contextStruct logicStruct.ContextStruct,
	errCode *[]structs.TypeError,
) {
	var (
		nmFunc = "DeleteByJobIDServiceLog"
		row    dbStruct.ServiceLog
	)
	row.JobID = contextStruct.JobID

	db := dbBase.Session()
	err := db.Begin()
	if err != nil {
		structs.ErrorCode.DatabaseError.String(errCode, err.Error(), nmFunc, logicName)
		return
	}

	err = DBSvcLog.DeleteByJobIDServiceLog(db, row)
	if err != nil {
		structs.ErrorCode.DatabaseError.String(errCode, err.Error(), nmFunc, logicName)
		err = db.Rollback()
		helper.CheckErr(nmFunc+" "+logicName, err)

		return
	}

	err = db.Commit()
	if err != nil {
		structs.ErrorCode.DatabaseError.String(errCode, err.Error(), nmFunc, logicName)
		err = db.Rollback()
		helper.CheckErr(nmFunc+" "+logicName, err)

		return
	}
}
