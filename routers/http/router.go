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
		// Start: Add Your HTTP Router Here //

		/* this router only for testing purpose */
		beego.NSNamespace("/test",
			beego.NSInclude(
				&ctrl.TestController{},
			),
		),

		// End : Add Your HTTP Router Here //

	)

	beego.AddNamespace(ns)
	beego.SetStaticPath("/storages", "storages")
	beego.InsertFilter("/*", beego.FinishRouter, AfterFunc, true)
}
