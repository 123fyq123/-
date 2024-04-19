package http_proxy_router

import (
	"fyqcode.top/go_gateway/controller"
	"fyqcode.top/go_gateway/http_proxy_middleware"
	"fyqcode.top/go_gateway/middleware"
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

	oauth := router.Group("oauth")
	oauth.Use(
		middleware.TranslationMiddleware(),
	)

	router.Use( // 使用中间件
		http_proxy_middleware.HTTPAccessModeMiddleware(), // 服务接入

		// 流量控制
		http_proxy_middleware.HTTPFlowCountMiddleware(), // 流量统计
		http_proxy_middleware.HTTPFlowLimitMiddleware(), // 限流

		// 权限校验
		http_proxy_middleware.HTTPJwtAuthTokenMiddleware(), // jwt验证
		http_proxy_middleware.HTTPJwtFlowCountMiddleware(), // 租户流量统计

		http_proxy_middleware.HTTPWhiteListMiddleware(), // ip白名单
		http_proxy_middleware.HTTPBlackListMiddleware(), // ip黑名单

		// 重写
		http_proxy_middleware.HTTPHeaderTransferMiddleware(), // heaer头转换
		http_proxy_middleware.HTTPStripUriMiddleware(),       // stripuri
		http_proxy_middleware.HTTPUrlRewriteMiddleware(),     // url重写

		http_proxy_middleware.HTTPReverseProxyMiddleware(), // 反向代理
	)

	{
		controller.OAuthRegister(oauth)
	}
	return router
}
