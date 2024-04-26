package grpc_proxy_middleware

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"fyqcode.top/go_gateway/dao"
	"fyqcode.top/go_gateway/public"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// JWT租户流量统计
func GrpcJwtFlowCountMiddleware(serviceDetail *dao.ServiceDetail) func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
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
		appCounter, err := public.FlowCounterHandler.GetCounter(public.FlowAppPrefix + appInfo.AppID)
		if err != nil {
			return err
		}
		appCounter.Increase() // 原子增加计数器

		if appInfo.Qpd > 0 && appCounter.TotalCount > appInfo.Qpd {
			return errors.New(fmt.Sprintf("租户日请求量限流 limit:%v", appInfo.Qpd))
		}

		if err := handler(srv, ss); err != nil { // 执行下游方法
			log.Printf("RPC failed with error %v\n", err)
			return err
		}
		return nil
	}
}
