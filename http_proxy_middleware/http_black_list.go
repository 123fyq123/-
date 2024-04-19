package http_proxy_middleware

import (
	"errors"
	"fmt"
	"strings"

	"fyqcode.top/go_gateway/dao"
	"fyqcode.top/go_gateway/middleware"
	"fyqcode.top/go_gateway/public"
	"github.com/gin-gonic/gin"
)

// ip黑名单，若有白名单则不验证黑名单，若无白名单有黑名单，则只有在黑名单的ip不能访问
func HTTPBlackListMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		serverInterface, ok := c.Get("service")
		if !ok {
			middleware.ResponseError(c, 2001, errors.New("service not found"))
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
		if serviceDetail.AccessControl.OpenAuth == 1 && len(whiteIpList) == 0 && len(blackIpList) > 0 {
			if public.InStringSlice(blackIpList, c.ClientIP()) { // 在黑名单内
				middleware.ResponseError(c, 3001, errors.New(fmt.Sprintf("%s is in black ip list", c.ClientIP())))
				c.Abort()
				return
			}
		}

		c.Next()
	}
}
