package logs

import (
	"os"
	"strconv"
	"template/helper/constant"
	"template/helper/timetn"

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

// InitLog - InitLog
func InitLog() {
	logFile := logFile()
	err := logs.SetLogger(logs.AdapterFile, `{"filename":"`+logFile+`",
		"level":7,"maxlines":20000,"maxsize":0,"daily":true,"maxdays":10}`)
	if err != nil {
		panic(err)
	}
}
