package http_proxy_middleware

import (
	"errors"

	"fyqcode.top/go_gateway/dao"
	"fyqcode.top/go_gateway/middleware"
	"fyqcode.top/go_gateway/reverse_proxy"
	"github.com/gin-gonic/gin"
)

// 反向代理
func HTTPReverseProxyMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		serverInterface, ok := c.Get("service")
		if !ok {
			middleware.ResponseError(c, 2001, errors.New("service not found"))
			c.Abort()
			return
		}
		serviceDetail := serverInterface.(*dao.ServiceDetail)

		lb, err := dao.LoadBalancerHandler.GetLoadBalancer(serviceDetail)
		if err != nil {
			middleware.ResponseError(c, 2002, err)
			c.Abort()
			return
		}
		trans, err := dao.TransportorHandler.GetTrans(serviceDetail)
		if err != nil {
			middleware.ResponseError(c, 2003, err)
			c.Abort()
			return
		}
		//middleware.ResponseSuccess(c,"ok")
		//return
		//创建 reverseproxy
		//使用 reverseproxy.ServerHTTP(c.Request,c.Response)
		// 创建反向代理
		proxy := reverse_proxy.NewLoadBalanceReverseProxy(c, lb, trans)
		proxy.ServeHTTP(c.Writer, c.Request)
		c.Abort()
		return
	}
}
