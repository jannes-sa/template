package testComponent

import (
	"encoding/json"
	"template/helper"
	"template/helper/timetn"
	"template/structs"
	structsRPC "template/structs/api/grpc"
	"template/thirdparty/rpc"
	"testing"

	"github.com/astaxie/beego"
)

func TestComponentGRPCSuccess(t *testing.T) {
	reqID := helper.GenJobID()

	var errorHeader structs.TypeGRPCError
	header := structsRPC.TypeHeaderRPC{
		ReqID:       reqID,
		Date:        timetn.Now(),
		ContentType: "application/json",
		RoundTrip:   "",
		Error:       errorHeader,
	}
	headerByte, _ := json.Marshal(header)

	resp, err := rpc.SendGRPCComponentTest(
		"/rpcTest",
		"127.0.0.1:58080",
		[]byte(`
			{"date":"2018-05-16","report_id":1}
		`),
		headerByte,
		reqID,
	)

	beego.Debug(err)
	beego.Debug(string(resp.Header))
	beego.Debug(string(resp.Body))
}

func TestComponentGRPC404(t *testing.T) {
	reqID := helper.GenJobID()

	var errorHeader structs.TypeGRPCError
	header := structsRPC.TypeHeaderRPC{
		ReqID:       reqID,
		Date:        timetn.Now(),
		ContentType: "application/json",
		RoundTrip:   "",
		Error:       errorHeader,
	}
	headerByte, _ := json.Marshal(header)

	resp, err := rpc.SendGRPCComponentTest(
		"/404",
		"127.0.0.1:58080",
		[]byte(`
			{"date":"2018-05-16","report_id":1}
		`),
		headerByte,
		reqID,
	)

	beego.Debug(err)
	beego.Debug(string(resp.Header))
	beego.Debug(string(resp.Body))
}

func TestComponentGRPCFailed(t *testing.T) {
	reqID := helper.GenJobID()

	var errorHeader structs.TypeGRPCError
	header := structsRPC.TypeHeaderRPC{
		ReqID:       reqID,
		Date:        timetn.Now(),
		ContentType: "application/json",
		RoundTrip:   "",
		Error:       errorHeader,
	}
	headerByte, _ := json.Marshal(header)

	resp, err := rpc.SendGRPCComponentTest(
		"/rpcFailed",
		"127.0.0.1:58080",
		[]byte(`
			{"date":"2018-05-16","report_id":1}
		`),
		headerByte,
		reqID,
	)

	beego.Debug(err)
	beego.Debug(string(resp.Header))
	beego.Debug(string(resp.Body))
}
