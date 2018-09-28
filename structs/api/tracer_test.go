package api

import (
	"log"
	"reflect"
	"testing"
)

func TestTracerStruct(t *testing.T) {
	var h HeaderTracer
	h.XReqID = "XReqID"
	h.SpanID = "SpanID"
	h.ParSpanID = "ParSpanID"
	h.TraceID = "TraceID"

	v := reflect.ValueOf(&h).Elem()
	vt := v.Type()
	for i, n := 0, v.NumField(); i < n; i++ {
		tag := vt.Field(i).Tag.Get("header")
		log.Println(tag)

		xx := v.Field(i).Interface()
		log.Println(xx.(string))

		// v.Field(i).SetString("TEST")
	}

	log.Println(h)
	t.Fatal("X")
}
