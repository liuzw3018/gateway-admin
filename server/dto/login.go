package dto

import (
	"github.com/gin-gonic/gin"
	"github.com/liuzw3018/gateway_admin/public"
	"time"
)

/**
 * @Author: liu zw
 * @Date: 2022/10/17 11:50
 * @File: admin.go
 * @Description: //TODO $
 * @Version:
 */

type AdminSessionInfo struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	LoginTime time.Time `json:"login_time"`
}

type AdminLoginInput struct {
	Username string `json:"username" form:"username" comment:"用户名" example:"admin" validate:"required,validUsername"` // 用户名
	Password string `json:"password" form:"password" comment:"密码" example:"123456" validate:"required"`               // 密码
}

func (a *AdminLoginInput) BindValidParam(c *gin.Context) error {
	return public.DefaultGetValidParams(c, a)
}

type AdminLoginOutput struct {
	Token string `json:"token" form:"token" comment:"token" example:"token" validate:""` // token
}
