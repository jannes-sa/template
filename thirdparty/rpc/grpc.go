package rpc

import (
	"strconv"
	"time"

	"template/helper"
	"template/helper/constant"
	"template/helper/timetn"

	middleware "template/routers/grpc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"

	structsAPI "template/structs/api"
	structsRPCAPI "template/structs/api/grpc"

	"context"

	pb "template/proto"
)

var (
	tls    = false
	caFile = constant.GOPATH + "/src/" + constant.GOAPP + "/proto/pem/ca.pem"

	serverHostOverride = string("x.test.youtube.com")
	maxTimeout         = time.Duration(25)
)

type connection struct {
	Client *grpc.ClientConn
}

// Dial ...
func (c *connection) Dial(
	serverAddr string,
) (err error) {
	var secureRPC grpc.DialOption
	if !tls {
		secureRPC = grpc.WithInsecure()
	} else {
		var sn string = serverHostOverride
		var creds credentials.TransportCredentials
		if caFile != "" {
			creds, err = credentials.NewClientTLSFromFile(caFile, sn)
			if err != nil {
				grpclog.Fatalf("Failed to create TLS credentials %v", err)
				return err
			}
		} else {
			creds = credentials.NewClientTLSFromCert(nil, sn)
		}
		secureRPC = grpc.WithTransportCredentials(creds)
	}

	conn, err := grpc.Dial(serverAddr,
		secureRPC,
		grpc.WithUnaryInterceptor(
			middleware.ChainUnaryClient(
				middleware.UnaryClientInterceptor(),
			),
		),
		grpc.WithStreamInterceptor(
			middleware.ChainStreamClient(
				middleware.StreamClientInterceptor(),
			),
		),
	)

	helper.CheckErr("Error Dial Call gRPC", err)
	c.Client = conn

	return err
}

// SendGRPC ...
func SendGRPC(
	route string,
	serverAddr string,
	body []byte,
	header []byte,
	reqID string,
	tracer structsAPI.HeaderTracer,
) (
	structsRPCAPI.TypeResponseRPC,
	error,
) {

	var respRPC structsRPCAPI.TypeResponseRPC
	var err error

	if constant.GOENV == constant.DEVCI {
		// structs.MappingMockRPC(route, body, &respRPC)
		return respRPC, err
	}

	respRPC, err = sendRPC(
		route,
		serverAddr,
		body,
		header,
		reqID,
		tracer,
	)

	return respRPC, err
}

// SendGRPCComponent ...
func SendGRPCComponentTest(
	route string,
	serverAddr string,
	body []byte,
	header []byte,
	reqID string,
	tracer structsAPI.HeaderTracer,
) (
	structsRPCAPI.TypeResponseRPC,
	error,
) {

	var respRPC structsRPCAPI.TypeResponseRPC
	var err error

	respRPC, err = sendRPC(
		route,
		serverAddr,
		body,
		header,
		reqID,
		tracer,
	)

	return respRPC, err
}

func sendRPC(
	route string,
	serverAddr string,
	body []byte,
	header []byte,
	reqID string,
	tracer structsAPI.HeaderTracer,
) (
	structsRPCAPI.TypeResponseRPC,
	error,
) {

	ms := timetn.Now().UnixNano() / int64(time.Millisecond)

	var respRPC structsRPCAPI.TypeResponseRPC
	var err error

	var client connection
	err = client.Dial(serverAddr)
	if err != nil {
		return respRPC, err
	}

	dial := client.Client
	defer func() {
		errC := client.Client.Close()
		helper.CheckErr("Failed Close gRPC", errC)
	}()

	c := pb.NewRestfulClient(dial)
	ctx, cancel := context.WithTimeout(context.Background(), maxTimeout*time.Second)
	defer cancel()

	tracer.GRPCSetHeaderTrace(&ctx)

	var md metadata.MD
	r, err := c.Echo(
		ctx,
		&pb.DoReq{
			Header: header,
			Body:   body,
			Route:  route,
			Round:  strconv.FormatInt(ms, 10),
			ReqID:  reqID,
		},
		grpc.Header(&md),
	)

	helper.CheckErr("Error Call Proto Method", err)
	if err != nil {
		return respRPC, err
	}

	respRPC.Header = r.Header
	respRPC.Body = r.Body

	var respmeta structsAPI.HeaderTracer
	respmeta.GRPCGetHeaderTrace(md)
	respRPC.Metadata = respmeta

	return respRPC, err
}
