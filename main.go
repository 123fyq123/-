package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"

	"fyqcode.top/go_gateway/golang_common/lib"
	"fyqcode.top/go_gateway/http_proxy_router"
	"fyqcode.top/go_gateway/router"
)

// endpoint dashboard(后台管理)  server(代理服务器)
// config ./conf/prod/ (对应配置文件夹)

var (
	endpoint = flag.String("endpoint", "", "input endpoint like dashboard or server")
	config   = flag.String("config", "", "input config file like ./conf/dev/")
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

		go func() {
			http_proxy_router.HttpServerRun()
		}()

		go func() {
			http_proxy_router.HttpsServerRun()
		}()

		quit := make(chan os.Signal)
		signal.Notify(quit, syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
		<-quit

		http_proxy_router.HttpServerStop()
		http_proxy_router.HttpsServerStop()
	}
}
