package grpc_proxy_router

import (
	"context"
	"fmt"
	"log"
	"net"

	"fyqcode.top/go_gateway/dao"
	"fyqcode.top/go_gateway/grpc_proxy_middleware"
	"fyqcode.top/go_gateway/reverse_proxy"
	"github.com/e421083458/grpc-proxy/proxy"
	"google.golang.org/grpc"
)

var grpcServerList = []*warpGrpcServer{}

type warpGrpcServer struct {
	Addr string
	*grpc.Server
}

type tcpHandler struct {
}

func (t *tcpHandler) ServeTCP(ctx context.Context, src net.Conn) {
	src.Write([]byte("tcpHandler\n"))
}

func GrpcServerRun() {
	serviceList := dao.ServiceManagerHandler.GetGrpcServiceList()
	for _, serviceItem := range serviceList {
		tempItem := serviceItem
		go func(serviceDetail *dao.ServiceDetail) {
			addr := fmt.Sprintf(":%d", serviceDetail.GRPCRule.Port)
			rb, err := dao.LoadBalancerHandler.GetLoadBalancer(serviceDetail)
			if err != nil {
				log.Fatalf(" [INFO] GetGrpcLoadBalancer %v err:%v\n", addr, err)
				return
			}
			lis, err := net.Listen("tcp", addr)
			if err != nil {
				log.Fatalf("Grpc failed to listen: %v", err)
			}
			grpcHandler := reverse_proxy.NewGrpcLoadBalanceHandler(rb)
			s := grpc.NewServer(
				grpc.ChainStreamInterceptor(
					grpc_proxy_middleware.GrpcFlowCountMiddleware(serviceDetail),       // 流量统计
					grpc_proxy_middleware.GrpcFlowLimitMiddleware(serviceDetail),       // 限流
					grpc_proxy_middleware.GrpcJwtAuthTokenMiddleware(serviceDetail),    // jwt验证
					grpc_proxy_middleware.GrpcJwtFlowCountMiddleware(serviceDetail),    // 租户流量统计
					grpc_proxy_middleware.GrpcJwtFlowLimitMiddleware(serviceDetail),    //Jwt客户端限流
					grpc_proxy_middleware.GrpcWhiteListMiddleware(serviceDetail),       // ip白名单
					grpc_proxy_middleware.GrpcBlackListMiddleware(serviceDetail),       // ip黑名单
					grpc_proxy_middleware.GrpcHeaderTransferMiddleware(serviceDetail)), // heaer头转换
				grpc.CustomCodec(proxy.Codec()),
				grpc.UnknownServiceHandler(grpcHandler))

			grpcServerList = append(grpcServerList, &warpGrpcServer{
				Addr:   addr,
				Server: s,
			})
			log.Printf(" [INFO] grpc_proxy_run %v\n", addr)
			if err := s.Serve(lis); err != nil {
				log.Fatalf(" [INFO] grpc_proxy_run %v err:%v\n", addr, err)
			}

		}(tempItem)
	}
}

func GrpcServerStop() {
	for _, grpcServer := range grpcServerList {
		grpcServer.GracefulStop()
		log.Printf(" [INFO] grpc_proxy_stop %v stopped\n", grpcServer.Addr)
	}
}
