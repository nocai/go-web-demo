package main

import (
	"flag"

	"github.com/golang/glog"
	"github.com/nocai/go-web-demo/api"
	"github.com/nocai/go-web-demo/infra"
	"github.com/nocai/go-web-demo/modules"
)

var (
	// Version version
	Version = "-"
	// GitHash git hash
	GitHash = "-"
	// BuildTime build time
	BuildTime = "-"
)

func init() {
	// glog 需要读取命令行参数
	flag.Parse()

	glog.Infof("Version: %v", Version)
	glog.Infof("GitHash: %v", GitHash)
	glog.Infof("BuildTime: %v", BuildTime)
}

// 日志用glog
//
// -log_dir="logs"
// -alsologtostderr -stderrthreshold=INFO
func main() {
	//退出时，确保日志写入文件中
	defer glog.Flush()

	app := infra.NewApp()

	var (
		yourServiceServer = modules.NewYourServiceServer()
	)
	api.RegisterYourServiceServer(app.GRPCServer, yourServiceServer)
	if err := api.RegisterYourServiceHandlerFromEndpoint(app.Ctx, app.ServeMux, app.GRPCAddr, app.DialOptions); err != nil {
		glog.Error(err)
	}

	if err := app.Run(); err != nil {
		glog.Error(err)
	}
}
