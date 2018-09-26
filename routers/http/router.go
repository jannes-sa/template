package http

import (
	"net/http"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func init() {
	DoRouter()
}

// DoRouter ...
func DoRouter() {
	ns := beego.NewNamespace("/template/v1",
		beego.NSBefore(Middleware),
		// Start: Add Your HTTP Router Here //

		// End : Add Your HTTP Router Here //
		beego.NSAfter(AfterFunc),
	)

	beego.AddNamespace(ns)

	beego.ErrorHandler("404", pageNotFound)
	// beego.InsertFilter("/*", beego.FinishRouter, AfterFunc, false)

	beego.SetStaticPath("/storages", "storages")
}

// pageNotFound ..
func pageNotFound(rw http.ResponseWriter, r *http.Request) {
	_, err := rw.Write([]byte(""))
	if err != nil {
		beego.Warning("NOT FOUND ERROR")
	}
}

// Middleware to call authentication
func Middleware(c *context.Context) {
	beego.Info("Middleware")
}

// AfterFunc to execute progress after response
func AfterFunc(c *context.Context) {
	// Code here for after execution event
}
