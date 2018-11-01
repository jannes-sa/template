package helper

import (
	"strconv"
	"strings"
	"template/helper/constant"
	"template/structs"
	"testing"

	"github.com/astaxie/beego"
	. "github.com/smartystreets/goconvey/convey"
)

func TestGenJobID(t *testing.T) {
	jobID := GenJobID()

	Convey("TestGenJobID \n", t, func() {
		Convey("GenJobID exists", func() {
			So(jobID, ShouldNotEqual, nil)
		})
	})
}

func TestFilterStrMonthDay(t *testing.T) {
	value := filterStrMonthDay("1")
	beego.Warning(value)
	Convey("TestFilterStrMonthDay \n", t, func() {
		Convey("FilterStrMonthDay exists", func() {
			So(value, ShouldNotEqual, nil)
			So(value, ShouldEqual, "01")
		})
	})
}

func TestRandom(t *testing.T) {
	min := 0
	max := 10
	value := Random(min, max)
	res, _ := strconv.Atoi(value)
	Convey("TestRandom \n", t, func() {
		Convey("TestRandom exists", func() {
			So(res, ShouldNotEqual, nil)
			So(res, ShouldBeLessThan, max)
			So(res, ShouldBeGreaterThan, min)
		})
	})
}

func TestGetSecureRandomNumber(t *testing.T) {
	length := 10
	value := getSecureRandomNumber(length)

	Convey("TestGetSecureRandomNumber \n", t, func() {
		Convey("GetSecureRandomNumber exists", func() {
			So(value, ShouldNotEqual, nil)
			So(len(value), ShouldEqual, length)
		})
	})
}

func TestCryptoRandSecure(t *testing.T) {
	max := 10
	value := cryptoRandSecure(int64(max))

	Convey("TestCryptoRandSecure \n", t, func() {
		Convey("CryptoRandSecure exists", func() {
			So(value, ShouldNotEqual, nil)
			So(value, ShouldBeLessThan, max)
		})
	})
}

func TestGenCRC32(t *testing.T) {
	value := GenCRC32()
	Convey("TestGenCRC32 \n", t, func() {
		Convey("GenCRC32 exists", func() {
			So(value, ShouldNotEqual, nil)
		})
	})
}

func TestRangeRandomNumber(t *testing.T) {
	max := 10
	min := 0
	value := RangeRandomNumber(int64(min), int64(max))
	Convey("RangeRandomNumber \n", t, func() {
		Convey("RangeRandomNumber exists", func() {
			So(value, ShouldNotEqual, nil)
			So(value, ShouldBeLessThan, max)
			So(value, ShouldBeGreaterThan, min)
		})
	})
}

func TestPathLogFile(t *testing.T) {
	logfile := PathLogFile(constant.GOPATH)
	Convey("TestPathLogFile \n", t, func() {
		Convey("PathLogFile exists", func() {
			So(logfile, ShouldNotEqual, nil)
		})
	})
}

func TestHeaderAll(t *testing.T) {
	ctx := ScaffoldContext()
	header := HeaderAll(ctx)

	Convey("TestGenJobID \n", t, func() {
		Convey("jobID exists", func() {
			So(header, ShouldNotEqual, nil)
		})
	})
}

func TestGetHeader(t *testing.T) {
	headerAll := `{"Accept":"application/json","Accept-Encoding":"UTF-8","Accept-Language":"en/id","Cache-Control":"no-cache","Connection":"keep-alive","Content-Length":"296","Content-Type":"application/json","Datetime":"2017-09-19T10:59:47.305411285+07:00","Postman-Token":"994fccb9-2673-4f9c-b400-7b5aeb5a5cdd","User-Agent":"POSTMAN","X-Caller-Domain":"CUS DOMAIN","X-Caller-Service":"CUS","X-Channel":"MBL","X-Device":"ANDROID DDD","X-Job-Id":"","X-Real-Ip":"192.168.99.100","X-Request-Id":"test"}`
	getHeader := GetHeader(headerAll)

	Convey("TestGetHeader \n", t, func() {
		Convey("header exists", func() {
			So(getHeader, ShouldNotEqual, nil)
		})
	})
}

func TestContextStruct(t *testing.T) {
	c := ScaffoldContext()
	contextStruct := ContextStruct(c)

	Convey("TestContextStruct \n", t, func() {
		Convey("context struct exists", func() {
			So(contextStruct.HeaderTracer, ShouldNotEqual, nil)
			So(contextStruct.HeaderAll, ShouldNotEqual, nil)
			So(contextStruct.Header, ShouldNotEqual, nil)
			So(contextStruct.JobID, ShouldNotEqual, nil)
		})
	})
}

func TestScaffoldContext(t *testing.T) {
	ctx := ScaffoldContext()

	Convey("TestScaffoldContext \n", t, func() {
		Convey("context struct exists", func() {
			So(ctx, ShouldNotEqual, nil)
		})
	})
}

func TestIsValidUUID(t *testing.T) {
	isValid := IsValidUUID("1234567890")

	Convey("TestIsValidUUID \n", t, func() {
		Convey("isValid exists", func() {
			So(isValid, ShouldNotEqual, nil)
		})
	})
}

func TestSetHeaderRPC(t *testing.T) {
	befMsStr := "1234"
	reqID := "1234"
	var errRPCCode structs.TypeGRPCError

	headerRPC := SetHeaderRPC(befMsStr, reqID, errRPCCode)
	result := string(headerRPC[:])
	resultSplit := strings.Split(result, ",")
	test := `"content-type":"application/grpc"`

	Convey("TestSetHeaderRPC \n", t, func() {
		Convey("header RPC exists", func() {
			So(headerRPC, ShouldNotEqual, nil)
			So(len(resultSplit), ShouldEqual, 5)
			So(resultSplit[2], ShouldEqual, test)
		})
	})
}

func TestIsNewRequest(t *testing.T) {
	ctx := ScaffoldContext()
	isNewRequest := IsNewRequest(ctx)

	Convey("TestIsNewRequest \n", t, func() {
		Convey("is new request exists", func() {
			So(isNewRequest, ShouldNotEqual, nil)
			So(isNewRequest, ShouldEqual, false)
		})
	})
}

func TestGetReqID(t *testing.T) {
	ctx := ScaffoldContext()
	getReqID := GetReqID(ctx)

	Convey("TestGetReqID \n", t, func() {
		Convey("req ID exists", func() {
			So(getReqID, ShouldNotEqual, nil)
		})
	})
}

func TestGetJobID(t *testing.T) {
	ctx := ScaffoldContext()
	getJobID := GetJobID(ctx)

	Convey("TestGetJobID \n", t, func() {
		Convey("job ID exists", func() {
			So(getJobID, ShouldNotEqual, nil)
			So(getJobID, ShouldEqual, "null")
		})
	})
}

func TestGetMessageID(t *testing.T) {
	ctx := ScaffoldContext()
	getMessageID := GetMessageID(ctx)

	Convey("TestGetMessageID \n", t, func() {
		Convey("message ID exists", func() {
			So(getMessageID, ShouldNotEqual, nil)
			So(getMessageID, ShouldEqual, "null")
		})
	})
}

func TestGetRoundTrip(t *testing.T) {
	ctx := ScaffoldContext()
	getRoundTrip := GetRoundTrip(ctx)

	Convey("TestGetRoundTrip \n", t, func() {
		Convey("round trip exists", func() {
			So(getRoundTrip, ShouldNotEqual, nil)
		})
	})
}

func TestGetRoundTripInternal(t *testing.T) {
	roundTripInternal := GetRoundTripInternal(60)

	Convey("TestGetRoundTripInternal \n", t, func() {
		Convey("round trip internal exists", func() {
			So(roundTripInternal, ShouldNotEqual, nil)
		})
	})
}

func TestGetRqBody(t *testing.T) {
	ctx := ScaffoldContext()
	getRqBody := GetRqBody(ctx)

	Convey("TestGetRqBody \n", t, func() {
		Convey("req body exists", func() {
			So(getRqBody, ShouldNotEqual, nil)
		})
	})
}

func TestGetRqBodyRev(t *testing.T) {
	var errRPCCode []structs.TypeError
	ctx := ScaffoldContext()
	getRqBodyRev := GetRqBodyRev(ctx, &errRPCCode)

	Convey("TestGetRqBodyRev \n", t, func() {
		Convey("req body rev exists", func() {
			So(getRqBodyRev, ShouldNotEqual, nil)
		})
	})
}
