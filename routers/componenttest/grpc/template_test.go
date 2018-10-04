package grpc

import (
	"encoding/json"
	"template/helper"
	"template/helper/timetn"
	"template/structs"
	structsAPI "template/structs/api"
	structsRPC "template/structs/api/grpc"
	"template/thirdparty/rpc"
	"testing"

	"github.com/astaxie/beego"
	. "github.com/smartystreets/goconvey/convey"
)

func TestTemplate(t *testing.T) {
	reqID := helper.GenJobID()

	var errorHeader structs.TypeGRPCError
	header := structsRPC.TypeHeaderRPC{
		ReqID:       reqID,
		Date:        timetn.Now(),
		ContentType: "application/grpc",
		RoundTrip:   "",
		Error:       errorHeader,
	}
	headerByte, _ := json.Marshal(header)

	var req structsRPC.ReqTest
	req.ID = 1
	req.Data = "requestdata"
	reqBy, _ := json.Marshal(req)

	var tracer structsAPI.HeaderTracer
	tracer.ParSpanID = "ParSpanID"
	tracer.SpanID = "SpanID"
	tracer.TraceID = "TraceID"
	tracer.XReqID = "XReqID"

	resp, err := rpc.SendGRPCComponentTest(
		prefix+"/template",
		host,
		reqBy,
		headerByte,
		reqID,
		tracer,
	)

	var resHeader structsRPC.TypeHeaderRPC
	json.Unmarshal(resp.Header, &resHeader)

	var resBody structsRPC.ResTest
	json.Unmarshal(resp.Body, &resBody)
	beego.Debug("resHeader => ", resHeader)
	beego.Debug("resBody => ", resBody)

	Convey("TestTemplate", t, func() {
		Convey("Should Success", func() {
			So(err, ShouldEqual, nil)
			So(len(resHeader.Error.Error), ShouldEqual, 0)
		})
	})
}
