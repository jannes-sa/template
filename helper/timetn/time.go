package timetn

import (
	"strconv"
	"template/helper/constant"
	"time"

	"github.com/astaxie/beego"
)

var addedTimeNow = 0

const (
	failedLoadLocMsg = "failed load location"
)

// Now replace helper.Now()
func Now() time.Time {
	if constant.TZ == "" {
		constant.TZ = "Asia/Bangkok"
	}
	loc, err := time.LoadLocation(constant.TZ)
	if err != nil {
		beego.Warning(failedLoadLocMsg, err)
	}

	return time.Now().AddDate(0, 0, addedTimeNow).In(loc)
}

// NowLoc - Time Now With Spesific Location Timezone
func NowLoc(timeLoc string) (
	dateTime time.Time,
	err error,
) {
	// "Asia/Bangkok"
	loc, err := time.LoadLocation(timeLoc)
	if err != nil {
		beego.Warning(failedLoadLocMsg, err)
	}
	dateTime = Now().In(loc)
	return
}

// ParseTimeLocToString - ParseTimeLocToString
func ParseTimeLocToString(
	layout string,
	valueTime string,
	timeLoc string,
) (
	dateTimeStr string,
	err error,
) {

	dateTime, err := ParseTimeLoc(layout, valueTime, timeLoc)
	if err != nil {
		beego.Warning("failed @ParseTimeLoc", err)
		return
	}

	dateTimeStr = dateTime.Format(layout)
	return
}

// ParseTimeLoc - ParseTimeLoc
func ParseTimeLoc(
	layout string,
	valueTime string,
	timeLoc string,
) (
	dateTimeRes time.Time,
	err error,
) {

	// Asia/Bangkok
	loc, err := time.LoadLocation(timeLoc)
	if err != nil {
		beego.Warning(failedLoadLocMsg, err)
		return
	}
	dateTime, err := time.Parse(layout, valueTime)
	if err != nil {
		beego.Warning("failed load parseInLocation", err)
		return
	}

	dateTimeRes = dateTime.In(loc)
	return
}

// SetTimeDateWithSpesificDate - SetTimeDateWithSpesificDate
func SetTimeDateWithSpesificDate(y int, m time.Month, d int) {
	if y != 0 && m != 0 && d != 0 {
		tm, err := time.LoadLocation("Asia/Bangkok")

		if err != nil {
			beego.Warning("error get set time", err)
		}

		now := time.Now()
		wayback := time.Date(y, m, d, 0, 0, 0, 0, tm)
		waynow := time.Date(
			now.Year(),
			now.Month(),
			now.Day(), 0, 0, 0, 0, tm)

		timeSubHours := int(wayback.Sub(waynow).Hours() / 24)
		addedTimeNow = timeSubHours
	}
}

// GetSetTimeDate ...
func GetSetTimeDate(num int) {
	addedTimeNow = num
}

// CreateTNDate Create t-n Date
func CreateTNDate(num int) string {
	now := Now()
	tn := now.AddDate(0, 0, num)
	d := tn.Day()
	m := int(tn.Month())
	y := tn.Year()

	tdate := strconv.Itoa(y) + `-` + filterDM(m) + `-` + filterDM(d)

	return tdate
}

// ParseDateToTimeTNDate ...
func ParseDateToTimeTNDate(num int) time.Time {
	now := Now()
	tn := now.AddDate(0, 0, num)
	d := tn.Day()
	m := int(tn.Month())
	y := tn.Year()

	tdate := strconv.Itoa(y) + `-` + filterDM(m) + `-` + filterDM(d)
	tdate += `T00:00:00.000Z`

	t, err := time.Parse(time.RFC3339, tdate)
	if err != nil {
		beego.Warning(err)
	}

	return t
}

// GetMNDate ...
func GetMNDate(num int) string {
	now := Now()
	tn := now.AddDate(0, 0, num)
	d := tn.Day()
	m := int(tn.Month())
	y := tn.Year()

	tdate := strconv.Itoa(y) + filterDM(m) + filterDM(d)

	return tdate
}

func filterDM(m int) string {
	mstr := strconv.Itoa(m)
	if len(mstr) < 2 {
		return `0` + mstr
	}
	return mstr
}
