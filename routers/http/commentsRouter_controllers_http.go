package http

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["template/controllers/http:TestController"] = append(beego.GlobalControllerRouter["template/controllers/http:TestController"],
		beego.ControllerComments{
			Method: "TestPost",
			Router: `/success`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["template/controllers/http:TestController"] = append(beego.GlobalControllerRouter["template/controllers/http:TestController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/success/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
