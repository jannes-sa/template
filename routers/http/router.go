package http

import (
	h "net/http"
	ctrl "template/controllers/http"

	"github.com/astaxie/beego"
)

// Router - Routing
func Router() {
	beego.InsertFilter("/*", beego.BeforeRouter, BeforeFunc, true)
	beego.ErrorHandler("404", pageNotFound)

	ns := beego.NewNamespace("/template/v1",
		// beego.NSBefore(Middleware),

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

	beego.InsertFilter("/*", beego.FinishRouter, AfterFunc, false)
}

// pageNotFound ..
func pageNotFound(rw h.ResponseWriter, r *h.Request) {
	_, err := rw.Write([]byte(""))
	if err != nil {
		beego.Warning("NOT FOUND ERROR")
	}
}
