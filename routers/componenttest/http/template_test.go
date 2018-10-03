package http

import (
	"encoding/json"
	"template/helper/constant"
	"template/routers/componenttest"
	"template/structs"
	httpStructs "template/structs/api/http"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
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
