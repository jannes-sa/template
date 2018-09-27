package http

import (
	"strconv"
	"strings"
	"template/helper"
	"template/helper/timetn"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

// Middleware - Middleware Struct
type Middleware struct{}

func (m *Middleware) log(c *context.Context) {
	jobID := helper.GetJobID(c)
	go func() {
		beego.Info(
			"REQ_JOBID", jobID,
			"REQ_URL", jobID, c.Input.URL(),
			"REQ_HEADER", jobID, helper.HeaderAll(c),
			"REQ_BODY",
			jobID,
			strings.Replace(
				string(c.Input.RequestBody),
				"\n", "", -1),
		)
	}()
}

func (m *Middleware) initialHeader(c *context.Context) {
	m.handleRoundTrip(c)
	m.handleReqID(c)
	m.handleJobID(c)
	m.setMessageID(c)
}

func (m *Middleware) handleRoundTrip(c *context.Context) {
	ms := timetn.Now().UnixNano() / int64(time.Millisecond)
	c.Input.SetData("x-roundtrip", strconv.FormatInt(ms, 10))
}

func (m *Middleware) handleReqID(c *context.Context) {
	reqID := c.Input.Header("x-request-id")
	c.Output.Header("x-request-id", reqID)
}

func (m *Middleware) handleJobID(c *context.Context) {
	jobID := c.Input.Header("x-job-id")
	newRequest := 0
	if jobID == "" {
		newRequest = 1
		c.Input.SetData("job-id", helper.GenJobID())
		c.Input.SetData("new_request", newRequest)
	} else {
		c.Input.SetData("job-id", jobID)
		c.Input.SetData("new_request", newRequest)
	}
}

func (m *Middleware) setMessageID(c *context.Context) {
	uuID := helper.GetUUID()
	c.Input.SetData("message-id", uuID)
}

// BeforeFunc - BeforeFunc
func BeforeFunc(c *context.Context) {
	var m Middleware

	m.initialHeader(c)
	m.log(c)
}

// AfterFunc to execute progress after response
func AfterFunc(c *context.Context) {

}
