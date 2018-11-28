package svclog

import (
	dbBase "template/models/db/pgsql2"
	"template/models/logic/logic2/dellogic2"
	"template/models/logic/svclog/delsvclog"
	"template/structs"
	dbStruct "template/structs/db"
	lStruct "template/structs/logic"

	"github.com/jinzhu/gorm/dialects/postgres"
)

func init() {
	delsvclog.Register("logicSvcLog", &logicSvcLog{})
}

type logicSvcLog struct{}

func (l logicSvcLog) Call() int {
	if _, ok := dellogic2.Function[dellogic2.LogicName]; !ok {
		panic("function not exists")
	}
	logicLogic2 := dellogic2.Function[dellogic2.LogicName]

	calc := logicLogic2.Receive(2)
	return calc
}

func (l logicSvcLog) Receive(input int) int {
	return input * 2
}

// delegate public function testing //

// GetAllServiceLog - GetAllServiceLog
func (l logicSvcLog) getAllServiceLog(errCode *[]structs.TypeError) (rows []dbStruct.ServiceLog) {
	rows, err := DBSvcLog.GetAllServiceLog()
	if err != nil {
		structs.ErrorCode.DatabaseError.String(errCode, err.Error(), "GetAllServiceLog ", logicName)
	}

	return
}

// GetOneByJobIDServiceLog - GetOneByJobIDServiceLog
func (l logicSvcLog) getOneByJobIDServiceLog(
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
func (l logicSvcLog) insertServiceLog(
	contextStruct lStruct.ContextStruct,
	errCode *[]structs.TypeError,
) {
	var (
		nmFunc = "InsertServiceLog"
		row    dbStruct.ServiceLog
	)

	row.JobID = contextStruct.JobID
	row.Req = postgres.Jsonb{[]byte(`{"update":"1"}`)}
	row.Res = postgres.Jsonb{[]byte("{}")}
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
func (l logicSvcLog) updateByJobIDServiceLog(
	contextStruct lStruct.ContextStruct,
	errCode *[]structs.TypeError,
) {
	var (
		nmFunc = "UpdateByJobIDServiceLog"
		row    dbStruct.ServiceLog
	)
	row.JobID = contextStruct.JobID
	row.Req = postgres.Jsonb{[]byte("{}")}

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
func (l logicSvcLog) updateReturnByJobIDServiceLog(
	contextStruct lStruct.ContextStruct,
	errCode *[]structs.TypeError,
) (rows []dbStruct.ServiceLog) {
	var (
		nmFunc = "UpdateReturnByJobIDServiceLog"
		row    dbStruct.ServiceLog
	)
	row.JobID = contextStruct.JobID
	row.Req = postgres.Jsonb{[]byte(`{"update":"1"}`)}

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
func (l logicSvcLog) deleteByJobIDServiceLog(
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
