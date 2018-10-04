package grpc

type (
	// ReqTest - Test Struct
	ReqTest struct {
		ID   int    `json:"id"`
		Data string `json:"data"`
	}

	// ResTest - Test Struct
	ResTest struct {
		ID  int    `json:"id"`
		Res string `json:"response"`
	}
)
