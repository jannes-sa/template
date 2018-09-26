package mq

import "testing"

func TestSendCommonMQ(t *testing.T) {
	var (
		post      = []byte(``)
		channel   = ""
		mq        = ""
		exchange  = ""
		eventName = ""
	)
	SendCommonMQ(post, channel, mq, exchange, eventName)
}
