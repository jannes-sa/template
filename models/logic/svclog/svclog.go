package svclog

import (
	"template/structs"
	dbStruct "template/structs/db"
)

// GetAll - GetAll svclog
func GetAll(errCode *[]structs.TypeError) (rows []dbStruct.ServiceLog, err error) {
	rows, err = DBSvcLog.GetAllServiceLog()
	if err != nil {
		structs.ErrorCode.DatabaseError.String(errCode, err.Error(), "GetAll svclog")
	}

	return
}
