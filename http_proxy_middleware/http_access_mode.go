package http_proxy_middleware

import (
	"fmt"

	"fyqcode.top/go_gateway/dao"
	"fyqcode.top/go_gateway/middleware"
	"fyqcode.top/go_gateway/public"
	"github.com/gin-gonic/gin"
)

// 基于请求信息 匹配接入方式
func HTTPAccessModeMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		service, err := dao.ServiceManagerHandler.HTTPAccessMode(c)
		if err != nil {
			middleware.ResponseError(c, 1001, err)
			c.Abort() // 停止中间件
			return
		}
		fmt.Println("matched service", public.Obj2Json(service))
		c.Set("service", service)
		c.Next()
	}
}
