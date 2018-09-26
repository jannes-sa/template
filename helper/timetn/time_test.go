package timetn

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	//"github.com/astaxie/beego"
	"time"
)

func TestParseTimeLocToString(t *testing.T) {

	dateTimeConvertChanPostDate, err := ParseTimeLocToString(
		time.RFC3339,
		"2018-06-06T23:00:00+06:00",
		"Asia/Bangkok",
	)

	Convey("TestParseTimeLocToString", t, func() {
		Convey("convert timezone correct", func() {
			So(err, ShouldBeNil)
			So(dateTimeConvertChanPostDate, ShouldEqual, "2018-06-07T00:00:00+07:00")
		})
	})

}

func TestParseTimeLocToString_plusZero(t *testing.T) {

	dateTimeConvertChanPostDate, err := ParseTimeLocToString(
		time.RFC3339,
		"2018-06-06T17:00:00+00:00",
		"Asia/Bangkok",
	)

	Convey("TestParseTimeLocToString_plusZero", t, func() {
		Convey("convert timezone correct", func() {
			So(err, ShouldBeNil)
			So(dateTimeConvertChanPostDate, ShouldEqual, "2018-06-07T00:00:00+07:00")
		})
	})

}
