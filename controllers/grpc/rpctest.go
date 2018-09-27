package grpc

import (
	pb "template/proto"
	"template/structs"

	"github.com/astaxie/beego"
)

// RPCTest - Controller RPC for testing only Success
func RPCTest(
	in *pb.DoReq,
	errRPCCode *structs.TypeGRPCError,
	body *[]byte,
) {
	beego.Info(in.GetBody())
	beego.Info(errRPCCode)
	*body = []byte(`{"tesResponse":"tesResponse"}`)
}

// RPCTestFailed - Controller RPC for testing only Failed
func RPCTestFailed(
	in *pb.DoReq,
	errRPCCode *structs.TypeGRPCError,
	body *[]byte,
) {
	var errCode []structs.TypeError
	structs.ErrorCode.UnexpectedError.String(&errCode)
	(*errRPCCode).Error = errCode

	beego.Info(in.GetBody())
	beego.Info(errRPCCode)
	*body = []byte(`{"tesResponseFailed":"tesResponseFailed"}`)
}
