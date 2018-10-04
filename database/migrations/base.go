package main

import (
	"io/ioutil"
	"log"
	"template/helper/constant"
)

// GetQuery ...
func GetQuery(table string, dt string, dbfile string) string {
	fl := constant.GOPATH + "/src/" + constant.GOAPP +
		"/database/db/" + table + "/" + dt + "/" + dbfile + ".sql"
	raw, err := ioutil.ReadFile(fl)
	if err != nil {
		log.Println("failed Open File SQL")
		panic(err)
	}
	query := string(raw)
	return query
}
