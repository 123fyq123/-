package grpc_proxy_middleware

import (
	"errors"
	"log"
	"strings"

	"fyqcode.top/go_gateway/dao"
	"fyqcode.top/go_gateway/public"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// jwt token 验证
func GrpcJwtAuthTokenMiddleware(serviceDetail *dao.ServiceDetail) func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		md, ok := metadata.FromIncomingContext(ss.Context())
		if !ok {
			return errors.New("miss metadata from context")
		}

		authToken := ""
		auth := md.Get("Authorization")
		if len(auth) > 0 { // 对于没开启权限验证的没有token
			authToken = auth[0]
		}
		appMatched := false
		token := strings.ReplaceAll(authToken, "Bearer ", "")
		if token != "" {
			claims, err := public.JwtDecode(token)
			if err != nil {
				return errors.New("JwtDecode err")
			}

			appList := dao.AppManagerHandler.GetAppList()
			for _, appInfo := range appList {
				if appInfo.AppID == claims.Issuer {
					md.Set("app", public.Obj2Json(appInfo))
					appMatched = true // 租户匹配成功
					break
				}
			}
			if serviceDetail.AccessControl.OpenAuth == 1 && !appMatched { // 开启验证但没有匹配到
				return errors.New("not matched valid app")
			}
		}
		if err := handler(srv, ss); err != nil {
			log.Printf("RPC failed with error %v\n", err)
			return err
		}
		return nil
	}

}
