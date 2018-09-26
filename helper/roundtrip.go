package helper

import (
	"strconv"
	"strings"
	"template/helper/timetn"
	"template/structs"
	"time"

	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

// Calculate ...
func Calculate(times *[]string, start time.Time) {
	now := timetn.Now()
	rountrip := now.Sub(start)

	*times = append(*times,
		strconv.FormatInt(rountrip.Nanoseconds()/(1*1000*1000), 10))
}

// SetXPerformance ...
func SetXPerformance(ctx *context.Context, title string, times []string) {
	ctx.Output.Header("x-performance", title+strings.Join(times, "|"))
}

// GetProfileRoundtrip ...
func GetProfileRoundtrip(data string) (roundtrip string) {
	var resHeader structs.ResHTTPHeader
	beego.Debug(data)
	err := json.Unmarshal([]byte(data), &resHeader)

	CheckErr("GetProfileRoundtrip", err)

	if err == nil {
		roundtrip = resHeader.XRoundTrip
	}

	return
}
