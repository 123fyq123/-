package controller

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	dao "fyqcode.top/go_gateway/dao"
	dto "fyqcode.top/go_gateway/dto"
	"fyqcode.top/go_gateway/golang_common/lib"
	middleware "fyqcode.top/go_gateway/middleware"
	public "fyqcode.top/go_gateway/public"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

type OAuthController struct{}

func OAuthRegister(group *gin.RouterGroup) {
	oauth := &OAuthController{}
	group.POST("/tokens", oauth.Tokens)
}

// Tokens godoc
// @Summary 获取Token
// @Description 获取Token
// @Tags OAuth
// @ID /oauth/tokens
// @Accept  json
// @Produce  json
// @Param body body dto.TokensInput true "body"
// @Success 200 {object} middleware.Response{data=dto.TokensOutput} "success"
// @Router /oauth/token [post]
func (oauth *OAuthController) Tokens(c *gin.Context) {
	params := &dto.TokensInput{}
	if err := params.BindValidParam(c); err != nil {
		middleware.ResponseError(c, 2000, err)
		return
	}

	splits := strings.Split(c.GetHeader("Authorization"), " ")
	if len(splits) != 2 {
		middleware.ResponseError(c, 2001, errors.New("用户名或密码错误"))
		return
	}

	appSecret, err := base64.StdEncoding.DecodeString(splits[1])
	if err != nil {
		middleware.ResponseError(c, 2002, err)
		return
	}

	// 取出app_id secret
	// 生成app_list
	// 匹配app_id
	// 基于jwt生成token
	// 生成output

	parts := strings.Split(string(appSecret), ":")
	if len(parts) != 2 {
		middleware.ResponseError(c, 2003, errors.New("用户名或密码错误"))
		return
	}

	appList := dao.AppManagerHandler.GetAppList()
	for _, appInfo := range appList {
		if appInfo.AppID == parts[0] && appInfo.Secret == parts[1] {
			claims := jwt.StandardClaims{
				Issuer:    appInfo.AppID,
				ExpiresAt: time.Now().Add(public.JwtExpires * time.Second).In(lib.TimeLocation).Unix(),
			}
			token, err := public.JwtEncode(claims)
			if err != nil {
				middleware.ResponseError(c, 2004, err)
				return
			}

			output := &dto.TokensOutput{
				ExpiresIn:   public.JwtExpires,
				TokenType:   "Bearer",
				AccessToken: token,
				Scope:       "read_write",
			}
			middleware.ResponseSuccess(c, output)
			return
		}
	}

	middleware.ResponseError(c, 2005, errors.New("未匹配到正确的APP信息"))
}

// ChangePwd godoc
// @Summary 修改密码
// @Description 修改密码
// @Tags 管理员接口
// @ID /admin/change_pwd
// @Accept  json
// @Produce  json
// @Param body body dto.ChangePwdInput true "body"
// @Success 200 {object} middleware.Response{data=string} "success"
// @Router /admin/change_pwd [post]
func (adminlogin *OAuthController) ChangePwd(c *gin.Context) {
	params := &dto.ChangePwdInput{}
	if err := params.BindValidParam(c); err != nil {
		middleware.ResponseError(c, 2000, err)
		return
	}

	//1. session读取用户信息到结构体 sessInfo
	//2. sessInfo.ID 读取数据库信息 adminInfo
	//3. params.password+adminInfo.salt sha256 saltPassword
	//4. saltPassword==> adminInfo.password 执行数据保存

	//session读取用户信息到结构体
	sess := sessions.Default(c)
	sessInfo := sess.Get(public.AdminSessionInfoKey)
	adminSessionInfo := &dto.AdminSessionInfo{}
	if err := json.Unmarshal([]byte(fmt.Sprint(sessInfo)), adminSessionInfo); err != nil {
		middleware.ResponseError(c, 2000, err)
		return
	}

	//从数据库中读取 adminInfo
	tx, err := lib.GetGormPool("default")
	if err != nil {
		middleware.ResponseError(c, 2001, err)
		return
	}
	adminInfo := &dao.Admin{}
	adminInfo, err = adminInfo.Find(c, tx, (&dao.Admin{UserName: adminSessionInfo.UserName}))
	if err != nil {
		middleware.ResponseError(c, 2002, err)
		return
	}

	//生成新密码 saltPassword
	saltPassword := public.GenSaltPassword(adminInfo.Salt, params.Password)
	adminInfo.Password = saltPassword

	//执行数据保存
	if err := adminInfo.Save(c, tx); err != nil {
		middleware.ResponseError(c, 2003, err)
		return
	}
	middleware.ResponseSuccess(c, "")
}
