package grpc

import (
	"net"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/astaxie/beego"

	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"

	"template/helper"
	"template/helper/constant"
	pb "template/proto"
	"template/structs"
)

var tls bool = false
var certFile = constant.GOPATH + "/src/" + constant.GOAPP + "/proto/pem/server1.pem"
var keyFile = constant.GOPATH + "/src/" + constant.GOAPP + "/proto/pem/server1.key"

type routeRPC struct{}

func (s *routeRPC) Echo(ctx context.Context, in *pb.DoReq) (*pb.DoResp, error) {
	var errRPCCode structs.TypeGRPCError
	var header []byte
	var body []byte

	if f, ok := routeMap[in.Route]; ok {
		f(in, &errRPCCode, &body)
	} else {
		s.NotFound(&errRPCCode)
	}

	header = helper.SetHeaderRPC(in.Round, in.ReqID, errRPCCode)
	return &pb.DoResp{Header: header, Body: body}, nil
}

// NotFound ...
func (s *routeRPC) NotFound(
	errRPCCode *structs.TypeGRPCError,
) error {
	err := status.Error(codes.NotFound, "NotFound")
	(*errRPCCode).Error = append((*errRPCCode).Error, structs.TypeError{
		Code:    "404",
		Case:    "Route Not Found",
		Message: "Route Not Found",
	})

	return err
}

type connection struct {
	Server *grpc.Server
}

func (c *connection) NewServer() error {
	if tls {
		creds, err := credentials.NewServerTLSFromFile(
			certFile,
			keyFile,
		)
		if err != nil {
			grpclog.Fatalf("Failed to generate credentials %v", err)
			return err
		}
		secureRPC := grpc.Creds(creds)

		server := grpc.NewServer(
			secureRPC,
			grpc.StreamInterceptor(ChainStreamServer(
				StreamServerInterceptor(),
			)),
			grpc.UnaryInterceptor(ChainUnaryServer(
				UnaryServerInterceptor(),
			)),
		)

		c.Server = server
		return nil
	}

	server := grpc.NewServer(
		grpc.StreamInterceptor(ChainStreamServer(
			StreamServerInterceptor(),
		)),
		grpc.UnaryInterceptor(ChainUnaryServer(
			UnaryServerInterceptor(),
		)),
	)

	c.Server = server
	return nil

}

// CreateGrpcServer ...
func CreateGrpcServer(portTemp string, customValidate ...string) {

	port := ":5" + portTemp
	if portTemp == "" {
		port = ":5" + strconv.Itoa(constant.APPPORT)
	}

	lis, errRPC := net.Listen("tcp", port)
	helper.CheckErr("Error Connect gRPC Server{}", errRPC)

	var server connection
	server.NewServer()
	s := server.Server
	pb.RegisterRestfulServer(s, &routeRPC{})
	reflection.Register(s)

	go func() {
		beego.Info("[*] gRPC Server Ready " + port)
		errGrpc := s.Serve(lis)
		helper.CheckErr("Error Connect gRPC to Serve()", errGrpc)
	}()

	if len(customValidate) > 0 {
		if customValidate[0] == "test" {
			return
		}
	}

	shutdown := make(chan os.Signal)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)
	<-shutdown
	s.GracefulStop()
	beego.Critical("Shutdown Gracefully gRPC")

}
