package grpc

import (
	"strconv"
	"template/helper/constant"
	"template/routers/componenttest"

	"github.com/astaxie/beego"
)

var (
	prefix = "/" + constant.GOAPP + "/" + constant.VERSION
	host   = "127.0.0.1:5" + strconv.Itoa(beego.BConfig.Listen.HTTPPort)
)

func init() {
	// componenttest.HTTPInit()
	componenttest.GRPCInit()
	componenttest.DBinit()
}
