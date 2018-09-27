package grpc

import (
	"context"
	"fmt"
	"time"

	"github.com/astaxie/beego"
	rpc "google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func StreamServerInterceptor() rpc.StreamServerInterceptor {

	return func(srv interface{}, stream rpc.ServerStream, info *rpc.StreamServerInfo, handler rpc.StreamHandler) error {
		startTime := time.Now()

		md, _ := metadata.FromIncomingContext(stream.Context())
		err := handler(srv, stream)
		md.Set("content-type", "application/grpc; charset=UTF-8")

		rpc.SendHeader(stream.Context(), md)

		dur := time.Since(startTime)

		beego.Info("StreamServerInterceptor => ", logInfo("Server", info, dur, md, nil, nil, err))

		return err
	}
}

func UnaryServerInterceptor() rpc.UnaryServerInterceptor {

	return func(ctx context.Context, req interface{}, info *rpc.UnaryServerInfo, handler rpc.UnaryHandler) (interface{}, error) {
		startTime := time.Now()
		md, _ := metadata.FromIncomingContext(ctx)
		resp, err := handler(ctx, req)

		md.Set("content-type", "application/grpc; charset=UTF-8")

		rpc.SendHeader(ctx, md)
		dur := time.Since(startTime)

		beego.Info("UnaryServerInterceptor => ", logInfo("Server", info, dur, md, req, resp, err))

		return resp, err
	}
}

func logInfo(
	kind string,
	info interface{},
	duration time.Duration,
	meta metadata.MD,
	req interface{},
	res interface{},
	err error,
) string {
	return fmt.Sprintf("[GRPC] <Kind:%v>, <Info:%+v>,  <Duration(ms):%v>,<MetaData:%+v> <Request:%+v>, <Response:%+v>, <Error:%v>",
		kind,
		info,
		duration.Nanoseconds()/1000000,
		meta,
		req,
		res,
		err,
	)
}

// ChainStreamServer creates a single interceptor out of a chain of many interceptors.
//
// Execution is done in left-to-right order, including passing of context.
// For example ChainUnaryServer(one, two, three) will execute one before two before three.
// If you want to pass context between interceptors, use WrapServerStream.
func ChainStreamServer(interceptors ...rpc.StreamServerInterceptor) rpc.StreamServerInterceptor {
	n := len(interceptors)

	if n > 1 {
		lastI := n - 1
		return func(srv interface{}, stream rpc.ServerStream, info *rpc.StreamServerInfo, handler rpc.StreamHandler) error {
			var (
				chainHandler rpc.StreamHandler
				curI         int
			)

			chainHandler = func(currentSrv interface{}, currentStream rpc.ServerStream) error {
				if curI == lastI {
					return handler(currentSrv, currentStream)
				}
				curI++
				err := interceptors[curI](currentSrv, currentStream, info, chainHandler)
				curI--
				return err
			}

			return interceptors[0](srv, stream, info, chainHandler)
		}
	}

	if n == 1 {
		return interceptors[0]
	}

	// n == 0; Dummy interceptor maintained for backward compatibility to avoid returning nil.
	return func(srv interface{}, stream rpc.ServerStream, _ *rpc.StreamServerInfo, handler rpc.StreamHandler) error {
		return handler(srv, stream)
	}
}

// ChainUnaryServer creates a single interceptor out of a chain of many interceptors.
//
// Execution is done in left-to-right order, including passing of context.
// For example ChainUnaryServer(one, two, three) will execute one before two before three, and three
// will see context changes of one and two.
func ChainUnaryServer(interceptors ...rpc.UnaryServerInterceptor) rpc.UnaryServerInterceptor {
	n := len(interceptors)

	if n > 1 {
		lastI := n - 1
		return func(ctx context.Context, req interface{}, info *rpc.UnaryServerInfo, handler rpc.UnaryHandler) (interface{}, error) {
			var (
				chainHandler rpc.UnaryHandler
				curI         int
			)

			chainHandler = func(currentCtx context.Context, currentReq interface{}) (interface{}, error) {
				if curI == lastI {
					return handler(currentCtx, currentReq)
				}
				curI++
				resp, err := interceptors[curI](currentCtx, currentReq, info, chainHandler)
				curI--
				return resp, err
			}

			return interceptors[0](ctx, req, info, chainHandler)
		}
	}

	if n == 1 {
		return interceptors[0]
	}

	// n == 0; Dummy interceptor maintained for backward compatibility to avoid returning nil.
	return func(ctx context.Context, req interface{}, _ *rpc.UnaryServerInfo, handler rpc.UnaryHandler) (interface{}, error) {
		return handler(ctx, req)
	}
}

// ChainUnaryClient creates a single interceptor out of a chain of many interceptors.
//
// Execution is done in left-to-right order, including passing of context.
// For example ChainUnaryClient(one, two, three) will execute one before two before three.
func ChainUnaryClient(interceptors ...rpc.UnaryClientInterceptor) rpc.UnaryClientInterceptor {
	n := len(interceptors)

	if n > 1 {
		lastI := n - 1
		return func(ctx context.Context, method string, req, reply interface{}, cc *rpc.ClientConn, invoker rpc.UnaryInvoker, opts ...rpc.CallOption) error {
			var (
				chainHandler rpc.UnaryInvoker
				curI         int
			)

			chainHandler = func(currentCtx context.Context, currentMethod string, currentReq, currentRepl interface{}, currentConn *rpc.ClientConn, currentOpts ...rpc.CallOption) error {
				if curI == lastI {
					return invoker(currentCtx, currentMethod, currentReq, currentRepl, currentConn, currentOpts...)
				}
				curI++
				err := interceptors[curI](currentCtx, currentMethod, currentReq, currentRepl, currentConn, chainHandler, currentOpts...)
				curI--
				return err
			}

			return interceptors[0](ctx, method, req, reply, cc, chainHandler, opts...)
		}
	}

	if n == 1 {
		return interceptors[0]
	}

	// n == 0; Dummy interceptor maintained for backward compatibility to avoid returning nil.
	return func(ctx context.Context, method string, req, reply interface{}, cc *rpc.ClientConn, invoker rpc.UnaryInvoker, opts ...rpc.CallOption) error {
		return invoker(ctx, method, req, reply, cc, opts...)
	}
}

// ChainStreamClient creates a single interceptor out of a chain of many interceptors.
//
// Execution is done in left-to-right order, including passing of context.
// For example ChainStreamClient(one, two, three) will execute one before two before three.
func ChainStreamClient(interceptors ...rpc.StreamClientInterceptor) rpc.StreamClientInterceptor {
	n := len(interceptors)

	if n > 1 {
		lastI := n - 1
		return func(ctx context.Context, desc *rpc.StreamDesc, cc *rpc.ClientConn, method string, streamer rpc.Streamer, opts ...rpc.CallOption) (rpc.ClientStream, error) {
			var (
				chainHandler rpc.Streamer
				curI         int
			)

			chainHandler = func(currentCtx context.Context, currentDesc *rpc.StreamDesc, currentConn *rpc.ClientConn, currentMethod string, currentOpts ...rpc.CallOption) (rpc.ClientStream, error) {
				if curI == lastI {
					return streamer(currentCtx, currentDesc, currentConn, currentMethod, currentOpts...)
				}
				curI++
				stream, err := interceptors[curI](currentCtx, currentDesc, currentConn, currentMethod, chainHandler, currentOpts...)
				curI--
				return stream, err
			}

			return interceptors[0](ctx, desc, cc, method, chainHandler, opts...)
		}
	}

	if n == 1 {
		return interceptors[0]
	}

	// n == 0; Dummy interceptor maintained for backward compatibility to avoid returning nil.
	return func(ctx context.Context, desc *rpc.StreamDesc, cc *rpc.ClientConn, method string, streamer rpc.Streamer, opts ...rpc.CallOption) (rpc.ClientStream, error) {
		return streamer(ctx, desc, cc, method, opts...)
	}
}

func UnaryClientInterceptor() rpc.UnaryClientInterceptor {

	return func(ctx context.Context, method string, req, reply interface{}, cc *rpc.ClientConn, invoker rpc.UnaryInvoker, opts ...rpc.CallOption) error {
		md, _ := metadata.FromOutgoingContext(ctx)
		err := invoker(ctx, method, req, reply, cc, opts...)
		beego.Info("Unary Client Interceptor => ", md)
		return err
	}
}

// StreamServerInterceptor returns a new streaming client interceptor that optionally logs the execution of external rpc calls.
func StreamClientInterceptor() rpc.StreamClientInterceptor {

	return func(ctx context.Context, desc *rpc.StreamDesc, cc *rpc.ClientConn, method string, streamer rpc.Streamer, opts ...rpc.CallOption) (rpc.ClientStream, error) {

		md, _ := metadata.FromOutgoingContext(ctx)
		clientStream, err := streamer(ctx, desc, cc, method, opts...)
		beego.Info("Stream Client Interceptor => ", md)
		return clientStream, err
	}
}
