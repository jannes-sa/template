package http

import (
	"strconv"
	"template/routers/componenttest"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/astaxie/beego"
)

func TestTemplateSuccess(t *testing.T) {
	host := "http://127.0.0.1:" + strconv.Itoa(beego.BConfig.Listen.HTTPPort)

	res := componenttest.SendHTTP(
		"POST",
		host+"/template/v1/template",
		[]byte(`{
			"rqBody":{
				"id":1
			}
		}`),
	)

	Convey("TestTemplateSuccess", t, func() {
		Convey("Should Success", func() {
			So(len(res.Error), ShouldEqual, 0)
		})
	})

}

func TestTemplateFailed(t *testing.T) {
	host := "http://127.0.0.1:" + strconv.Itoa(beego.BConfig.Listen.HTTPPort)

	res := componenttest.SendHTTP(
		"POST",
		host+"/template/v1/template",
		[]byte(`{
			"rqBody":{
			}
		}`),
	)

	Convey("TestTemplateFailed", t, func() {
		Convey("Should Failed", func() {
			So(len(res.Error), ShouldEqual, 1)
		})
	})
}
