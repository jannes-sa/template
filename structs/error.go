package structs

import (
	"log"
	"reflect"
	"strings"
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
		Code:    str["code"],
		Case:    str["case"],
		Message: strMsg,
	})
}

// GetString ...
func (e Code) GetString() (string, string, string) {
	str := getStr(e)
	return str["code"], str["case"], str["msg"]
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
				Code:    vt.Field(i).Tag.Get("code"),
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
		msgs = strings.Replace(msgs, `{field}`, arrMisField[1], -1)
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
				msgFilter := strings.Replace(msgsStr, `{field}`, arrType[2], -1)
				*errCode = append(*errCode, TypeError{
					Code:    codeStr,
					Case:    casesStr,
					Message: msgFilter,
				})
			} else if arrType[1] == "int" || arrType[1] == "int64" || arrType[1] == "float64" {
				msgFilter := strings.Replace(msgsNum, `{field}`, arrType[2], -1)
				*errCode = append(*errCode, TypeError{
					Code:    codeNum,
					Case:    casesNum,
					Message: msgFilter,
				})
			} else if arrType[1] == "time" {
				msgFilter := strings.Replace(msgsTim, `{field}`, arrType[2], -1)
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
				Code:    codeTag,
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
	// Wallet Transaction Format Error //
	UnexpectedError            Code `code:"E02030000" case:"Unexpected error" msg:"Unexpected error"`
	InvalidAccountNumber       Code `code:"E02030001" case:"Beneficiary account number not found" msg:"Invalid To Account Number"`
	InsufficientFund           Code `code:"E02030002" case:"Source available balance - transaction amount < 0" msg:"Insufficient Fund"`
	RestrictionAmountFound1    Code `code:"E02030003" case:"Beneficiary account balance > Maximum Balance (after transaction)" msg:"Exceed maximum balance"`
	RestrictionAmountFound2    Code `code:"E02030004" case:"Beneficiary account's total_credit_month_to_date > maximum_credit_per_month (after transaction)" msg:"Exceed maximum credit transaction amount per month"`
	InvalidAccountNumberSource Code `code:"E02030005" case:"From account number not found" msg:"Invalid From Account Number"`
	ProfileFailed              Code `code:"E02030006" case:"TXN Domain return an error when our domain send a post transaction request" msg:"Your transaction failed, please try again later"`
	// Wallet Transaction Format Error //

	// Spesific Request Format Error //
	BeneficiaryExceedChar    Code `code:"E02030007" case:"Beneficiary account number exceed 20 characters" msg:"To Account No. must not be more than 20 characters"`
	BeneficiaryAccNotNumeric Code `code:"E02030008" case:"Beneficiary account number is not numeric" msg:"To Account No. must be numeric only"`
	BeneficiaryEmpty         Code `code:"E02030009" case:"Missing beneficiary account number" msg:"To Account No. is required"`

	TransactionAmountLength      Code `code:"E02030010" case:"Transaction amount length > (18,2)" msg:"Amount must not be more than 999.999.999.999.999,99"`
	TransactionAmountOnlyNumeric Code `code:"E02030011" case:"Transaction amount contain value other than numeric" msg:"Amount must be numeric only"`
	TransactionAmountEmpty       Code `code:"E02030012" case:"Missing transaction amount" msg:"Transaction Amount is Required"`

	SourceAccountNumberNotNumeric Code `code:"E02030013" case:"From Accont Number is not numeric" msg:"From Account No. must be numeric only"`
	SourceAccountNumberExceed     Code `code:"E02030014" case:"From account number exceed 20 characters" msg:"From Account No. must not be more than 20 characters"`
	SourceAccountEmpty            Code `code:"E02030015" case:"Missing from account number" msg:"From Account No. is required"`

	SourceBenefeciarySameNumber Code `code:"E02030016" case:"From and beneficiary account number are same" msg:"From and Beneficiary Account are same"`
	ServiceBranchLength         Code `code:"E02030059" case:"Service Branch length > 4" msg:"Service Branch must not be more than 4 digits"`
	ServiceBranchRequired       Code `code:"E02030060" case:"Service Branch must not be 0" msg:"Service Branch must not be 0"`
	// Spesific Request Format Error //
	InvalidToBankCodeLength Code `code:"E02030017" case:"Invalid To Bank ID (length)" msg:"To Bank Code must not be more than 3 characters"`
	InvalidToBankCodeType   Code `code:"E02030018" case:"Invalid To Bank ID (type)" msg:"To Bank Code must be numeric only"`
	MissingToBankCode       Code `code:"E02030019" case:"Missing To Bank ID" msg:"To Bank Code is required"`

	InvalidSourceBankCodeLength Code `code:"E02030020" case:"Invalid Source Bank ID (length)" msg:"From Bank Code must not be more than 3 characters"`
	InvalidSourceBankCodeType   Code `code:"E02030021" case:"Invalid Source Bank ID (type)" msg:"From Bank Code must be numeric only"`
	MissingSourceBankCode       Code `code:"E02030022" case:"Missing Source Bank ID" msg:"From Bank Code is required"`

	MissingAccountNumber     Code `code:"E02030026" case:"Missing Account Number" msg:"Account Number is required"`
	AccountNumberOnlyNumeric Code `code:"E02030027" case:"Account Number is not numeric" msg:"Account Number must be numeric"`
	AccountNumberExceed      Code `code:"E02030028" case:"CccountNumber exceed 20 characters" msg:"Account Number must not be more than 20 characters"`
	AccountNumberNotFound    Code `code:"E02030029" case:"Account Number not found" msg:"Invalid Account Number"`
	DateTimeExceed           Code `code:"E02030030" case:"From Date is more than 6 months" msg:"From Date is more than 6 months"`

	MissingEarningPoint     Code `code:"E02030033" case:"Missing earning points" msg:"Missing earning points"`
	InvalidCurrency         Code `code:"E02030034" case:"Invalid currency (RWP)" msg:"Currency must be in RWP"`
	InsufficientPoints      Code `code:"E02030035" case:"Source available points - transaction amount < 0" msg:"Insufficient Points"`
	MissingRedemptionPoints Code `code:"E02030036" case:"Missing redemption points" msg:"Redeeming Points is required"`
	AccountIsNotCorrect     Code `code:"E02030038" case:"Account is not in the correct product code" msg:"This account is not in {product_code} product"`
	TransactionAmountZero   Code `code:"E02030039" case:"Transaction amount must not be zero" msg:"Transaction amount must not be zero"`

	ValueRedemptionLength      Code `code:"E02030040" case:"Value of redemption length > (18,2)" msg:"Value of redemption must not be more than 999.999.999.999.999,99"`
	ValueRedemptionZero        Code `code:"E02030041" case:"Value of redemption must not be zero" msg:"Value of redemption must not be zero"`
	ValueRedemptionOnlyNumeric Code `code:"E02030042" case:"Value redemption contain value other than numeric" msg:"Value of redemption must be numeric only"`
	AccountIsNotLoyalty        Code `code:"E02030043" case:"Account is not in the loyalty product code" msg:"This account is not in 'Loyalty' product"`
	FromDateInvalid            Code `code:"E02030052" case:"From Date > To Date" msg:"From Date must Not be Greater than To Date"`
	MaturityTrfAccNotExist     Code `code:"E02030053" case:"Maturity Transfer Account does not exist" msg:"Invalid Maturity Transfer Account"`
	MaturityTrfAccNotFound     Code `code:"E02030054" case:"Maturity Transfer Account is blank" msg:"Maturity Transfer Account is Required"`
	MaturityAccMustBlank       Code `code:"E02030055" case:"Maturity Transfer Account Must be Blank" msg:"Maturity Transfer Account Must be Blank"`
	// Batch Earning Format Error //
	BatchEarningTotalNotValid           Code `code:"E02030029" case:"Total record in trailer <> total number of record" msg:"Total record in the trailer is not equal to total number of records"`
	BatchEarningDateNotValid            Code `code:"E02030030" case:"Data in file <> System Date" msg:"Date in batch file is not equal in actual date"`
	BatchAlreadyExecuted                Code `code:"E02030031" case:"Batch file can only be processed once per day" msg:"Cannot process this file. Earning Point Batch file for date [DD/MM/YYYY] Had been processed"`
	BatchEarningFilenameRequired        Code `code:"E02030032" case:"Filename is Missing" msg:"Filename is Required"`
	BatchEarningFileNotExist            Code `code:"E02030033" case:"Missing File Batch Earning Points" msg:"File Not Exists"`
	BatchEarningCreateFileResult        Code `code:"E02030034" case:"Error Creating File Result" msg:"Failed to Create File Result"`
	BatchEarningFormatFileNotValid      Code `code:"E02030035" case:"Format File Not Valid" msg:"Format File Not Valid"`
	BatchEarningFileGenerateWithWarning Code `code:"E02030036" case:"File Generated With Warning" msg:"File Generated With Warning"`
	BatchEarningStatusCreateFailed      Code `code:"E02030037" case:"Error Create/Update Batch Earning Status" msg:"Failed to Insert Batch Earning Status"`
	BatchEarningFilenameNotValid        Code `code:"E02030074" case:"Filename Not Valid" msg:"Filename is not valid"`
	BatchEarningRefJobIDNotEmpty        Code `code:"E02030075" case:"RefJobID Not Empty For Earning" msg:"ref_job_id is not empty for earning"`
	BatchAdjustEarningRefJobIDRequired  Code `code:"E02030076" case:"RefJobID Is Required For Adjust Earning" msg:"ref_job_id is required for adjusted earning"`

	// Batch Earning Format Error //

	// InterestPosting Format Error //
	IntPostingInvalidAccount Code `code:"E02009996" case:"Account Number is not exit" msg:"Account Number is not exit"`

	ChannelPostDateFormat       Code `code:"E02030069" case:"Channel Post Date format is not (RFC 3339)" msg:"Channel Post Date format must (RFC 3339)"`
	BatchDateLoyaltyReportDaily Code `code:"E02030070" case:"Date format for generate loyalty report daily is not (YYYY-MM-DD)" msg:"Date format for generate loyalty report daily must (YYYY-MM-DD)"`
	MissingChannelPostDate      Code `code:"E02030073" case:"Missing Channel Post Date" msg:"Channel Post Date is required"`

	// Common Format Error //
	RPCInvalid               Code `code:"E02009999" case:"Connection RPC Invalid" msg:"Connection RPC Invalid"`
	ServiceNotAllowedToBatch Code `code:"E02009997" case:"Service Not Allowed to Batch" msg:"Service Not Allowed to Batch"`
	URLFilterInvalid         Code `code:"E02009998" case:"URL Filter Invalid" msg:"URL Filter Invalid"`

	RequestMalformed             Code `code:"E02000003" case:"Request malformed / Unacceptable" msg:"Request malformed / Unacceptable"`
	DatabaseError                Code `code:"E02000100" case:"Database Error" msg:"Database Error"`
	DatabaseConnError            Code `code:"E02000101" case:"Database Connection Error" msg:"Database Connection Error"`
	FormatError                  Code `code:"E02000401" case:"Format Error" msg:"Format Error"`
	MismatchParamValue           Code `code:"E02000402" case:"Mismatch Parameter Value" msg:"Mismatch Parameter Value"`
	TransactionFailedFromProfile Code `code:"E02030006" case:"TXN Domain return an error when our domain send a post transaction request" msg:"Your transaction failed, please try again later"`

	TargetEndPointNotFound Code `code:"E02000004" case:"Target Endpoint not found" msg:"Target Endpoint Not Found"`

	MissingField Code `code:"E02000005" case:"Missing Field {field}" msg:"{field} is Required"`

	ErorTypeString  Code `code:"E02000006" case:"Key type is string, but the value isn't" msg:"Field {field} must be string only"`
	ErorTypeNumeric Code `code:"E02000007" case:"Key type is numeric, but the value isn't" msg:"Field {field} must be numeric only"`
	ErorTypeTime    Code `code:"E02000008" case:"Key Datetime is (ISO 8601), but the value isn't" msg:"Field {field} must be ISO-8601 only"`

	ErrorDataOutstandingNotFound   Code `code:"E02099992" case:"Error Data Outstanding Not Found" msg:"Error Data Outstanding Not Found"`
	ErrorDataInterclearingNotFound Code `code:"E02099993" case:"Error Data Interclearing Not Found" msg:"Error Data Interclearing Not Found"`
	ErrorDataAccrueNotFound        Code `code:"E02099994" case:"Error Data Accrue Not Found" msg:"Error Data Accrue Not Found"`
	CreateExtractAccrueData        Code `code:"E02009995" case:"Error Create Extract file for Accrual Process" msg:"Error Create Extract file for Accrual Process"`

	ErrorHasNoAccount             Code `code:"E020299993" case:"This Account {acc_num} has no Account" msg:"This Account {acc_num} has no Account"`
	CreateDateGreatherThanTxnTime Code `code:"E02030037" case:"Account creation time > transaction time" msg:"Account {acc_num} creation time > transaction time"`
	// Common Format Error //

	// Error From Shell Script //
	FailedCallURLSh    Code `code:"E02088881" case:"Failed Call URL" msg:"Failed Call URL"`
	FailedReadRespSh   Code `code:"E02088882" case:"Failed Read Response" msg:"Failed Read Response"`
	FailedStatusCodeSh Code `code:"E02088883" case:"Status Code not 200" msg:"Status Code not 200"`
	// Error From Shell Script //

	JobIDNotExists        Code `code:"E02030140" case:"pre-job-id not exists" msg:"pre-job-id not exists"`
	ReconFileLoyaltyEmpty Code `code:"E02030071" case:"Data Not Available" msg:"Data Not Available"`

	TransactionAccountMinus Code `code:"E0209898984" case:"Transaction amount should not minus" msg:"Transaction amount should not minus"`
	ValueRedemption         Code `code:"E0209898985" case:"Value redemption should not minus" msg:"Value redemption should not minus"`

	InterDomainError Code `code:"E0209898986" case:"Interdomain Error" msg:"There is error from destination domain"`

	FailedAccountBranchRecon Code `code:"E02030068" case:"Interdomain Error" msg:"Failed to generate {report_name} for Date {date}. It contains invalid data, please check {filename}"`

	//Loyalty Report
	OriginalReverseJobIdNotFound Code `code:"E02030072" case:"Original Reverse JobId NotFound" msg:"Original Reverse JobId Not Found"`

	AccountClosed Code `code:"E02030077" case:"Account Already Closed" msg:"This account number was closed"`

	//Reverese Redemption
	TransactionReversed               Code `code:"E02030061" case:"Transaction has been reversed" msg:"Transaction has been reversed"`
	AmountNotMatchWithOriginal        Code `code:"E02030062" case:"Amount Not Match" msg:"Amount Not Match with the Original Transaction"`
	AccountNumberNotMatchWithOriginal Code `code:"E02030064" case:"Account Number Not Match" msg:"Account Number Not Match with the Original Transaction"`
	TransactionIsNotRedemption        Code `code:"E02030065" case:"Transaction is not redemption" msg:"Transaction is not redemption"`
	OriginalJobIDNotFound             Code `code:"E02030063" case:"Original Job ID not found" msg:"Invalid Original Job ID"`

	// Batch Clearing Point Format Error
	BatchClearingPointCreateFileResult   Code `code:"E02030064" case:"Error Creating File Result" msg:"Failed to Create File Result"`
	BatchClearingPointFormatFileNotValid Code `code:"E02030035" case:"Format File Not Valid" msg:"Format File Not Valid"`
	BatchClearingPointTotalNotValid      Code `code:"E02030029" case:"Total record in trailer <> total number of record" msg:"Total record in the trailer is not equal to total number of records"`
	BatchClearingPointDateNotValid       Code `code:"E02030030" case:"Data in file <> System Date" msg:"Date in batch file is not equal in actual date"`
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
