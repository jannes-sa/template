/*
Version : 2.0
Author  : Jannes Santoso
Noted   : Use it only for validation request external data

Adding Array Validation
*/

package customvalidator

import (
	"encoding/json"
	"math/big"
	"reflect"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/astaxie/beego"

	"log"
)

// TypeStructAfterScan ...
type TypeStructAfterScan struct {
	Type        string
	NameField   string
	AssignField string
	Code        string
	Value       interface{}
	Validate    string
}

// Validate Custom Validating
func Validate(
	st interface{},
	overflowStruct interface{}) []string {

	var codeError []string

	var t reflect.Value
	ValidateNested(st, overflowStruct, t, &codeError)

	// start : handle slice struct //
	if len(codeError) == 0 {
		v := reflect.ValueOf(st)
		ve := reflect.ValueOf(overflowStruct).Elem()
		for i, n := 0, v.NumField(); i < n; i++ {
			if v.Field(i).Kind() == reflect.Slice {
				s := v.Field(i)

				nse := reflect.MakeSlice(ve.Field(i).Type(), s.Len(), s.Len())
				reflect.Copy(nse, ve.Field(i))
				ve.Field(i).Set(nse)

				for x := 0; x < s.Len(); x++ {
					vv := s.Index(x)
					vve := ve.Field(i).Index(x)

					var tt interface{}
					ValidateNested(vv.Interface(), tt, vve, &codeError)
				}

			}
		}
	}
	// end : handle slice struct //

	return codeError
}

// ValidateNested Custom Validating
func ValidateNested(
	st interface{},
	overflowStruct interface{},
	overflowStructValue reflect.Value,
	codeError *[]string) []string {

	v := reflect.ValueOf(st)
	vt := v.Type()

	var ve reflect.Value
	if reflect.ValueOf(overflowStruct).Kind() == reflect.Ptr {
		ve = reflect.ValueOf(overflowStruct).Elem()
	} else {
		ve = overflowStructValue
	}

	var scanDataRequest []TypeStructAfterScan

	for i, n := 0, v.NumField(); i < n; i++ {
		if v.Field(i).Kind() != reflect.Slice {
			f := v.Field(i)
			ft := vt.Field(i)

			var realVal interface{}
			var realType string
			stateType := true
			getTypeAndVal(f, ft, &stateType, &realVal, &realType)

			// validate missing field //
			validateFieldMissing(f, ft, codeError)
			// validate missing field //

			// validate type must match //
			if len(*codeError) == 0 {
				validateType(ft, stateType, codeError)
			}
			// validate type must match //

			// running validate //
			if ft.Tag.Get("validate") != "" && len(*codeError) == 0 {
				runningValidate(f, ft, stateType, realVal, realType, &scanDataRequest,
					codeError)
			}
			// running validate //

			// Parse Value //
			if ft.Tag.Get("convert") != "" && len(*codeError) == 0 {
				parseConvertValue(f, ft, realType, &realVal, codeError)
			}
			// Parse Value //

			// Validate After Scan //
			if ft.Tag.Get("validate") != "" && len(*codeError) == 0 {
				scanAfterValidate(f, ft, stateType, realVal, realType, &scanDataRequest,
					codeError)
				validateAfterScan(scanDataRequest, codeError)
			}
			// Validate After Scan //

			if len(*codeError) == 0 {
				// Set Value //
				if realType == "string" {
					ve.Field(i).SetString(realVal.(string))
				} else if realType == "int" || realType == "int64" {
					ve.Field(i).SetInt(int64(realVal.(float64)))
				} else if realType == "float64" {
					ve.Field(i).SetFloat(realVal.(float64))
				} else if realType == "time" {
					t, err := time.Parse(time.RFC3339, realVal.(string))
					if err != nil {
						log.Println("Failed Validate Time")
						log.Println(ft.Name)
						log.Println(realVal.(string))
					}
					ve.Field(i).Set(reflect.ValueOf(t))
				}
				// Set Value //

			} else {
				log.Println(ft.Name)
			}

		}
	}

	return *codeError
}

func parseConvertValue(f reflect.Value, ft reflect.StructField, realType string,
	realVal *interface{}, extractCodeError *[]string) {
	convert := ft.Tag.Get("convert")
	spr := strings.Split(convert, "=")
	valReal := *realVal

	if realType == "string" {
		if spr[0] == "removezero" {
			// convert:"removezero={code_error}"
			// to remove zero from string
			cnt, _ := new(big.Int), big.NewInt(1)
			cnt.SetString(valReal.(string), 10)
			*realVal = cnt.String()
		} else if spr[0] == "fixlength" {
			// fixlen:"fixlen={value}={code_error}"
			// using for thai fixedLength
			fixLen, err := strconv.Atoi(spr[1])
			CheckErr("Failed line 104 validate js", err)

			if fixLen > utf8.RuneCountInString(valReal.(string)) {
				negLen := fixLen - utf8.RuneCountInString(valReal.(string))
				zeros := strings.Repeat("0", negLen)
				*realVal = zeros + valReal.(string)
			} else {
				*extractCodeError = append(*extractCodeError, spr[len(spr)-1])
			}
		} else {
			*extractCodeError = append(*extractCodeError, spr[len(spr)-1])
		}
	} else {
		*extractCodeError = append(*extractCodeError, spr[len(spr)-1])
	}
}

func runningValidate(f reflect.Value, ft reflect.StructField, stateType bool, realVal interface{},
	realType string, scanDataRequest *[]TypeStructAfterScan, extractCodeError *[]string) {
	validateStr := ft.Tag.Get("validate")
	validateArr := strings.Split(validateStr, ",")

	for _, val := range validateArr {
		valArr := strings.Split(val, "=")
		if valArr[0] == "type" && !stateType {
			*extractCodeError = append(*extractCodeError, valArr[1])
		} else if stateType {
			if valArr[0] == "required" {
				if len(valArr) == 2 {
					requiredValidate(realType, realVal, extractCodeError, valArr[1])
				}
			} else if valArr[0] == "stringnumericonly" {
				if len(valArr) == 2 {
					stringnumericonlyValidate(realType, realVal, extractCodeError, valArr[1])
				}
			} else if valArr[0] == "gte" || valArr[0] == "lte" || valArr[0] == "len" {
				if len(valArr) == 3 {
					// len={value}={code_error}
					gteLteLenValidate(realType, realVal, extractCodeError, valArr)
				}
			} else if valArr[0] == "email" {
				if len(valArr) == 2 {
					emailValidate(realType, realVal, extractCodeError, valArr[1])
				}
			} else if valArr[0] == "should" {
				if len(valArr) == 4 {
					shouldValidate(realType, realVal, extractCodeError, valArr[1], valArr[2],
						valArr[3])
				}
			} else if valArr[0] == "commacheck" {
				if len(valArr) == 3 {
					commaCheck(realType, realVal, extractCodeError, valArr[1], valArr[2])
				}
			} else if valArr[0] == "enum" {
				if len(valArr) == 3 {
					// enum={value|value|value}={code_error}
					enumValidate(realType, realVal, valArr[1], valArr[2], extractCodeError)
				}
			} else if valArr[0] == "timevalid" {
				if len(valArr) == 2 {
					// timevalid={error_code}
					timeValidValidate(realType, realVal, extractCodeError, valArr[1])
				}
			}
		}
	}

}

// Create Validation Here //
func validateFieldMissing(f reflect.Value, ft reflect.StructField,
	extractCodeError *[]string) {
	if f.Interface() == nil {
		*extractCodeError = append(*extractCodeError, "missing_field|"+ft.Tag.Get("json"))
	}
}

func validateType(ft reflect.StructField, stateType bool, extractCodeError *[]string) {
	if !stateType {
		*extractCodeError = append(*extractCodeError, "mismatch_type|"+ft.Tag.Get("type")+"|"+ft.Tag.Get("json"))
	}
}

func enumValidate(realType string, realVal interface{},
	enumVal string, code string, extractCodeError *[]string) {

	arrEnumChar := strings.Split(enumVal, `|`)

	if realType == "string" {
		if !contains(arrEnumChar, realVal.(string)) {
			*extractCodeError = append(*extractCodeError, code)
		}
	} else if realType == "int" || realType == "int64" || realType == "float64" {
		valStr := strconv.FormatFloat(realVal.(float64), 'f', 0, 64)
		if !contains(arrEnumChar, valStr) {
			*extractCodeError = append(*extractCodeError, code)
		}
	}
}

func requiredValidate(realType string, realVal interface{}, extractCodeError *[]string,
	code string) {
	if realType == "string" {
		if realVal.(string) == "" {
			*extractCodeError = append(*extractCodeError, code)
		}
	} else if realType == "int" {
		if realVal.(float64) == 0 {
			*extractCodeError = append(*extractCodeError, code)
		}
	} else if realType == "float64" {
		if realVal.(float64) == float64(0) {
			*extractCodeError = append(*extractCodeError, code)
		}
	} else if realType == "time" {
		if realVal.(string) == "" {
			*extractCodeError = append(*extractCodeError, code)
		}
	}
}

func shouldValidate(realType string, realVal interface{}, extractCodeError *[]string,
	fixVal interface{}, commaDelimiter string, code string) {

	if realType == "string" {
		if realVal.(string) != fixVal.(string) {
			*extractCodeError = append(*extractCodeError, code)
		}
	} else if realType == "int" {
		str := strconv.FormatFloat(realVal.(float64), 'f', 0, 64)
		if str != fixVal.(string) {
			*extractCodeError = append(*extractCodeError, code)
		}
	} else if realType == "float64" {
		comInt, _ := strconv.Atoi(commaDelimiter)
		str := strconv.FormatFloat(realVal.(float64), 'f', comInt, 64)
		if str != fixVal.(string) {
			*extractCodeError = append(*extractCodeError, code)
		}
	}
}

func commaCheck(realType string, realVal interface{}, extractCodeError *[]string,
	commaDelimiter string, code string) {

	if realType == "float64" {
		commaDelimiterInt, _ := strconv.Atoi(commaDelimiter)
		nn, _ := json.Marshal(realVal.(float64))
		sepVal := strings.Split(string(nn), ".")

		if len(sepVal) == 2 {
			if len(sepVal[1]) > commaDelimiterInt {
				*extractCodeError = append(*extractCodeError, code)
			}
		}
	} else {
		*extractCodeError = append(*extractCodeError, code)
	}
}

func stringnumericonlyValidate(realType string, realVal interface{}, extractCodeError *[]string,
	code string) {
	if realType == "string" && realVal.(string) != "" {
		// _, errConv := strconv.ParseUint(realVal.(string), 10, 64)
		_, errConv := strconv.ParseFloat(realVal.(string), 64)
		if errConv != nil {
			log.Println(errConv)
			*extractCodeError = append(*extractCodeError, code)
		}
	}
}

func gteLteLenValidate(realType string, realVal interface{}, extractCodeError *[]string,
	valArr []string) {
	intNil, errAtoi := strconv.ParseFloat(valArr[1], 64)
	CheckErr("Failed Convert custom validate line 69", errAtoi)

	if realType == "string" {
		stCheck := false
		if valArr[0] == "gte" {
			if float64(utf8.RuneCountInString(realVal.(string))) >= intNil {
				stCheckAsgn(&stCheck)
			}
		} else if valArr[0] == "lte" {
			if float64(utf8.RuneCountInString(realVal.(string))) <= intNil {
				stCheckAsgn(&stCheck)
			}
		} else if valArr[0] == "len" {
			if float64(utf8.RuneCountInString(realVal.(string))) == intNil {
				stCheckAsgn(&stCheck)
			}
		}
		if !stCheck {
			*extractCodeError = append(*extractCodeError, valArr[2])
		}
	} else if realType == "int" {
		stCheck := false
		if valArr[0] == "gte" {
			if realVal.(float64) >= intNil {
				stCheckAsgn(&stCheck)
			}
		} else if valArr[0] == "lte" {
			if realVal.(float64) <= intNil {
				stCheckAsgn(&stCheck)
			}
		} else if valArr[0] == "len" {
			if realVal.(float64) == intNil {
				stCheckAsgn(&stCheck)
			}
		}
		if !stCheck {
			*extractCodeError = append(*extractCodeError, valArr[2])
		}
	} else if realType == "float64" {
		stCheck := false
		if valArr[0] == "gte" {
			if realVal.(float64) >= intNil {
				stCheckAsgn(&stCheck)
			}
		} else if valArr[0] == "lte" {
			if realVal.(float64) <= intNil {
				stCheckAsgn(&stCheck)
			}
		} else if valArr[0] == "len" {
			if realVal.(float64) == intNil {
				stCheckAsgn(&stCheck)
			}
		}
		if !stCheck {
			*extractCodeError = append(*extractCodeError, valArr[2])
		}
	}
}
func stCheckAsgn(check *bool) {
	*check = true
}

func emailValidate(realType string, realVal interface{}, extractCodeError *[]string,
	code string) {
	if realType == "string" && realVal.(string) != "" {
		errMail := ValidateFormatMail(realVal.(string))
		if errMail != nil {
			*extractCodeError = append(*extractCodeError, code)
		}
	}
}

func timeValidValidate(realType string, realVal interface{}, extractCodeError *[]string,
	code string) {

	if realType == "string" && realVal.(string) != "" {
		tm, err := time.Parse(time.RFC3339, realVal.(string)+"T00:00:00.000Z")
		log.Println(tm)
		if err != nil {
			log.Println(err)
			*extractCodeError = append(*extractCodeError, code)
		}
	}
}

///////////////////////

// Validate After Scan //
func scanAfterValidate(f reflect.Value, ft reflect.StructField, stateType bool, realVal interface{},
	realType string, scanDataRequest *[]TypeStructAfterScan,
	extractCodeError *[]string) {
	validateStr := ft.Tag.Get("validate")
	validateArr := strings.Split(validateStr, ",")

	for _, val := range validateArr {
		valArr := strings.Split(val, "=")
		if valArr[0] == "identicField" {
			*scanDataRequest = append(*scanDataRequest, TypeStructAfterScan{
				Type:        realType,
				NameField:   ft.Name,
				AssignField: valArr[1],
				Code:        valArr[2],
				Value:       realVal,
				Validate:    valArr[0],
			})
		}
	}
}

func validateAfterScan(scanDataRequest []TypeStructAfterScan, extractCodeError *[]string) {
	for _, val := range scanDataRequest {
		if val.Validate == "identicField" {
			st := validateIdentical(scanDataRequest, val)
			if st == false {
				*extractCodeError = append(*extractCodeError, val.Code)
			}
		}
	}
}

func validateIdentical(scanDataRequest []TypeStructAfterScan, valAfterScan TypeStructAfterScan) bool {
	state := true
	for _, val := range scanDataRequest {
		if val.NameField == valAfterScan.AssignField && val.Type == valAfterScan.Type {
			if val.Type == "string" {
				if val.Value.(string) == valAfterScan.Value.(string) {
					state = false
				}
			} else if val.Type == "int" {
				if val.Value.(float64) == valAfterScan.Value.(float64) {
					state = false
				}
			} else if val.Type == "float64" {
				if val.Value.(float64) == valAfterScan.Value.(float64) {
					state = false
				}
			}
		}
	}
	return state
}

////////////////////////

// Parse Value //
func removeZero() {

}

/////////////////

func getTypeAndVal(f reflect.Value, ft reflect.StructField, stateType *bool, realVal *interface{},
	realType *string) {
	strType := ft.Tag.Get("type")
	arrType := strings.Split(strType, ",")

	checkType(f.Interface(), stateType, arrType, realVal, realType, strType)
	if !(*stateType) && ft.Tag.Get("validate") != "" {
		log.Println("FAILED VALIDATE !!!!!")
		log.Println(ft.Name)
		log.Println(f.Interface())
		log.Println("FAILED VALIDATE !!!!!")
	}
}

func checkType(mpt interface{}, state *bool, status []string, val *interface{},
	typeVal *string, strType string) {
	*state = false
	switch v := mpt.(type) {
	case int:
		if contains(status, "int") {
			*state = true
		}
		*val = v
		*typeVal = "int"
	case float64:
		if strType == "float64" {
			if contains(status, "float64") {
				*state = true
			}
			*val = v
			*typeVal = "float64"
		} else if strType == "int" {
			if contains(status, "int") {
				*state = true
			}
			*val = v
			*typeVal = "int"
		}
	case string:
		if strType == "string" {
			if contains(status, "string") {
				*state = true
			}
			*typeVal = "string"
		} else if strType == "time" {
			if contains(status, "time") {
				*state = true
			}
			*typeVal = "time"
		}
		*val = v
	case *big.Int:
		beego.Debug(v)
		*val = v
	default:
		*state = false
	}
}

func contains(slice []string, item string) bool {
	set := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		set[s] = struct{}{}
	}

	_, ok := set[item]
	return ok
}

// CheckErr ...
func CheckErr(msg string, err error) {
	if err != nil {
		log.Println(msg)
		panic(err)
	}
}
