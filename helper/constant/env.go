package constant

import (
	"os"
	"strings"

	"github.com/astaxie/beego"
	"github.com/joho/godotenv"
)

var (
	GOPATH    string
	GOAPP     string
	GOENV     string
	CREDPGSQL string
	CREDMONGO string

	DBMAXCONN string
	DBMAXIDLE string

	AUTH    string
	AUTHKEY string
	AUTHEXP string
	TZ      string

	VERSION string
)

func init() {
	LoadEnv()

	GOPATH = os.Getenv("GOPATH")
	GOAPP = os.Getenv("GOAPP")
	GOENV = os.Getenv("GOENV")

	CREDPGSQL = os.Getenv("CRED_PGSQL")
	CREDMONGO = os.Getenv("CRED_MONGO")

	DBMAXCONN = os.Getenv("DBMAXCONN")
	DBMAXIDLE = os.Getenv("DBMAXIDLE")

	AUTH = os.Getenv("AUTH")
	AUTHKEY = os.Getenv("AUTHKEY")
	AUTHEXP = os.Getenv("AUTHEXP")
	TZ = os.Getenv("TZ")

	VERSION = os.Getenv("VERSION")
}

// LoadEnv - LoadEnv
func LoadEnv() {
	if os.Getenv("GOENV") == DEVCI || strings.ToLower(os.Getenv("GOENV")) == LOCAL {
		errEnv := godotenv.Load(
			os.Getenv("GOPATH") + "/src/" + os.Getenv("GOAPP") +
				"/conf/env")
		if errEnv != nil {
			beego.Critical("fatal load env", errEnv)
		}
	}
}
