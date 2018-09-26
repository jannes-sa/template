package helper

import (
	"errors"
	"template/helper/constant"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

var (
	exceptionLastInsertID = "no LastInsertId available"
)

// CheckErr - CheckErr Message
func CheckErr(val string, err error) {
	if err != nil {
		beego.Warning(val, err)
	}
}

// CheckErrReturn ...
func CheckErrReturn(msg string, err *error) {
	if *err != nil {
		beego.Warning(msg, err)
		*err = errors.New(msg)
	}
}

// CheckErrWithRollBack ...
func CheckErrWithRollBack(
	msg string,
	forceRB bool,
	o orm.Ormer,
	stateTx *bool,
	err error,
) {
	if forceRB {
		*stateTx = false
		beego.Warning(msg, "Force Rollback")
		errRollback := o.Rollback()
		CheckErr("errRollback@CheckErrRB", errRollback)
	} else {
		if err != nil {
			*stateTx = false
			beego.Warning(msg, err)
			errRollback := o.Rollback()
			CheckErr("errRollback@CheckErrRB", errRollback)
		}
	}
}

// CheckErrRB Rollback if Error
func CheckErrRB(msg string, forceRB bool, o orm.Ormer, stateTx *bool,
	err error) {
	beego.Warning(msg, "=>", err)
	if forceRB {
		*stateTx = false
		beego.Warning(msg, "Force Rollback")
		errRollback := o.Rollback()
		CheckErr("errRollback@CheckErrRB", errRollback)
	} else if err != nil && err.Error() != exceptionLastInsertID {
		*stateTx = false
		errRollback := o.Rollback()
		CheckErr("errRollback@CheckErrRB", errRollback)
	}
}

// CheckErrRollback Rollback if Error
func CheckErrRollback(msg string, forceRB bool, o orm.Ormer, err error) {
	beego.Warning(msg, "=>", err)
	if forceRB {
		beego.Warning(msg, "Force Rollback")
		errRollback := o.Rollback()
		CheckErr("errRollback@CheckErrRB", errRollback)
	} else if err != nil && err.Error() != constant.ExceptionLastInsertID {
		errRollback := o.Rollback()
		CheckErr("errRollback@CheckErrRB", errRollback)
	}
}

// CheckErrorLatestInsert ...
func CheckErrorLatestInsert(
	err error,
) (statusErr bool) {
	statusErr = false
	if err != nil && err.Error() != exceptionLastInsertID {
		statusErr = true
	}
	return
}
