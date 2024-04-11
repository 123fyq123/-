package dto

import (
	"time"

	"github.com/e421083458/go_gateway_demo/public"
	"github.com/gin-gonic/gin"
)

type AdminSessionInfo struct {
	ID        int       `json:"id"`
	UserName  string    `json:"user_name"`
	LoginTime time.Time `json:"login_time"`
}

type AdminLoginInput struct {
	UserName string `form:"username" json:"username" comment:"用户名"  validate:"required,is_valid_username" example:"admin"` // 用户名
	Password string `form:"password" json:"password" comment:"密码"  validate:"required" example:"123456"`                   // 密码
}

// 绑定结构体并校验参数
func (param *AdminLoginInput) BindValidParam(c *gin.Context) error {
	return public.DefaultGetValidParams(c, param)
}

type AdminLoginOutput struct {
	Token string `form:"token" json:"token" comment:"token"  validate:"" example:"token"` // token
}
