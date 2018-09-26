package helper

import (
	"testing"
	"time"
	"txn/helper/constant"

	. "github.com/smartystreets/goconvey/convey"
)

func TestDebugElapsedTime(t *testing.T) {
	DebugElapsedTime("tes elapsed time", time.Now())
}

func TestPathLogFile(t *testing.T) {
	PathLogFile(constant.GOPATH)
}

func TestGenJobID(t *testing.T) {
	jobID := GenJobID()

	Convey("TestGenJobID \n", t, func() {
		Convey("jobID exists", func() {
			So(jobID, ShouldNotEqual, nil)
		})
	})
}
