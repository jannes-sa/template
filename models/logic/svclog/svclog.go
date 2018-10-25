package svclog

import (
	dbBase "template/models/db/pgsql2"
	"template/structs"
	dbStruct "template/structs/db"
	lStruct "template/structs/logic"
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
	contextStruct lStruct.ContextStruct,
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
	contextStruct lStruct.ContextStruct,
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

	tx := dbBase.DB.Begin()

	_, err := DBSvcLog.InsertServiceLog(tx, &row)
	if err != nil {
		structs.ErrorCode.DatabaseError.String(errCode, err.Error(), nmFunc, logicName)
		tx.Rollback()

		return
	}

	tx.Commit()
}

// UpdateByJobIDServiceLog - UpdateByJobIDServiceLog
func UpdateByJobIDServiceLog(
	contextStruct lStruct.ContextStruct,
	errCode *[]structs.TypeError,
) {
	var (
		nmFunc = "UpdateByJobIDServiceLog"
		row    dbStruct.ServiceLog
	)
	row.JobID = contextStruct.JobID
	row.Req = "REQ UPDATE DATA"

	tx := dbBase.DB.Begin()

	err := DBSvcLog.UpdateByJobIDServiceLog(tx, row)
	if err != nil {
		structs.ErrorCode.DatabaseError.String(errCode, err.Error(), nmFunc, logicName)
		tx.Rollback()

		return
	}

	tx.Commit()
}

// UpdateReturnByJobIDServiceLog - UpdateReturnByJobIDServiceLog
func UpdateReturnByJobIDServiceLog(
	contextStruct lStruct.ContextStruct,
	errCode *[]structs.TypeError,
) (rows []dbStruct.ServiceLog) {
	var (
		nmFunc = "UpdateReturnByJobIDServiceLog"
		row    dbStruct.ServiceLog
	)
	row.JobID = contextStruct.JobID
	row.Req = "REQ UPDATE DATA RETURN"

	tx := dbBase.DB.Begin()

	rows, err := DBSvcLog.UpdateReturnByJobIDServiceLog(tx, row)
	if err != nil {
		structs.ErrorCode.DatabaseError.String(errCode, err.Error(), nmFunc, logicName)
		tx.Rollback()

		return
	}

	tx.Commit()

	return
}

// DeleteByJobIDServiceLog - DeleteByJobIDServiceLog
func DeleteByJobIDServiceLog(
	contextStruct lStruct.ContextStruct,
	errCode *[]structs.TypeError,
) {
	var (
		nmFunc = "DeleteByJobIDServiceLog"
		row    dbStruct.ServiceLog
	)
	row.JobID = contextStruct.JobID

	tx := dbBase.DB.Begin()

	err := DBSvcLog.DeleteByJobIDServiceLog(tx, row)
	if err != nil {
		structs.ErrorCode.DatabaseError.String(errCode, err.Error(), nmFunc, logicName)
		tx.Rollback()

		return
	}

	tx.Commit()
}
