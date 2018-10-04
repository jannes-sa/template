package http

import (
	"encoding/json"
	"template/helper"
	"template/structs"
	structsAPI "template/structs/api/http"

	"github.com/astaxie/beego"
)

// TemplateController operations for Template
type TemplateController struct {
	beego.Controller
}

// URLMapping ...
func (c *TemplateController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Create
// @Description create Template
// @Param	body		body 	models.Template	true		"body for Template content"
// @Success 201 {object} models.Template
// @Failure 403 body is empty
// @router / [post]
func (c *TemplateController) Post() {
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

	SendOutput(c.Ctx, res, errCode)

}

// GetOne ...
// @Title GetOne
// @Description get Template by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Template
// @Failure 403 :id is empty
// @router /:id [get]
func (c *TemplateController) GetOne() {
	errCode := make([]structs.TypeError, 0)

	id := c.Ctx.Input.Param(":id")
	beego.Debug(id)

	SendOutput(c.Ctx, c.Data["json"], errCode)

}

// GetAll ...
// @Title GetAll
// @Description get Template
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Template
// @Failure 403
// @router / [get]
func (c *TemplateController) GetAll() {

}

// Put ...
// @Title Put
// @Description update the Template
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Template	true		"body for Template content"
// @Success 200 {object} models.Template
// @Failure 403 :id is not int
// @router /:id [put]
func (c *TemplateController) Put() {

}

// Delete ...
// @Title Delete
// @Description delete the Template
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *TemplateController) Delete() {

}
