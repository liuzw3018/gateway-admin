package v1

import (
	"encoding/json"
	"github.com/e421083458/golang_common/lib"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/liuzw3018/gateway_admin/dao"
	"github.com/liuzw3018/gateway_admin/dto"
	"github.com/liuzw3018/gateway_admin/middleware"
	"github.com/liuzw3018/gateway_admin/public"
	"time"
)

/**
 * @Author: liu zw
 * @Date: 2022/10/17 11:46
 * @File: admin.go
 * @Description: //TODO $
 * @Version:
 */

type AdminLoginApi struct{}

func AdminLoginApiRegister(router *gin.RouterGroup) {
	adminLogin := &AdminLoginApi{}
	userApi := router.Group("/admin")
	{
		userApi.POST("/login", adminLogin.Login)
		userApi.GET("/logout", adminLogin.Logout)
	}
}

// Login godoc
// @Summary 管理员登录
// @Description 管理员登录
// @Tags 管理员接口
// @ID /api/admin/login
// @Accept json
// @Produce json
// @Param body body dto.AdminLoginInput true "body"
// @Success 200 {object} middleware.Response{data=dto.AdminLoginOutput} "success"
// @Router /api/admin/login [post]
func (a *AdminLoginApi) Login(c *gin.Context) {
	params := &dto.AdminLoginInput{}
	if err := params.BindValidParam(c); err != nil {
		middleware.ResponseError(c, 1001, err)
		return
	}
	tx, err := lib.GetGormPool("default")
	if err != nil {
		middleware.ResponseError(c, 2000, err)
		return
	}

	admin := &dao.Admin{}
	admin, err = admin.LoginCheck(c, tx, params)
	if err != nil {
		middleware.ResponseError(c, 2001, err)
		return
	}

	// session
	sessInfo := &dto.AdminSessionInfo{
		ID:        admin.Id,
		Username:  admin.Username,
		LoginTime: time.Now(),
	}
	sessBts, err := json.Marshal(&sessInfo)
	if err != nil {
		middleware.ResponseError(c, 2002, err)
		return
	}
	sess := sessions.Default(c)
	sess.Set(public.SessionKey, string(sessBts))
	sess.Save()

	out := &dto.AdminLoginOutput{Token: admin.Username}
	middleware.ResponseSuccess(c, out)
}

// Logout godoc
// @Summary 管理员退出
// @Description 管理员退出
// @Tags 管理员接口
// @ID /api/admin/logout
// @Accept json
// @Produce json
// @Success 200 {object} middleware.Response{data=string} "success"
// @Router /api/admin/logout [get]
func (a *AdminLoginApi) Logout(c *gin.Context) {
	// session
	sess := sessions.Default(c)
	sess.Delete(public.SessionKey)
	sess.Save()

	middleware.ResponseSuccess(c, "OK")
}
