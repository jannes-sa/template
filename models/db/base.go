package db

import (
	"strconv"
	"template/helper"
	"template/helper/constant"

	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
)

// Session Set Session for DB
func Session() orm.Ormer {
	o := orm.NewOrm()

	return o
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

	errRegisterDataBase := orm.RegisterDataBase("default", "postgres",
		constant.CREDPGSQL, maxIdle, maxConn)
	helper.CheckErr("errRegisterDataBase@Register", errRegisterDataBase)

	RegisterModelPgSQL()
}

// RegisterModelPgSQL Register Model for pgSQL
func RegisterModelPgSQL() {
	// orm.ResetModelCache()

	if constant.GOENV == constant.LOCAL || constant.DEBUG == "1" {
		orm.Debug = true
	}

}
