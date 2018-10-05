package svclog

import (
	"template/helper/constant"
	db "template/models/db/pgsql"

	_ "github.com/lib/pq"
)

func init() {
	initialize()
}

func initialize() {
	constant.LoadEnv()
	db.RegisterPGSQL()
}
