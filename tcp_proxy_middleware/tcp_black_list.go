package tcp_proxy_middleware

import (
	"fmt"
	"strings"

	"fyqcode.top/go_gateway/dao"
	"fyqcode.top/go_gateway/public"
)

// ip黑名单，若有白名单则不验证黑名单，若无白名单有黑名单，则只有在黑名单的ip不能访问
func TCPBlackListMiddleware() func(c *TcpSliceRouterContext) {
	return func(c *TcpSliceRouterContext) {
		serverInterface := c.Get("service")
		if serverInterface == nil {
			c.conn.Write([]byte("get service empty"))
			c.Abort()
			return
		}
		serviceDetail := serverInterface.(*dao.ServiceDetail)

		whiteIpList := []string{}
		blackIpList := []string{}
		if serviceDetail.AccessControl.WhiteList != "" {
			whiteIpList = strings.Split(serviceDetail.AccessControl.WhiteList, ",")
		}
		if serviceDetail.AccessControl.BlackList != "" {
			blackIpList = strings.Split(serviceDetail.AccessControl.BlackList, ",")
		}
		splits := strings.Split(c.conn.RemoteAddr().String(), ":")
		clientIP := ""
		if len(splits) == 2 {
			clientIP = splits[0]
		}
		if serviceDetail.AccessControl.OpenAuth == 1 && len(whiteIpList) == 0 && len(blackIpList) > 0 {
			if public.InStringSlice(blackIpList, clientIP) { // 在黑名单内
				c.conn.Write([]byte(fmt.Sprintf("%s is in black ip list", clientIP)))
				c.Abort()
				return
			}
		}

		c.Next()
	}
}
