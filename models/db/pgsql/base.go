package pgsql

import (
	"strconv"
	"template/helper"
	"template/helper/constant"

	"github.com/astaxie/beego"

	"github.com/astaxie/beego/orm"
)

// Session Set Session for DB
func Session() orm.Ormer {
	db := orm.NewOrm()
	return db
}

// RegisterPGSQL - Registration pgSQL DB
func RegisterPGSQL() {
	var maxIdle = 50
	var maxConn = 50
	if constant.DBMAXIDLE != "" {
		maxIdleTemp, err := strconv.Atoi(constant.DBMAXIDLE)
		helper.CheckErr("", err)
		maxIdle = maxIdleTemp
	}
	if constant.DBMAXCONN != "" {
		maxConnTemp, err := strconv.Atoi(constant.DBMAXCONN)
		helper.CheckErr("", err)
		maxConn = maxConnTemp
	}

	errRegisterDriver := orm.RegisterDriver("postgres", orm.DRPostgres)
	helper.CheckErr("errRegisterDriver@Register", errRegisterDriver)
	if errRegisterDriver != nil {
		beego.Critical(errRegisterDriver)
		panic(1)
	}

	errRegisterDataBase := orm.RegisterDataBase("default", "postgres",
		constant.CREDPGSQL, maxIdle, maxConn)
	helper.CheckErr("errRegisterDataBase@Register", errRegisterDataBase)
	if errRegisterDataBase != nil {
		beego.Critical(errRegisterDataBase)
		panic(1)
	}

	RegisterModelPgSQL()
}

// RegisterModelPgSQL Register Model for pgSQL
func RegisterModelPgSQL() {
	// orm.ResetModelCache()

	// if constant.GOENV == constant.LOCAL || constant.DEBUG == "1" {
	orm.Debug = true
	// }

}
