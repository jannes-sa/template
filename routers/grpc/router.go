package grpc

import (
	ctrl "template/controllers/grpc"
	"template/helper/constant"
	pb "template/proto"
	"template/structs"
)

var (
	prefix = "/" + constant.GOAPP + "/" + constant.VERSION
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
		/*:STARTGRPC*/
		prefix + "/template": ctrl.RPCtrl,
		/*:ENDGRPC*/
	}
}
