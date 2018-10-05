package svclog

import (
	"template/helper/constant"
	iSvcLog "template/models/db/interfaces/svclog"
	pgSvcLog "template/models/db/pgsql/svclog"
	stSvcLog "template/models/stub/svclog"
)

var (
	logicName = "@svcLog"
	DBSvcLog  iSvcLog.ISvcLog
)

func init() {
	if constant.GOENV == constant.DEVCI {
		DBSvcLog = new(stSvcLog.SvcLog)
	} else {
		DBSvcLog = new(pgSvcLog.SvcLog)
	}
}
