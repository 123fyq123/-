package dto

import (
	"github.com/e421083458/go_gateway_demo/public"
	"github.com/gin-gonic/gin"
)

type AdminLoginInput struct {
	UserName string `form:"username" json:"username" comment:"用户名"  validate:"required,is_valid_username" example:"admin"`
	Password string `form:"password" json:"password" comment:"密码"  validate:"required" example:"123456"`
}

// 绑定结构体并校验参数
func (param *AdminLoginInput) BindValidParam(c *gin.Context) error {
	return public.DefaultGetValidParams(c, param)
}
