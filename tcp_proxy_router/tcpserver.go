package tcp_proxy_router

import (
	"context"
	"fmt"
	"log"
	"net"

	"fyqcode.top/go_gateway/dao"
	"fyqcode.top/go_gateway/tcp_server"
)

var tcpServerList = []*tcp_server.TcpServer{}

type tcpHandler struct {
}

func (t *tcpHandler) ServeTCP(ctx context.Context, src net.Conn) {
	src.Write([]byte("tcpHandler\n"))
}

func TcpServerRun() {
	serviceList := dao.ServiceManagerHandler.GetTcpServiceList()
	for _, serviceItem := range serviceList {
		tmpItem := serviceItem
		log.Printf(" [INFO] tcp_proxy_run:%v\n", tmpItem.TCPRule.Port)

		go func(serviceDetail *dao.ServiceDetail) {
			addr := fmt.Sprintf(":%d", serviceDetail.TCPRule.Port)
			tcpServer := &tcp_server.TcpServer{
				Addr:    addr,
				Handler: &tcpHandler{},
			}
			tcpServerList = append(tcpServerList, tcpServer) // 用于后面关闭
			if err := tcpServer.ListenAndServe(); err != nil && err != tcp_server.ErrServerClosed {
				log.Fatalf(" [INFO] tcp_proxy_run :%v err:%v\n", tmpItem.TCPRule.Port, err)
			}
		}(tmpItem)
	}
}

func TcpServerStop() {
	for _, tcpServer := range tcpServerList {
		tcpServer.Close()
		log.Printf(" [INFO] tcp_proxy_run stopped %v\n", tcpServer.Addr)
	}
}
