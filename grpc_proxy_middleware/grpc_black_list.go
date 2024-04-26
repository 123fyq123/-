package grpc_proxy_middleware

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"fyqcode.top/go_gateway/dao"
	"fyqcode.top/go_gateway/public"
	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
)

// ip黑名单，若有白名单则不验证黑名单，若无白名单有黑名单，则只有在黑名单的ip不能访问
func GrpcBlackListMiddleware(serviceDetail *dao.ServiceDetail) func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		whiteIpList := []string{}
		blackIpList := []string{}

		peerCtx, ok := peer.FromContext(ss.Context())
		if !ok {
			return errors.New("peer not found with context")
		}
		peerAddr := peerCtx.Addr.String()
		addrPos := strings.LastIndex(peerAddr, ":")
		clientIP := peerAddr[0:addrPos]

		if serviceDetail.AccessControl.WhiteList != "" {
			whiteIpList = strings.Split(serviceDetail.AccessControl.WhiteList, ",")
		}
		if serviceDetail.AccessControl.BlackList != "" {
			blackIpList = strings.Split(serviceDetail.AccessControl.BlackList, ",")
		}
		if serviceDetail.AccessControl.OpenAuth == 1 && len(whiteIpList) == 0 && len(blackIpList) > 0 {
			if public.InStringSlice(blackIpList, clientIP) { // 在黑名单内
				return errors.New(fmt.Sprintf("%s is in black ip list", clientIP))
			}
		}
		if err := handler(srv, ss); err != nil {
			log.Printf("RPC failed with error %v\n", err)
			return err
		}
		return nil

	}
}
