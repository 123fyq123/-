package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"

	"fyqcode.top/go_gateway/dao"
	"fyqcode.top/go_gateway/golang_common/lib"
	"fyqcode.top/go_gateway/grpc_proxy_router"
	"fyqcode.top/go_gateway/http_proxy_router"
	"fyqcode.top/go_gateway/router"
	"fyqcode.top/go_gateway/tcp_proxy_router"
)

// endpoint dashboard(后台管理)  server(代理服务器)
// config ./conf/prod/ (对应配置文件夹)

var (
	endpoint = flag.String("endpoint", "server", "input endpoint like dashboard or server")
	config   = flag.String("config", "./conf/dev/", "input config file like ./conf/dev/")
)

func main() {
	flag.Parse()
	if *endpoint == "" { // 用户没输入参数，直接退出
		flag.Usage()
		os.Exit(1)
	}
	if *config == "" { // 用户没输入参数，直接退出
		flag.Usage()
		os.Exit(1)
	}

	if *endpoint == "dashboard" {
		lib.InitModule(*config)
		defer lib.Destroy()
		router.HttpServerRun()

		quit := make(chan os.Signal)
		signal.Notify(quit, syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
		<-quit

		router.HttpServerStop()
	} else {
		lib.InitModule(*config)
		defer lib.Destroy()
		dao.ServiceManagerHandler.LoadOnce() // 加载服务列表
		dao.AppManagerHandler.LoadOnce()
		go func() {
			http_proxy_router.HttpServerRun()
		}()

		go func() {
			http_proxy_router.HttpsServerRun()
		}()

		go func() {
			tcp_proxy_router.TcpServerRun()
		}()

		go func() {
			grpc_proxy_router.GrpcServerRun()
		}()

		quit := make(chan os.Signal)
		signal.Notify(quit, syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
		<-quit

		tcp_proxy_router.TcpServerStop()
		grpc_proxy_router.GrpcServerStop()
		http_proxy_router.HttpServerStop()
		http_proxy_router.HttpsServerStop()
	}
}
