package http

import (
	ctrl "template/controllers/http"

	"github.com/astaxie/beego"
)

func init() {
	Router()
}

// Router - Routing
func Router() {
	beego.InsertFilter("/*", beego.BeforeRouter, BeforeFunc, true)
	beego.ErrorHandler("404", pageNotFound)

	ns := beego.NewNamespace("/template/v1",
		/*:STARTHTTP*/

		/* this router only for testing purpose */
		beego.NSNamespace("/test",
			beego.NSInclude(
				&ctrl.TestController{},
			),
		),

		/*:ENDHTTP*/
	)

	beego.AddNamespace(ns)
	beego.SetStaticPath("/storages", "storages")
	beego.InsertFilter("/*", beego.FinishRouter, AfterFunc, true)
}
