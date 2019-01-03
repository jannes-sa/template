package constant

import (
	"os"
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
