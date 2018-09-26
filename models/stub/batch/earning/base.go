package earning

import (
	"errors"

	"github.com/astaxie/beego"
)

var (
	// TxAccount  ...
	TxAccount string
	// TxLog ...
	TxLog string
	// TxHistory ...
	TxHistory string
	// BatchEarning ...
	BatchEarning string
	// EarningBatchHistory ...
	EarningBatchHistory string
)

// CheckErrReturn ...
func CheckErrReturn(msg string, err *error) {
	if *err != nil {
		beego.Warning(msg)
		beego.Warning(err)
		*err = errors.New(msg)
	}
}
