package v1

import (
	"github.com/e421083458/golang_common/lib"
	"github.com/gin-gonic/gin"
	"github.com/liuzw3018/gateway_admin/dao"
	"github.com/liuzw3018/gateway_admin/dto"
	"github.com/liuzw3018/gateway_admin/middleware"
	"github.com/liuzw3018/gateway_admin/public"
	"github.com/pkg/errors"
	"time"
)

/**
 * @Author: liu zw
 * @Date: 2022/10/18 18:58
 * @File: app.go
 * @Description: //TODO $
 * @Version:
 */

type APPApi struct{}

func APPApiRegister(router *gin.RouterGroup) {
	aar := APPApi{}
	aApi := router.Group("/app")
	{
		aApi.GET("/list", aar.List)
		aApi.GET("/detail", aar.Detail)
		aApi.GET("/stat", aar.Statistics)
		aApi.GET("/delete", aar.Delete)
		aApi.POST("/add", aar.Add)
		aApi.POST("/update", aar.Update)
	}
}

// List godoc
// @Summary 租户列表
// @Description 租户列表
// @Tags 租户管理接口
// @ID /api/app/list
// @Accept json
// @Produce json
// @Param info query string false "关键词"
// @Param page_size query string true "每页条数"
// @Param page_no query string true "页数"
// @Success 200 {object} middleware.Response{data=dto.APPListOutput} "success"
// @Router /api/app/list [get]
func (a *APPApi) List(c *gin.Context) {
	params := &dto.APPListInput{}
	if err := params.GetValidParams(c); err != nil {
		middleware.ResponseError(c, 1001, err)
		return
	}
	info := &dao.App{}
	list, total, err := info.APPList(c, lib.GORMDefaultPool, params)
	if err != nil {
		middleware.ResponseError(c, 1001, err)
		return
	}
	var outputList []dto.APPListItemOutput
	for _, item := range list {
		realQps := int64(0)
		realQpd := int64(0)
		outputList = append(outputList, dto.APPListItemOutput{
			ID:       item.ID,
			AppID:    item.AppID,
			Name:     item.Name,
			Secret:   item.Secret,
			WhiteIPS: item.WhiteIPS,
			Qpd:      item.Qpd,
			Qps:      item.Qps,
			RealQpd:  realQpd,
			RealQps:  realQps,
		})
	}
	output := dto.APPListOutput{
		List:  outputList,
		Total: total,
	}
	middleware.ResponseSuccess(c, output)
}

// Detail godoc
// @Summary 租户详情
// @Description 租户详情
// @Tags 租户管理接口
// @ID /api/app/detail
// @Accept  json
// @Produce  json
// @Param id query string true "租户ID"
// @Success 200 {object} middleware.Response{data=dao.App} "success"
// @Router /api/app/detail [get]
func (a *APPApi) Detail(c *gin.Context) {
	params := &dto.APPDetailInput{}
	if err := params.GetValidParams(c); err != nil {
		middleware.ResponseError(c, 2001, err)
		return
	}
	search := &dao.App{
		ID: params.ID,
	}
	detail, err := search.Find(c, lib.GORMDefaultPool, search)
	if err != nil {
		middleware.ResponseError(c, 2002, err)
		return
	}
	middleware.ResponseSuccess(c, detail)
	return
}

// Delete godoc
// @Summary 租户删除
// @Description 租户删除
// @Tags 租户管理接口
// @ID /api/app/delete
// @Accept  json
// @Produce  json
// @Param id query string true "租户ID"
// @Success 200 {object} middleware.Response{data=string} "success"
// @Router /api/app/delete [get]
func (a *APPApi) Delete(c *gin.Context) {
	params := &dto.APPDetailInput{}
	if err := params.GetValidParams(c); err != nil {
		middleware.ResponseError(c, 2001, err)
		return
	}
	search := &dao.App{
		ID: params.ID,
	}
	info, err := search.Find(c, lib.GORMDefaultPool, search)
	if err != nil {
		middleware.ResponseError(c, 2002, err)
		return
	}
	info.IsDelete = 1
	if err := info.Save(c, lib.GORMDefaultPool); err != nil {
		middleware.ResponseError(c, 2003, err)
		return
	}
	middleware.ResponseSuccess(c, "")
	return
}

// Add godoc
// @Summary 租户添加
// @Description 租户添加
// @Tags 租户管理接口
// @ID /api/app/add
// @Accept  json
// @Produce  json
// @Param body body dto.APPAddHttpInput true "body"
// @Success 200 {object} middleware.Response{data=string} "success"
// @Router /api/app/add [post]
func (a *APPApi) Add(c *gin.Context) {
	params := &dto.APPAddHttpInput{}
	if err := params.GetValidParams(c); err != nil {
		middleware.ResponseError(c, 2001, err)
		return
	}

	//验证app_id是否被占用
	search := &dao.App{
		AppID: params.AppID,
	}
	if _, err := search.Find(c, lib.GORMDefaultPool, search); err == nil {
		middleware.ResponseError(c, 2002, errors.New("租户ID被占用，请重新输入"))
		return
	}
	if params.Secret == "" {
		params.Secret = public.MD5(params.AppID)
	}
	tx := lib.GORMDefaultPool
	info := &dao.App{
		AppID:    params.AppID,
		Name:     params.Name,
		Secret:   params.Secret,
		WhiteIPS: params.WhiteIPS,
		Qps:      params.Qps,
		Qpd:      params.Qpd,
	}
	if err := info.Save(c, tx); err != nil {
		middleware.ResponseError(c, 2003, err)
		return
	}
	middleware.ResponseSuccess(c, "")
	return
}

// Update godoc
// @Summary 租户更新
// @Description 租户更新
// @Tags 租户管理接口
// @ID /api/app/update
// @Accept  json
// @Produce  json
// @Param body body dto.APPUpdateHttpInput true "body"
// @Success 200 {object} middleware.Response{data=string} "success"
// @Router /api/app/update [post]
func (a *APPApi) Update(c *gin.Context) {
	params := &dto.APPUpdateHttpInput{}
	if err := params.GetValidParams(c); err != nil {
		middleware.ResponseError(c, 2001, err)
		return
	}
	search := &dao.App{
		ID: params.ID,
	}
	info, err := search.Find(c, lib.GORMDefaultPool, search)
	if err != nil {
		middleware.ResponseError(c, 2002, err)
		return
	}
	if params.Secret == "" {
		params.Secret = public.MD5(params.AppID)
	}
	info.Name = params.Name
	info.Secret = params.Secret
	info.WhiteIPS = params.WhiteIPS
	info.Qps = params.Qps
	info.Qpd = params.Qpd
	if err := info.Save(c, lib.GORMDefaultPool); err != nil {
		middleware.ResponseError(c, 2003, err)
		return
	}
	middleware.ResponseSuccess(c, "")
	return
}

// Statistics godoc
// @Summary 租户统计
// @Description 租户统计
// @Tags 租户管理接口
// @ID /api/app/stat
// @Accept  json
// @Produce  json
// @Param id query string true "租户ID"
// @Success 200 {object} middleware.Response{data=dto.StatisticsOutput} "success"
// @Router /api/app/stat [get]
func (a *APPApi) Statistics(c *gin.Context) {
	params := &dto.APPDetailInput{}
	if err := params.GetValidParams(c); err != nil {
		middleware.ResponseError(c, 2001, err)
		return
	}

	search := &dao.App{
		ID: params.ID,
	}
	_, err := search.Find(c, lib.GORMDefaultPool, search)
	if err != nil {
		middleware.ResponseError(c, 2002, err)
		return
	}

	//今日流量全天小时级访问统计
	var todayStat []int64
	for i := 0; i <= time.Now().In(lib.TimeLocation).Hour(); i++ {
		todayStat = append(todayStat, 0)
	}

	//昨日流量全天小时级访问统计
	var yesterdayStat []int64
	for i := 0; i <= 23; i++ {
		yesterdayStat = append(yesterdayStat, 0)
	}
	stat := dto.StatisticsOutput{
		Today:     todayStat,
		Yesterday: yesterdayStat,
	}
	middleware.ResponseSuccess(c, stat)
	return
}
