package main

import (
	"os"
	"strconv"
	"template/helper/constant"
	"template/helper/timetn"
	pg "template/models/db/pgsql2"
	"template/routers/grpc"
	_ "template/routers/http"
	_ "template/structs/db"
	"time"

	_ "github.com/lib/pq"

	"net/http"
	_ "net/http/pprof"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

func logFile() string {
	hostname, _ := os.Hostname()
	now := timetn.Now()

	year := now.Year()
	month := int(now.Month())
	day := now.Day()
	hours := now.Hour()

	strTime := strconv.Itoa(year) + `-` +
		strconv.Itoa(month) + `-` +
		strconv.Itoa(day) + `-` +
		strconv.Itoa(hours)

	pathLog := "logs/" + hostname + "_" + strTime +
		"_" + constant.GOAPP + ".log"
	return pathLog
}

func main() {
	setup()
	if constant.GOENV != "" && constant.GOAPP == beego.BConfig.AppName {
		time.AfterFunc(1*time.Second, initialData)

		beego.BConfig.Listen.Graceful = true
		if beego.BConfig.RunMode == "dev" {
			beego.BConfig.WebConfig.DirectoryIndex = true
			beego.BConfig.Listen.EnableAdmin = true
			beego.BConfig.Listen.AdminAddr = "localhost"

			adminPort, _ := strconv.Atoi("2" + strconv.Itoa(beego.BConfig.Listen.HTTPPort))
			beego.BConfig.Listen.AdminPort = adminPort
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

	logFile := logFile()
	err := logs.SetLogger(logs.AdapterFile, `{"filename":"`+logFile+`",
		"level":7,"maxlines":20000,"maxsize":0,"daily":true,"maxdays":10}`)
	if err != nil {
		panic(err)
	}

	// db.RegisterPGSQL()
	pg.ConnectDB()
}

func initialData() {
	// path := constant.GOPATH + "/src/" + constant.GOAPP + "/storages"
	// if _, err := os.Stat(path); os.IsNotExist(err) {
	// 	os.MkdirAll(path, os.ModePerm)
	// } else {
	// 	beego.Debug(err)
	// }

	grpc.CreateGrpcServer("")
	go initPprof()
}

func initPprof() {
	portHTTP := "1" + strconv.Itoa(beego.BConfig.Listen.HTTPPort)
	http.ListenAndServe(":"+portHTTP, nil)
}
