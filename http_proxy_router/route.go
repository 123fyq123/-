package http_proxy_router

import (
	"fyqcode.top/go_gateway/http_proxy_middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter(middlewares ...gin.HandlerFunc) *gin.Engine {
	router := gin.Default()
	router.Use(middlewares...) // 使用中间件
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.Use(http_proxy_middleware.HTTPAccessModeMiddleware())
	return router
}
