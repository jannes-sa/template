package grpc

import (
	pb "template/proto"
	"template/structs"

	"github.com/astaxie/beego"
)

func RPCTest(
	in *pb.DoReq,
	errRPCCode *structs.TypeGRPCError,
	body *[]byte,
) {
	beego.Info(in.GetBody())
	*body = []byte(`{"tesResponse":"tesResponse"}`)
}

func RPCTestFailed(
	in *pb.DoReq,
	errRPCCode *structs.TypeGRPCError,
	body *[]byte,
) {
	var errCode []structs.TypeError
	structs.ErrorCode.UnexpectedError.String(&errCode)
	(*errRPCCode).Error = errCode

	beego.Info(in.GetBody())
	*body = []byte(`{"tesResponseFailed":"tesResponseFailed"}`)
}
