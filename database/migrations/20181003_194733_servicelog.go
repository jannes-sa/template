package main

import (
	"template/helper/constant/tablename"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Servicelog_20181003_194733 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Servicelog_20181003_194733{}
	m.Created = "20181003_194733"

	migration.Register("Servicelog_20181003_194733", m)
}

// Run the migrations
func (m *Servicelog_20181003_194733) Up() {
	dt := "20181003"
	m.SQL(GetQuery(tablename.ServiceLog, dt, "service_log"))
}

// Reverse the migrations
func (m *Servicelog_20181003_194733) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
