package http_proxy_middleware

import (
	"errors"

	"fyqcode.top/go_gateway/dao"
	"fyqcode.top/go_gateway/middleware"
	"fyqcode.top/go_gateway/public"
	"github.com/gin-gonic/gin"
)

// 流量统计
func HTTPFlowCountMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		serverInterface, ok := c.Get("service")
		if !ok {
			middleware.ResponseError(c, 2001, errors.New("service not found"))
			c.Abort()
			return
		}
		serviceDetail := serverInterface.(*dao.ServiceDetail)

		// 1.全站统计
		// 2.服务统计
		// 3.租户统计

		// 全站
		totalCounter, err := public.FlowCounterHandler.GetCounter(public.FlowTotal)
		if err != nil {
			middleware.ResponseError(c, 4001, err)
			c.Abort()
			return
		}
		totalCounter.Increase() // 原子增加计数器
		// dayCount, _ := totalCounter.GetDayData(time.Now())
		// fmt.Printf("totalCounter qps:%v  daycount:%v", totalCounter.QPS, dayCount)

		// 服务
		serviceCounter, err := public.FlowCounterHandler.GetCounter(public.FlowServicePrefix + serviceDetail.Info.ServiceName)
		if err != nil {
			middleware.ResponseError(c, 4002, err)
			c.Abort()
			return
		}
		serviceCounter.Increase() // 原子增加计数器

		// dayServiceCount, _ := serviceCounter.GetDayData(time.Now())
		// fmt.Printf("serviceCounter qps:%v  daycount:%v", serviceCounter.QPS, dayServiceCount)

		c.Next()
	}
}
