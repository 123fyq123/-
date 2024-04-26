package grpc_proxy_middleware

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strings"

	"fyqcode.top/go_gateway/dao"
	"fyqcode.top/go_gateway/public"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
)

// JWT限流
func GrpcJwtFlowLimitMiddleware(serviceDetail *dao.ServiceDetail) func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		md, ok := metadata.FromIncomingContext(ss.Context())
		if !ok {
			return errors.New("miss metadata from context")
		}
		appInfos := md.Get("app")
		if len(appInfos) == 0 {
			if err := handler(srv, ss); err != nil { // 执行下游方法
				log.Printf("RPC failed with error %v\n", err)
				return err
			}
			return nil
		}
		appInfo := &dao.App{}
		if err := json.Unmarshal([]byte(appInfos[0]), appInfo); err != nil {
			return err
		}
		peerCtx, ok := peer.FromContext(ss.Context())
		if !ok {
			return errors.New("peer not found with context")
		}
		peerAddr := peerCtx.Addr.String()
		addrPos := strings.LastIndex(peerAddr, ":")
		clientIP := peerAddr[0:addrPos]

		// 客户端限流
		if appInfo.Qps > 0 {
			clientLimiter, err := public.FlowLimiterHandler.GetLimiter(public.FlowAppPrefix+appInfo.AppID+"_"+clientIP,
				float64(appInfo.Qps))
			if err != nil {
				return err
			}
			if !clientLimiter.Allow() {
				return errors.New(fmt.Sprintf("%v limit %v", clientIP, appInfo.Qps))
			}
		}
		if err := handler(srv, ss); err != nil { // 执行下游方法
			log.Printf("RPC failed with error %v\n", err)
			return err
		}
		return nil
	}
}
