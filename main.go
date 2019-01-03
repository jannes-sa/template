package main

import (
	"net/http"
	_ "net/http/pprof"
	"strconv"
	"template/helper/constant"
	"template/routers/grpc"
	_ "template/routers/http"

	appLog "template/logs"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

func main() {
	setup()
	if constant.GOENV != "" && constant.GOAPP == beego.BConfig.AppName {
		beego.BConfig.Listen.Graceful = true
		if beego.BConfig.RunMode == "dev" {
			beego.BConfig.WebConfig.DirectoryIndex = true
			beego.BConfig.Listen.EnableAdmin = true
			beego.BConfig.Listen.AdminAddr = "localhost"

			adminPort, _ := strconv.Atoi("2" + strconv.Itoa(beego.BConfig.Listen.HTTPPort))
			beego.BConfig.Listen.AdminPort = adminPort
			beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
		} else {
			// beego.SetLevel(beego.LevelInformational)
			beego.BConfig.RecoverPanic = true
			beego.BConfig.Listen.ServerTimeOut = 680
		}

		beego.Run()

	} else {
		beego.Error("SETUP GOENV && GOAPP FIRST")
	}
}

func setup() {
	logs.Async()
	appLog.InitLog()
	go grpc.CreateGrpcServer("")
	go initPprof()
	// gc.SetProfiler()

	// db.RegisterPGSQL()
}

func initPprof() {
	portHTTP := "1" + strconv.Itoa(beego.BConfig.Listen.HTTPPort)
	http.ListenAndServe("localhost:"+portHTTP, nil)
}
