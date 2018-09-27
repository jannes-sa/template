package http

import (
	"encoding/json"
	"template/helper"
	"template/structs"
	structsAPI "template/structs/api/http"

	"github.com/astaxie/beego"
)

// TestController -
type TestController struct {
	beego.Controller
}

// URLMapping ...
func (c *TestController) URLMapping() {
	c.Mapping("get", c.Get)
	// c.Mapping("TestPost", c.TestPost)
}

// Get ...
// @Title Get
// @router /success/:id [get]
func (c *TestController) Get() {
	errCode := make([]structs.TypeError, 0)
	id := c.Ctx.Input.Param(":id")

	c.Data["json"] = `{"data":"id":"` + id + `"}`

	SendOutput(c.Ctx, c.Data["json"], errCode)
}

// TestPost ...
// @Title TestPost TestPost
// @router /success [post]
func (c *TestController) TestPost() {
	errCode := make([]structs.TypeError, 0)

	var (
		reqInterface structsAPI.TestInterface
		req          structsAPI.ReqTest
		res          structsAPI.ResTest
	)

	rqBodyByte := helper.GetRqBodyRev(c.Ctx, &errCode)
	if len(errCode) > 0 {
		SendOutput(c.Ctx, c.Data["json"], errCode)
		return
	}

	err := json.Unmarshal(rqBodyByte, &reqInterface)
	if err != nil {
		structs.ErrorCode.UnexpectedError.String(&errCode)
		SendOutput(c.Ctx, c.Data["json"], errCode)
		return
	}

	reqInterface.ValidateRequest(&req, &errCode)
	if len(errCode) > 0 {
		SendOutput(c.Ctx, c.Data["json"], errCode)
		return
	}

	res.ID = req.ID

	c.Data["json"] = res

	SendOutput(c.Ctx, c.Data["json"], errCode)
}
