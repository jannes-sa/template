package http

import (
	"template/routers/componenttest"
)

func init() {
	componenttest.HTTPInit()
	componenttest.DBinit()
}
