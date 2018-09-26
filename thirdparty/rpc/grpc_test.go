package rpc

import (
	"log"
	"testing"
)

func TestGrpc(t *testing.T) {
	var (
		route      = ""
		serverAddr = "txn@127.0.0.1:8084"
		body       = []byte(``)
		header     = []byte(``)
		reqID      = ``
	)

	resRPC, err := SendGRPC(route, serverAddr, body, header, reqID)

	log.Println(err)
	log.Println(string(resRPC.Body))
	log.Println(string(resRPC.Header))
}
