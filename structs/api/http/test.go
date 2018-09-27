package http

import (
	"template/structs"

	"github.com/jannes-sa/customvalidator"
)

type (
	// TestInterface - Test Interface Struct
	TestInterface struct {
		ID interface{} `json:"id" type:"int" validate:"type:E02030011"`
	}
	// ReqTest - Test Struct
	ReqTest struct {
		ID int
	}

	// ResTest - Test Struct
	ResTest struct {
		ID int `json:"id"`
	}
)

// ValidateRequest ...
func (st *TestInterface) ValidateRequest(
	assignStruct *ReqTest,
	errCode *[]structs.TypeError,
) {
	var paramsInterChange interface{} = *st
	codeError := customvalidator.Validate(
		paramsInterChange, assignStruct,
	)
	structs.GetCodeError(codeError, errCode)
}
