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

type AdminInfoOutput struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	LoginTime    time.Time `json:"login_time"`
	Avatar       string    `json:"avatar"`
	Introduction string    `json:"introduction"`
	Roles        []string  `json:"roles"`
}

type ChangePwdInput struct {
	//OldPassword string `json:"old_password" form:"old_password" comment:"旧密码" example:"123456" validate:"required"` // 旧密码
	Password string `json:"password" form:"password" comment:"密码" example:"123456" validate:"required"` // 密码
}

func (cpi *ChangePwdInput) BindValidParam(c *gin.Context) error {
	return public.DefaultGetValidParams(c, cpi)
}

type ChangeInfoInput struct {
}
