package v1

import (
	"encoding/json"
	"fmt"
	"github.com/e421083458/golang_common/lib"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/liuzw3018/gateway_admin/dao"
	"github.com/liuzw3018/gateway_admin/dto"
	"github.com/liuzw3018/gateway_admin/middleware"
	"github.com/liuzw3018/gateway_admin/public"
)

/**
 * @Author: liu zw
 * @Date: 2022/10/17 11:46
 * @File: admin.go
 * @Description: //TODO $
 * @Version:
 */

type AdminApi struct{}

func AdminApiRegister(router *gin.RouterGroup) {
	adminLogin := &AdminApi{}
	userApi := router.Group("/admin")
	{
		userApi.GET("/info", adminLogin.Info)
		userChangeApi := userApi.Group("/change")
		{
			userChangeApi.POST("/password", adminLogin.ChangePwd)
			userChangeApi.POST("/info", adminLogin.ChangeInfo)
		}
	}
}

// Info godoc
// @Summary 管理员用户信息
// @Description 管理员用户信息
// @Tags 管理员用户信息接口
// @ID /api/admin/info
// @Accept json
// @Produce json
// @Success 200 {object} middleware.Response{data=dto.AdminInfoOutput} "success"
// @Router /api/admin/info [get]
func (a *AdminApi) Info(c *gin.Context) {
	sess := sessions.Default(c)
	sessInfo := sess.Get(public.SessionKey)
	adminSessInfo := &dto.AdminSessionInfo{}
	if err := json.Unmarshal([]byte(fmt.Sprint(sessInfo)), adminSessInfo); err != nil {
		middleware.ResponseError(c, 2000, err)
		return
	}

	out := &dto.AdminInfoOutput{
		ID:           adminSessInfo.ID,
		Name:         adminSessInfo.Username,
		LoginTime:    adminSessInfo.LoginTime,
		Avatar:       "https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif",
		Introduction: "I am a super administrator",
		Roles:        []string{"admin"},
	}
	middleware.ResponseSuccess(c, out)
}

// ChangePwd godoc
// @Summary 管理员修改密码
// @Description 管理员修改密码
// @Tags 管理员用户信息接口
// @ID /api/admin/change/password
// @Accept json
// @Produce json
// @Param body body dto.ChangePwdInput true "body"
// @Success 200 {object} middleware.Response{data=string} "success"
// @Router /api/admin/change/password [post]
func (a *AdminApi) ChangePwd(c *gin.Context) {
	params := &dto.ChangePwdInput{}
	if err := params.BindValidParam(c); err != nil {
		middleware.ResponseError(c, 1001, err)
		return
	}

	sess := sessions.Default(c)
	sessInfo := sess.Get(public.SessionKey)
	adminSessInfo := &dto.AdminSessionInfo{}
	if err := json.Unmarshal([]byte(fmt.Sprint(sessInfo)), adminSessInfo); err != nil {
		middleware.ResponseError(c, 2000, err)
		return
	}

	tx, err := lib.GetGormPool("default")
	if err != nil {
		middleware.ResponseError(c, 2000, err)
		return
	}

	adminInfo := &dao.Admin{}
	adminInfo, err = adminInfo.Find(c, tx, &dao.Admin{
		Username: adminSessInfo.Username,
		Id:       adminSessInfo.ID})
	if err != nil {
		middleware.ResponseError(c, 2001, err)
		return
	}

	saltPassword := public.GenSaltPassword(adminInfo.Salt, params.Password)
	adminInfo.Password = saltPassword
	err = adminInfo.Save(c, tx)
	if err != nil {
		middleware.ResponseError(c, 2002, err)
		return
	}

	middleware.ResponseSuccess(c, "OK")
}

// ChangeInfo godoc
// @Summary 管理员修改信息
// @Description 管理员修改信息
// @Tags 管理员用户信息接口
// @ID /api/admin/change/info
// @Accept json
// @Produce json
// @Param body body dto.ChangeInfoInput true "body"
// @Success 200 {object} middleware.Response{data=string} "success"
// @Router /api/admin/change/info [post]
func (a *AdminApi) ChangeInfo(c *gin.Context) {
	middleware.ResponseSuccess(c, "")
}
