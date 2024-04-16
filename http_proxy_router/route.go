package http_proxy_router

import (
	"fyqcode.top/go_gateway/http_proxy_middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter(middlewares ...gin.HandlerFunc) *gin.Engine {
	// 优化1
	// router := gin.Default()
	router := gin.New()
	router.Use(middlewares...) // 使用中间件
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.Use( // 使用中间件
		http_proxy_middleware.HTTPAccessModeMiddleware(),     // 服务接入
		http_proxy_middleware.HTTPFlowCountMiddleware(),      // 流量统计
		http_proxy_middleware.HTTPFlowLimitMiddleware(),      // 限流
		http_proxy_middleware.HTTPWhiteListMiddleware(),      // ip白名单
		http_proxy_middleware.HTTPBlackListMiddleware(),      // ip黑名单
		http_proxy_middleware.HTTPHeaderTransferMiddleware(), // heaer头转换
		http_proxy_middleware.HTTPStripUriMiddleware(),       // stripuri
		http_proxy_middleware.HTTPUrlRewriteMiddleware(),     // url重写
		http_proxy_middleware.HTTPReverseProxyMiddleware(),   // 反向代理
	)
	return router
}
