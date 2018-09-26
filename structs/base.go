package structs

import (
	"time"

	"github.com/astaxie/beego"
	validator "gopkg.in/go-playground/validator.v9"
)

type (
	//ReqHTTPHeader Base All Requesr HTTP Header
	ReqHTTPHeader struct {
		XRequestID     string    `json:"X-Request-Id" validate:"required"`
		XJobID         string    `json:"X-Job-Id"`
		XRealIP        string    `json:"X-Real-Ip" validate:"required"`
		XCallerService string    `json:"X-Caller-Service" validate:"required"`
		XCallerDomain  string    `json:"X-Caller-Domain" validate:"required"`
		XDevice        string    `json:"X-Device"`
		XApplication   string    `json:"X-Application"`
		XChannel       string    `json:"X-Channel"`
		UserAgent      string    `json:"User-Agent"`
		Datetime       time.Time `json:"Datetime" validate:"required"`
		Accept         string    `json:"Accept" validate:"required"`
		AcceptLanguage string    `json:"Accept-Language"`
		AcceptEncoding string    `json:"Accept-Encoding"`
		DomainID       string    `json:"domain-id"`
	}

	//ResHTTPHeader Base All Response HTTP Header
	ResHTTPHeader struct {
		XRequestID  string    `json:"x-request-id"`
		XJobID      string    `json:"x-job-id"`
		Datetime    time.Time `json:"datetime"`
		ContentType string    `json:"content-type"`
		XRoundTrip  string    `json:"x-roundtrip"`
	}

	// ReqData Base All Request
	ReqData struct {
		ReqBody interface{} `json:"rqBody"`
	}

	// RespData Base All Response
	RespData struct {
		// ReqHeader JSONHeaderResp    `json:"rsHeader"`
		ResponseBody interface{}       `json:"rsBody"`
		Error        []FilterErrorCode `json:"error"`
	}

	// FilterErrorCode ...
	FilterErrorCode struct {
		Code    string `json:"errorCode"`
		Message string `json:"errorDesc"`
	}

	// JSONHeader JSON header request for async/sync message
	JSONHeader struct {
		JobID         string    `json:"job-id"`                                        // job unique id, Generate by API gateway
		MessageID     string    `json:"message-id" validate:"required"`                // Message ID, Generated with RFC4122
		ReqSvc        string    `json:"requested-service" validate:"required"`         // requested service
		ReqSvcVer     string    `json:"requested-service-version" validate:"required"` // request service version
		CallSvc       string    `json:"caller-service" validate:"required"`            // Caller Service
		CallDmn       string    `json:"caller-domain"`                                 // Caller Domain
		CallRtnSvc    string    `json:"caller-return-service"`                         // return service (send response back to this service)
		CallRtnSvcVer string    `json:"caller-return-service-version"`                 // return service version (send response back to this version)
		DateTime      time.Time `json:"datetime" validate:"required"`                  // Date and time of request
		AcceptLang    string    `json:"accept-language" validate:"required"`           // language which is acceptable for the response (e.g en/th)
		AcceptEnc     string    `json:"accept-encoding" validate:"required"`           // Character encoding which is acceptable for the response (UTF-8)
	}

	// JSONHeaderResp JSON Header Response for async/sync message
	JSONHeaderResp struct {
		JobID         string    `json:"job-id"`                        // job unique id, Generate by API gateway
		MessageID     string    `json:"message-id"`                    // Message ID, Generated with RFC4122
		CallRtnSvc    string    `json:"caller-return-service"`         // return service (send response back to this service)
		CallRtnSvcVer string    `json:"caller-return-service-version"` // return service version (send response back to this version)
		ReqDmn        string    `json:"requested-domain"`              //
		DateTime      time.Time `json:"datetime"`                      // Date and time of request
		AcceptLang    string    `json:"accept-language"`               // language which is acceptable for the response (e.g en/th)
		AcceptEnc     string    `json:"accept-encoding"`
	}

	// LogType logging struct for standard centralize
	LogType struct {
		JobID         interface{} `json:"job_id"`
		URL           interface{} `json:"url"`
		HeaderReq     interface{} `json:"header_req"`
		BodyReq       interface{} `json:"body_req"`
		HeaderRes     interface{} `json:"header_res"`
		BodyRes       interface{} `json:"body_res"`
		Type          string      `json:"type"`
		ToRPCApp      string      `json:"to_rpc"`
		RealRoundTrip string      `json:"real_rountrip"`
	}
)

// ValidateReqHeader ...
func (st *JSONHeader) ValidateReqHeader() error {
	validate := validator.New()
	err := validate.Struct(st)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			beego.Warning(err)
		}

		for _, err := range err.(validator.ValidationErrors) {
			beego.Warning(err.Namespace())
			beego.Warning(err.Field())
			beego.Warning(err.StructNamespace())
			beego.Warning(err.StructField())
			beego.Warning(err.Tag())
			beego.Warning(err.ActualTag())
			beego.Warning(err.Kind())
			beego.Warning(err.Type())
			beego.Warning(err.Value())
			beego.Warning(err.Param())
		}

	}
	return err
}

// ValidateReqHTTPHeader ...
func (st *ReqHTTPHeader) ValidateReqHTTPHeader() error {
	validate := validator.New()
	err := validate.Struct(st)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			beego.Warning(err)
		}

		for _, err := range err.(validator.ValidationErrors) {
			beego.Warning(err.Namespace())
			beego.Warning(err.Field())
			beego.Warning(err.StructNamespace())
			beego.Warning(err.StructField())
			beego.Warning(err.Tag())
			beego.Warning(err.ActualTag())
			beego.Warning(err.Kind())
			beego.Warning(err.Type())
			beego.Warning(err.Value())
			beego.Warning(err.Param())
		}

	}
	return err
}

// CheckErr ...
func CheckErr(msg string, err error) {
	if err != nil {
		beego.Warning(msg, err)
	}
}
