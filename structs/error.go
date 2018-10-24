package structs

import (
	"log"
	"reflect"
	"strings"
	c "template/helper/constant"
)

const (
	fieldCall = `{field}`
)

var strs = make(map[reflect.Type]map[int]map[string]string)

func register(t reflect.Type, v int, s map[string]string) {
	m, ok := strs[t]
	if !ok {
		m = make(map[int]map[string]string)
		strs[t] = m
	}
	m[v] = s
}

func getStr(e interface{}) map[string]string {
	t := reflect.TypeOf(e)
	v := int(reflect.ValueOf(e).Int())
	str := strs[t][v]

	return str
}

func initEnums(estr interface{}) {
	v := reflect.ValueOf(estr).Elem()
	vt := v.Type()
	for i, n := 0, v.NumField(); i < n; i++ {
		f := v.Field(i)
		maping := map[string]string{
			"code": vt.Field(i).Tag.Get("code"),
			"case": vt.Field(i).Tag.Get("case"),
			"msg":  vt.Field(i).Tag.Get("msg"),
		}
		register(f.Type(), i, maping)
		f.SetInt(int64(i))
	}
}

// Code is
type Code int

// String convert
func (e Code) String(errCode *[]TypeError, customErrorMessage ...string) {
	str := getStr(e)
	strMsg := str["msg"]
	if len(customErrorMessage) > 0 {
		strMsg += " " + strings.Join(customErrorMessage, ", ")
	}

	*errCode = append(*errCode, TypeError{
		Code:    c.PREFIXERRCODE + str["code"],
		Case:    str["case"],
		Message: strMsg,
	})
}

// GetString ...
func (e Code) GetString() (string, string, string) {
	str := getStr(e)
	return c.PREFIXERRCODE + str["code"], str["case"], str["msg"]
}

// GetCodeError ...
func GetCodeError(code []string, errCode *[]TypeError) {
	v := reflect.ValueOf(ErrorCode)
	vt := v.Type()
	for _, val := range code {
		filterMissingField(val, errCode)
		filterMismatchType(val, errCode)
	}

	for i, n := 0, v.NumField(); i < n; i++ {
		if contain(code, vt.Field(i).Tag.Get("code")) {
			*errCode = append(*errCode, TypeError{
				Code:    c.PREFIXERRCODE + vt.Field(i).Tag.Get("code"),
				Case:    vt.Field(i).Tag.Get("case"),
				Message: vt.Field(i).Tag.Get("msg"),
			})
		}
	}

}

func filterMissingField(val string, errCode *[]TypeError) {
	misField := strings.Contains(val, "missing_field")
	if misField {
		codes, cases, msgs := ErrorCode.MissingField.GetString()

		arrMisField := strings.Split(val, `|`)
		msgs = strings.Replace(msgs, fieldCall, arrMisField[1], -1)
		*errCode = append(*errCode, TypeError{
			Code:    codes,
			Case:    cases,
			Message: msgs,
		})
	}
}

func filterMismatchType(val string, errCode *[]TypeError) {
	if len(*errCode) == 0 {
		codeStr, casesStr, msgsStr := ErrorCode.ErorTypeString.GetString()
		codeNum, casesNum, msgsNum := ErrorCode.ErorTypeNumeric.GetString()
		codeTim, casesTim, msgsTim := ErrorCode.ErorTypeTime.GetString()

		if strings.Contains(val, "mismatch_type") {
			arrType := strings.Split(val, `|`)
			if arrType[1] == "string" {
				msgFilter := strings.Replace(msgsStr, fieldCall, arrType[2], -1)
				*errCode = append(*errCode, TypeError{
					Code:    codeStr,
					Case:    casesStr,
					Message: msgFilter,
				})
			} else if arrType[1] == "int" || arrType[1] == "int64" || arrType[1] == "float64" {
				msgFilter := strings.Replace(msgsNum, fieldCall, arrType[2], -1)
				*errCode = append(*errCode, TypeError{
					Code:    codeNum,
					Case:    casesNum,
					Message: msgFilter,
				})
			} else if arrType[1] == "time" {
				msgFilter := strings.Replace(msgsTim, fieldCall, arrType[2], -1)
				*errCode = append(*errCode, TypeError{
					Code:    codeTim,
					Case:    casesTim,
					Message: msgFilter,
				})
			}
		}
	}
}

// GetCodeErrorWithParseMsg ...
func GetCodeErrorWithParseMsg(code []string, errCode *[]TypeError, st interface{}) {
	v := reflect.ValueOf(ErrorCode)
	vt := v.Type()
	for i, n := 0, v.NumField(); i < n; i++ {
		if contain(code, vt.Field(i).Tag.Get("code")) {

			codeTag := vt.Field(i).Tag.Get("code")
			caseTag := vt.Field(i).Tag.Get("case")
			msgTag := vt.Field(i).Tag.Get("msg")

			vv := reflect.ValueOf(st)
			vvtt := vv.Type()
			for ii, nn := 0, vv.NumField(); ii < nn; ii++ {
				msgParse := vvtt.Field(ii).Tag.Get("msgparse")
				if msgParse != "" {
					msgParseSplit := strings.Split(msgParse, "=")
					log.Println(msgParseSplit)
					msgTag = strings.Replace(msgTag, msgParseSplit[0], msgParseSplit[1], -1)
				}
			}

			*errCode = append(*errCode, TypeError{
				Code:    c.PREFIXERRCODE + codeTag,
				Case:    caseTag,
				Message: msgTag,
			})
		}
	}
}
func contain(slice []string, item string) bool {
	set := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		set[s] = struct{}{}
	}

	_, ok := set[item]
	return ok
}

// ErrorCode Struct
var ErrorCode struct {

	// Common Format Error //
	UnexpectedError          Code `code:"0000" case:"Unexpected error" msg:"Unexpected error"`
	RPCInvalid               Code `code:"0001" case:"Connection RPC Invalid" msg:"Connection RPC Invalid"`
	ServiceNotAllowedToBatch Code `code:"0002" case:"Service Not Allowed to Batch" msg:"Service Not Allowed to Batch"`
	URLFilterInvalid         Code `code:"0003" case:"URL Filter Invalid" msg:"URL Filter Invalid"`

	RequestMalformed   Code `code:"0004" case:"Request malformed / Unacceptable" msg:"Request malformed / Unacceptable"`
	DatabaseError      Code `code:"0005" case:"Database Error" msg:"Database Error"`
	DatabaseConnError  Code `code:"0006" case:"Database Connection Error" msg:"Database Connection Error"`
	FormatError        Code `code:"0007" case:"Format Error" msg:"Format Error"`
	MismatchParamValue Code `code:"0008" case:"Mismatch Parameter Value" msg:"Mismatch Parameter Value"`

	MissingField Code `code:"0009" case:"Missing Field {field}" msg:"{field} is Required"`

	ErorTypeString  Code `code:"0010" case:"Key type is string, but the value isn't" msg:"Field {field} must be string only"`
	ErorTypeNumeric Code `code:"0011" case:"Key type is numeric, but the value isn't" msg:"Field {field} must be numeric only"`
	ErorTypeTime    Code `code:"0012" case:"Key Datetime is (ISO 8601), but the value isn't" msg:"Field {field} must be ISO-8601 only"`

	TokenInvalid        Code `code:"0013" case:"Token Invalid" msg:"Token Invalid"`
	TokenExpired        Code `code:"0014" case:"Token Expired" msg:"Token Expired"`
	TokenGenerateFailed Code `code:"0015" case:"Generate Token Failed" msg:"Generate Token Failed"`
	TokenGenerateDenied Code `code:"0016" case:"Permission Denied" msg:"Permission Denied"`
	// Common Format Error //

}

func init() {
	initEnums(&ErrorCode)
}

type (
	// TypeError ..
	TypeError struct {
		Code    string `json:"errorCode"`
		Case    string `json:"errorCase"`
		Message string `json:"errorDesc"`
	}
	// TypeGRPCError ...
	TypeGRPCError struct {
		Error []TypeError `json:"error"`
	}
	// TypeHTTPError ...
	TypeHTTPError struct {
		Error []TypeError `json:"error"`
	}
)
