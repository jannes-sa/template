package routers

import (
	"template/routers/grpc"
	"template/routers/http"
)

func init() {
	grpc.Router()
	http.Router()
}
