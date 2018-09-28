package constant

import (
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

var (
	GOPATH        string
	GOAPP         string
	GOENV         string
	SMITHBANKCODE string

	MQ   string
	HTTP string

	DBMAXIDLE                  string
	DBMAXCONN                  string
	PATHBATCHEARNING           string
	PATHBATCHLOYALTYRECON      string
	PATHBATCHDEACTIVATE        string
	PATHBATCHCLEARINGPOINT     string
	PATHBATCHREVERSEREDEMPTION string
	NUMROUTINE                 string

	BANKCODE    string
	BANKNAME    string
	DOMAIN      string
	DOMAINID    string
	DOMAINIDFIX string

	CREDMONGODB string
	CREDMQ      string
	CREDREDIS   string
	CREDPGSQL   string
	CREDSQLITE  string

	BATCH string

	FIXTXNDOMAINHOST string
	GLDOMAINHOST     string

	RPCRULESTXN string
	HTTPTXN     string
	HTTPSELFSVC string
	HTTPOTHER   string
	MQADDRESSGL string

	STATICPATH string

	DEBUG string
)

func init() {
	if os.Getenv("GOENV") == "" || strings.ToLower(os.Getenv("GOENV")) == LOCAL {
		errEnv := godotenv.Load(
			os.Getenv("GOPATH") + "/src/" + os.Getenv("GOAPP") +
				"/environment/env")
		if errEnv != nil {
			log.Fatal("fatal load env", errEnv)
		}
	}

	GOPATH = os.Getenv("GOPATH")
	GOAPP = os.Getenv("GOAPP")
	GOENV = os.Getenv("GOENV")
	SMITHBANKCODE = os.Getenv("SMITHBANKCODE")

	MQ = os.Getenv("MQ")
	HTTP = os.Getenv("HTTP")

	DBMAXIDLE = os.Getenv("DB_MAXIDLE")
	DBMAXCONN = os.Getenv("DB_MAXCONN")
	PATHBATCHEARNING = os.Getenv("PATH_BATCH_EARNING")
	PATHBATCHLOYALTYRECON = os.Getenv("PATH_BATCH_LOYALTY_RECON")
	PATHBATCHDEACTIVATE = os.Getenv("PATH_BATCH_DEACTIVATE")
	PATHBATCHCLEARINGPOINT = os.Getenv("PATH_BATCH_CLEARING_POINT")
	PATHBATCHREVERSEREDEMPTION = os.Getenv("PATH_BATCH_REVERSE_REDEMPTION")
	NUMROUTINE = os.Getenv("NUM_ROUTINE")

	BANKCODE = os.Getenv("BANK_CODE")
	BANKNAME = os.Getenv("BANK_NAME")
	DOMAIN = os.Getenv("DOMAIN")
	DOMAINID = os.Getenv("DOMAIN_ID")
	DOMAINIDFIX = os.Getenv("DOMAIN_ID_FIX")

	CREDMONGODB = os.Getenv("CRED_MONGODB")
	CREDMQ = os.Getenv("CRED_MQ")
	CREDREDIS = os.Getenv("CRED_REDIS")
	CREDPGSQL = os.Getenv("CRED_PGSQL")
	CREDSQLITE = os.Getenv("CRED_SQLITE")

	BATCH = os.Getenv("batch")

	FIXTXNDOMAINHOST = os.Getenv("FIX_TXN_DOMAIN_HOST")
	GLDOMAINHOST = os.Getenv("GL_DOMAIN_HOST")

	RPCRULESTXN = os.Getenv("RPC_RULESTXN")
	HTTPTXN = os.Getenv("HTTP_TXN")
	HTTPSELFSVC = os.Getenv("HTTP_SELFSVC")
	HTTPOTHER = os.Getenv("HTTP_OTHER")
	MQADDRESSGL = os.Getenv("MQ_ADDRESS_GL")

	STATICPATH = os.Getenv("STATIC_PATH")

	DEBUG = os.Getenv("DEBUG")
}
