package http

import (
	"template/controllers"
	"template/structs"

	"github.com/astaxie/beego/context"
)

type fnURLFilter func(
	*context.Context,
	*[]structs.TypeError,
)

var filterMap map[string]fnURLFilter

func init() {
	filterMap = map[string]fnURLFilter{}

}

// Filtering URL before go into controller
func filterURL(c *context.Context) {
	var errCode []structs.TypeError

	if valFunc, ok := filterMap[c.Input.URL()]; ok {
		valFunc(c, &errCode)
		if len(errCode) > 0 {
			var t interface{}
			controllers.SendOutput(c, t, errCode)
		}
	}
}
