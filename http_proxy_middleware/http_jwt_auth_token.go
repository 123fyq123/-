package http_proxy_middleware

import (
	"errors"
	"strings"

	"fyqcode.top/go_gateway/dao"
	"fyqcode.top/go_gateway/middleware"
	"fyqcode.top/go_gateway/public"
	"github.com/gin-gonic/gin"
)

// jwt token 验证
func HTTPJwtAuthTokenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		serverInterface, ok := c.Get("service")
		if !ok {
			middleware.ResponseError(c, 2001, errors.New("service not found"))
			c.Abort()
			return
		}
		serviceDetail := serverInterface.(*dao.ServiceDetail)

		// decode jwt token信息 获取claims信息，里面又appID
		// appID 与 applist内容进行匹配  取出appInfo信息
		// appInfo 放到 gin.context里面
		// 之后就可以从gin.context里面获取信息  进行流量统计等

		appMatched := false
		token := strings.ReplaceAll(c.GetHeader("Authorization"), "Bearer ", "")
		if token != "" {
			claims, err := public.JwtDecode(token)
			if err != nil {
				middleware.ResponseError(c, 2002, err)
				c.Abort()
				return
			}

			appList := dao.AppManagerHandler.GetAppList()
			for _, appInfo := range appList {
				if appInfo.AppID == claims.Issuer {
					c.Set("app", appInfo)
					appMatched = true
					break
				}
			}
			if serviceDetail.AccessControl.OpenAuth == 1 && !appMatched { // 开启验证但没有匹配到
				middleware.ResponseError(c, 2003, errors.New("not matched valid app"))
				c.Abort()
				return
			}
		}
		c.Next()
	}
}
