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
 * @Date: 2022/10/18 09:30
 * @File: dashboard.go
 * @Description: //TODO $
 * @Version:
 */

type DashboardApi struct{}

func DashboardApiRegister(router *gin.RouterGroup) {
	srv := &DashboardApi{}
	srvApi := router.Group("/dashboard")
	{
		srvApi.GET("/panelGroupData", srv.PanelGroupData)
		srvApi.GET("/flowStat", srv.FlowStat)
		srvApi.GET("/serviceStat", srv.ServiceStat)
	}
}

// PanelGroupData godoc
// @Summary 指标统计
// @Description 指标统计
// @Tags 首页大盘接口
// @ID /api/dashboard/panelGroupData
// @Accept json
// @Produce json
// @Success 200 {object} middleware.Response{data=dto.DashboardOutput} "success"
// @Router /api/dashboard/panelGroupData [get]
func (d *DashboardApi) PanelGroupData(c *gin.Context) {
	serviceInfo := &dao.ServiceInfo{}
	_, serviceNum, err := serviceInfo.PageList(c, lib.GORMDefaultPool, &dto.ServiceListInput{PageSize: 1, PageNo: 1})
	if err != nil {
		middleware.ResponseError(c, 2002, err)
		return
	}

	appInfo := &dao.App{}
	_, appNum, err := appInfo.APPList(c, lib.GORMDefaultPool, &dto.APPListInput{PageNo: 1, PageSize: 1})
	if err != nil {
		middleware.ResponseError(c, 2002, err)
		return
	}

	out := &dto.DashboardOutput{
		ServiceNum:      serviceNum,
		APPNum:          appNum,
		TodayRequestNum: 0,
		CurrentQPS:      0,
	}
	middleware.ResponseSuccess(c, out)
}

// FlowStat godoc
// @Summary 流量统计
// @Description 流量统计
// @Tags 首页大盘接口
// @ID /api/dashboard/flowStat
// @Accept json
// @Produce json
// @Success 200 {object} middleware.Response{data=dto.DashboardFlowStatOutput} "success"
// @Router /api/dashboard/flowStat [get]
func (d *DashboardApi) FlowStat(c *gin.Context) {
	// 统计数据
	var todayList []int64
	var yesterdayList []int64
	for i := 0; i <= time.Now().Hour(); i++ {
		todayList = append(todayList, 0)
	}
	for i := 0; i <= 23; i++ {
		yesterdayList = append(yesterdayList, 0)
	}

	middleware.ResponseSuccess(c, &dto.DashboardFlowStatOutput{
		Today:     todayList,
		Yesterday: yesterdayList,
	})
}

// ServiceStat godoc
// @Summary 服务统计
// @Description 服务统计
// @Tags 首页大盘接口
// @ID /api/dashboard/serviceStat
// @Accept json
// @Produce json
// @Success 200 {object} middleware.Response{data=dto.DashboardServiceStatOutput} "success"
// @Router /api/dashboard/serviceStat [get]
func (d *DashboardApi) ServiceStat(c *gin.Context) {

	serviceInfo := &dao.ServiceInfo{}
	list, err := serviceInfo.GroupByLoadType(c, lib.GORMDefaultPool)
	if err != nil {
		middleware.ResponseError(c, 2000, err)
		return
	}

	var legend []string
	for index, item := range list {
		name, ok := public.LoadTypeMap[item.LoadType]
		if !ok {
			middleware.ResponseError(c, 2003, errors.New("load_type not found"))
			return
		}
		legend = append(legend, name)
		list[index].Name = name
	}

	out := &dto.DashboardServiceStatOutput{
		Legend: legend,
		Data:   list,
	}
	middleware.ResponseSuccess(c, out)
}
