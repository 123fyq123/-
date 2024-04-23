package tcp_proxy_middleware

import (
	"fmt"
	"strings"

	"fyqcode.top/go_gateway/dao"
	"fyqcode.top/go_gateway/public"
)

// ip白名单
func TCPWhiteListMiddleware() func(c *TcpSliceRouterContext) {
	return func(c *TcpSliceRouterContext) {
		serverInterface := c.Get("service")
		if serverInterface == nil {
			c.conn.Write([]byte("get service empty"))
			c.Abort()
			return
		}
		serviceDetail := serverInterface.(*dao.ServiceDetail)

		splits := strings.Split(c.conn.RemoteAddr().String(), ":")
		clientIP := ""
		if len(splits) == 2 {
			clientIP = splits[0]
		}

		iplist := []string{}
		if serviceDetail.AccessControl.WhiteList != "" {
			iplist = strings.Split(serviceDetail.AccessControl.WhiteList, ",")
		}
		if serviceDetail.AccessControl.OpenAuth == 1 && len(iplist) > 0 {
			if !public.InStringSlice(iplist, clientIP) { // 不在ip白名单内
				c.conn.Write([]byte(fmt.Sprintf("%s is not in white ip list", clientIP)))
				c.Abort()
				return
			}
		}

		c.Next()
	}
}
