package svclog

import (
	"template/helper/constant"
	iSvcLog "template/models/db/interfaces/svclog"
	pgSvcLog "template/models/db/pgsql/svclog"
)

var DBSvcLog iSvcLog.ISvcLog

func init() {
	if constant.GOENV == constant.DEVCI {

	} else {
		DBSvcLog = new(pgSvcLog.SvcLog)
	}
}
