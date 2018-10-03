package componenttest

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"strconv"
	"template/models/db"
	ctrlRPC "template/routers/grpc"
	routerHTTP "template/routers/http"
	"template/structs"
	"time"

	_ "github.com/lib/pq"

	"github.com/astaxie/beego"
)

// DBinit ...
func DBinit() {
	db.RegisterPGSQL()
}

// HTTPInit ...
func HTTPInit() {
	routerHTTP.Router()
	appath, _ := filepath.Abs(filepath.Dir("../../../http"))
	beego.TestBeegoInit(appath)
}

// SendHTTP ...
func SendHTTP(
	method string,
	URL string,
	bodyByte []byte,
) (response structs.RespData) {
	newRequest, errNewRequest := http.NewRequest(method, URL, bytes.NewBuffer(bodyByte))
	checkError(errNewRequest)

	// 5. set http header
	newRequest.Header.Set("X-Request-Id", "req-id")
	newRequest.Header.Set("X-Job-Id", "")
	newRequest.Header.Set("X-Real-Ip", "ip")
	newRequest.Header.Set("X-Caller-Service", "serv")
	newRequest.Header.Set("X-Caller-Domain", "dom")
	newRequest.Header.Set("X-Device", "device")
	newRequest.Header.Set("User-Agent", "agent")
	newRequest.Header.Set("Datetime", "2006-01-02T15:04:05Z")
	newRequest.Header.Set("Accept", "accept")

	// 6. test send request
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, newRequest)

	// 7. get response after send request
	errUnmarshal := json.Unmarshal(w.Body.Bytes(), &response)
	checkError(errUnmarshal)

	return
}

// GRPCInit ...
func GRPCInit() {
	ctrlRPC.CreateGrpcServer(strconv.Itoa(beego.BConfig.Listen.HTTPPort), "test")
	time.Sleep(2 * time.Second)
}

func checkError(err error) {
	if err != nil {
		beego.Warning(err)
	}
}
