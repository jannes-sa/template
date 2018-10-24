package svclog

import (
	"template/helper/constant"
	iSvcLog "template/models/db/interfaces/svclog"
	pgSvcLog "template/models/db/pgsql/svclog"
	stSvcLog "template/models/stub/svclog"
)

const (
	logicName = "@svcLog"
)

var (
	// DBSvcLog ...
	DBSvcLog iSvcLog.ISvcLog
)

func init() {
	if constant.GOENV == constant.DEVCI {
		DBSvcLog = new(stSvcLog.SvcLog)
	} else {
		DBSvcLog = new(pgSvcLog.SvcLog)
	}
}
