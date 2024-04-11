package dto

import (
	"time"

	"github.com/e421083458/go_gateway_demo/public"
	"github.com/gin-gonic/gin"
)

type AdminInfoOutput struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	LoginTime    time.Time `json:"login_time"`
	Avatar       string    `json:"avatar"`
	Introduction string    `json:"introduction"`
	Roles        []string  `json:"roles"`
}

type ChangePwdInput struct {
	Password string `form:"password" json:"password" comment:"密码"  validate:"required" example:"123456"` // 密码
}

// 绑定结构体并校验参数
func (param *ChangePwdInput) BindValidParam(c *gin.Context) error {
	return public.DefaultGetValidParams(c, param)
}
