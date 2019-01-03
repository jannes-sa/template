package grpc

import (
	"encoding/json"
	"template/helper"
	pb "template/proto"
	"template/structs"
	rpcStructs "template/structs/api/grpc"
)

// RPCtrlTemplate - RPCtrlTemplate Controllers
func RPCtrl(
	in *pb.DoReq,
	errRPCCode *structs.TypeGRPCError,
	body *[]byte,
) {

	var (
		req rpcStructs.ReqTest
		res rpcStructs.ResTest
	)

	err := json.Unmarshal(in.GetBody(), &req)
	if err != nil {
		helper.CheckErr("failed unmarshal @RPCtrlTemplate", err)
		structs.ErrorCode.UnexpectedError.String(&errRPCCode.Error)
		return
	}

	res.ID = req.ID
	res.Res = "response"
	resBy, err := json.Marshal(res)
	if err != nil {
		helper.CheckErr("failed marshal &RPCtrlTemplate", err)
		structs.ErrorCode.UnexpectedError.String(&errRPCCode.Error)
		return
	}

	*body = resBy
}
