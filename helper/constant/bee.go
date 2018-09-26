package constant

import (
	"github.com/astaxie/beego"
)

var (
	APPPORT = beego.BConfig.Listen.HTTPPort
)
