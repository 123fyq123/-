package tcp_proxy_middleware

import (
	"fyqcode.top/go_gateway/dao"
	"fyqcode.top/go_gateway/public"
)

// 流量统计
func TCPFlowCountMiddleware() func(c *TcpSliceRouterContext) {
	return func(c *TcpSliceRouterContext) {
		serverInterface := c.Get("service")
		if serverInterface == nil {
			c.conn.Write([]byte("get service empty"))
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
			c.conn.Write([]byte(err.Error()))
			c.Abort()
			return
		}
		totalCounter.Increase() // 原子增加计数器
		// dayCount, _ := totalCounter.GetDayData(time.Now())
		// fmt.Printf("totalCounter qps:%v  daycount:%v", totalCounter.QPS, dayCount)

		// 服务
		serviceCounter, err := public.FlowCounterHandler.GetCounter(public.FlowServicePrefix + serviceDetail.Info.ServiceName)
		if err != nil {
			c.conn.Write([]byte(err.Error()))
			c.Abort()
			return
		}
		serviceCounter.Increase() // 原子增加计数器

		// dayServiceCount, _ := serviceCounter.GetDayData(time.Now())
		// fmt.Printf("serviceCounter qps:%v  daycount:%v", serviceCounter.QPS, dayServiceCount)

		c.Next()
	}
}
