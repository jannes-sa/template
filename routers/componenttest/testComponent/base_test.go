package testComponent

import (
	"template/routers/componenttest"
)

func init() {
	componenttest.HTTPInit()
	componenttest.GRPCInit()
	componenttest.DBinit()
}
