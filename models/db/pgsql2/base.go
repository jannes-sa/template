package pgsql2

import (
	"template/helper"
	"template/helper/constant"
	"time"

	"github.com/astaxie/beego"
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// DB - DB
var DB *gorm.DB

// ConnectDB - ConnectDB
func ConnectDB() {
	db, err := gorm.Open("postgres", constant.CREDPGSQL)
	if err != nil {
		helper.CheckErr("Failed Connect", err)
		panic(err)
	}

	db.LogMode(true)
	db.DB().SetMaxIdleConns(0)
	db.DB().SetMaxOpenConns(0)
	db.DB().SetConnMaxLifetime(10 * time.Minute)

	beego.Info("DB Connected", db.DB().Ping())

	DB = db

	return
}
