package api

import (
	"encoding/hex"
	"math/rand"
	"reflect"

	ctxGO "context"

	"github.com/astaxie/beego/context"
	"google.golang.org/grpc/metadata"
)

var tag string = "header"

type (
	// HeaderTracer - Header for Jaeger Tracer
	HeaderTracer struct {
		XReqID    string `header:"x-request-id"`
		SpanID    string `header:"x-b3-spanid"`
		ParSpanID string `header:"x-b3-parentspanid"`
		TraceID   string `header:"x-b3-traceid"`
	}
)

func (h *HeaderTracer) extractStruct() (
	v reflect.Value,
	vt reflect.Type,
) {
	v = reflect.ValueOf(h).Elem()
	vt = v.Type()

	return
}

// HTTPGetHeaderTrace - Get Header Tracer HTTP
func (h *HeaderTracer) HTTPGetHeaderTrace(ctx *context.Context) {
	v, vt := h.extractStruct()
	for i, n := 0, v.NumField(); i < n; i++ {
		v.Field(i).SetString(
			ctx.Input.Header(
				vt.Field(i).Tag.Get(tag),
			),
		)
	}

	manipulateSpanIDTracer(h)
}

// HTTPSetHeaderTrace - Set Header Tracer HTTP
func (h *HeaderTracer) HTTPSetHeaderTrace(ctx *context.Context) {
	v, vt := h.extractStruct()
	for i, n := 0, v.NumField(); i < n; i++ {
		ctx.Output.Header(
			vt.Field(i).Tag.Get(tag),
			v.Field(i).Interface().(string),
		)
	}
}

// GRPCGetHeaderTrace - Get Header Tracer GRPC
func (h *HeaderTracer) GRPCGetHeaderTrace(md metadata.MD) {
	v, vt := h.extractStruct()
	for i, n := 0, v.NumField(); i < n; i++ {
		v.Field(i).SetString(
			md[vt.Field(i).Tag.Get(tag)][0],
		)
	}

	manipulateSpanIDTracer(h)
}

// GRPCSetHeaderTrace - Set Header Tracer GRPC
func (h *HeaderTracer) GRPCSetHeaderTrace(
	ctx *ctxGO.Context,
) {
	v, vt := h.extractStruct()
	for i, n := 0, v.NumField(); i < n; i++ {
		*ctx = metadata.AppendToOutgoingContext(
			*ctx,
			vt.Field(i).Tag.Get(tag),
			v.Field(i).Interface().(string),
		)
	}
}

func manipulateSpanIDTracer(h *HeaderTracer) {
	hex16Str := hex16Encode()
	(*h).ParSpanID = (*h).SpanID
	(*h).SpanID = hex16Str
}

func hex16Encode() string {
	src := []byte(randStringBytes(8))
	dst := make([]byte, hex.EncodedLen(len(src)))
	hex.Encode(dst, src)
	return string(dst)
}

const letterBytes = "abcdefghijklmnopqrstuvwxyz1234567890"

func randStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
