package http_proxy_middleware

import (
	"errors"
	"fmt"

	"fyqcode.top/go_gateway/dao"
	"fyqcode.top/go_gateway/middleware"
	"fyqcode.top/go_gateway/public"
	"github.com/gin-gonic/gin"
)

// JWT租户流量统计
func HTTPJwtFlowCountMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// serverInterface, ok := c.Get("service")
		// if !ok {
		// 	middleware.ResponseError(c, 2001, errors.New("service not found"))
		// 	c.Abort()
		// 	return
		// }
		// serviceDetail := serverInterface.(*dao.ServiceDetail)

		appInterface, ok := c.Get("app")
		if !ok {
			c.Next()
			return
		}
		appInfo := appInterface.(*dao.App)
		appCounter, err := public.FlowCounterHandler.GetCounter(public.FlowAppPrefix + appInfo.AppID)
		if err != nil {
			middleware.ResponseError(c, 2002, err)
			c.Abort()
			return
		}
		appCounter.Increase() // 原子增加计数器

		if appInfo.Qpd > 0 && appCounter.TotalCount > appInfo.Qpd {
			middleware.ResponseError(c, 2003, errors.New(fmt.Sprintf("租户日请求量限流 limit:%v", appInfo.Qpd)))
			c.Abort()
			return
		}

		fmt.Printf("appCounter qps:%v  daycount:%v", appCounter.QPS, appCounter.TotalCount)

		c.Next()
	}
}
