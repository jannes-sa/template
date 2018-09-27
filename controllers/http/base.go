package http

import (
	"encoding/json"
	"template/helper"
	"template/structs"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

// SendOutput ...
func SendOutput(
	c *context.Context,
	jsonBody interface{},
	arrErrType []structs.TypeError,
) {

	if len(arrErrType) > 0 {
		c.Output.SetStatus(400)
	}

	SendResponse(c, jsonBody, arrErrType)
}

// SendResponse Sending Response
func SendResponse(
	c *context.Context,
	jsonBody interface{},
	arrErrType []structs.TypeError,
) {

	// Set Header //
	structHeaderResponse := helper.ConstructHTTPHeader(c)
	// Set Header //

	var (
		hasIndent   = true
		hasEncoding = false
	)
	if beego.BConfig.RunMode == "prod" {
		hasIndent = false
	}

	var respData structs.RespData
	respData.ResponseBody = jsonBody

	filErrCode := make([]structs.FilterErrorCode, 0)
	if len(arrErrType) > 0 {
		var t interface{}
		respData.ResponseBody = t
		filErrCode = filterErrorCode(arrErrType)
	}
	respData.Error = filErrCode

	var t interface{} = respData

	go printResponseLog(c, structHeaderResponse, respData)
	err := c.Output.JSON(t, hasIndent, hasEncoding)
	if err != nil {
		panic("ERROR OUTPUT JSON LEVEL MIDDLEWARE")
	}
}

func printResponseLog(
	c *context.Context,
	structHeaderResponse structs.ResHTTPHeader,
	respData structs.RespData,
) {

	jsonByte, err := json.Marshal(respData)
	helper.CheckErr("Failed Marshal line 156 controller/base", err)

	headerByte, err := json.Marshal(structHeaderResponse)
	helper.CheckErr("Failed Marshal line 159 controller/base", err)

	beego.Info(
		"RES_JOBID", helper.GetJobID(c),
		"RES_URL", c.Input.URL(),
		"RES_HEADER", string(headerByte),
		"RES_BODY", string(jsonByte),
	)
}

func filterErrorCode(arrErrType []structs.TypeError) []structs.FilterErrorCode {
	var filter []structs.FilterErrorCode
	for _, val := range arrErrType {
		filter = append(filter, structs.FilterErrorCode{
			Code:    val.Code,
			Message: val.Message,
		})
	}
	return filter
}
