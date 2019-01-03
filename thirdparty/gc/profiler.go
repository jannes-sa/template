package gc

// import (
// 	"log"
// 	"os"
// 	"path/filepath"
// 	"template/helper/constant"

// 	"cloud.google.com/go/profiler"
// )

// // SetProfiler - SetProfiler Console
// func SetProfiler() {
// 	file := filepath.Join(
// 		os.Getenv("GOPATH"),
// 		"src",
// 		constant.GOAPP,
// 		"thirdparty", "gc", "APIProject-key.json",
// 	)
// 	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", file)
// 	go func() {
// 		if err := profiler.Start(profiler.Config{
// 			Service:        constant.GOAPP + "-" + constant.GOENV,
// 			ServiceVersion: "1.0",
// 			ProjectID:      "jannessantoso12", // optional on GCP
// 		}); err != nil {
// 			log.Fatalf("Cannot start the profiler: %v", err)
// 		}
// 	}()
// }
