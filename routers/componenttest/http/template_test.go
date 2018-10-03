package http

import (
	"encoding/json"
	"strconv"
	"template/helper/constant"
	"template/routers/componenttest"
	"template/structs"
	httpStructs "template/structs/api/http"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/astaxie/beego"
)

var (
	host string = "http://127.0.0.1:" + strconv.Itoa(beego.BConfig.Listen.HTTPPort)
)

func TestTemplateSuccess(t *testing.T) {

	var req structs.ReqData
	var reqBody httpStructs.ReqTest
	reqBody.ID = 1
	req.ReqBody = reqBody
	by, _ := json.Marshal(req)

	res := componenttest.SendHTTP(
		"POST",
		host+"/"+constant.DOMAINNAME+"/v1/template",
		by,
	)

	Convey("TestTemplateSuccess", t, func() {
		Convey("Should Success", func() {
			So(len(res.Error), ShouldEqual, 0)
		})
	})
}
