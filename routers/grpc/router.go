package grpc

import (
	ctrl "template/controllers/grpc"
	pb "template/proto"
	"template/structs"
)

type fnRouteRPC func(
	*pb.DoReq,
	*structs.TypeGRPCError,
	*[]byte,
)

var routeMap map[string]fnRouteRPC

func init() {
	Router()
}

func Router() {
	routeMap = map[string]fnRouteRPC{
		"/rpcTest":   ctrl.RPCTest,
		"/rpcFailed": ctrl.RPCTestFailed,
	}
}
