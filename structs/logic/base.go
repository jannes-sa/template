package logic

import (
	"template/structs"
)

type (
	// ContextStruct - Struct for contain all data related to context and put it in argument
	ContextStruct struct {
		HeaderAll string
		Header    structs.ReqHTTPHeader
		JobID     string
	}
)
