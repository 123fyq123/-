package http_proxy_middleware

import (
	"errors"
	"fmt"

	"fyqcode.top/go_gateway/dao"
	"fyqcode.top/go_gateway/middleware"
	"fyqcode.top/go_gateway/public"
	"github.com/gin-gonic/gin"
)

// JWT限流
func HTTPJwtFlowLimitMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		appInterface, ok := c.Get("app")
		if !ok {
			c.Next()
			return
		}
		appInfo := appInterface.(*dao.App)

		// 客户端限流
		if appInfo.Qps > 0 {
			clientLimiter, err := public.FlowLimiterHandler.GetLimiter(public.FlowAppPrefix+appInfo.AppID+"_"+c.ClientIP(),
				float64(appInfo.Qps))
			if err != nil {
				middleware.ResponseError(c, 5001, err)
				c.Abort()
				return
			}
			if !clientLimiter.Allow() {
				middleware.ResponseError(c, 5002, errors.New(fmt.Sprintf("%v limit %v", c.ClientIP(), appInfo.Qps)))
				c.Abort()
				return
			}
		}
		c.Next()
	}
}
