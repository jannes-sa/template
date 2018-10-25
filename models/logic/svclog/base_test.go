package svclog

import (
	"template/helper/constant"
	db "template/models/db/pgsql"
	pg "template/models/db/pgsql2"

	_ "github.com/lib/pq"
)

func init() {
	initialize()
}

func initialize() {
	constant.LoadEnv()
	db.RegisterPGSQL()
	pg.ConnectDB()
}
