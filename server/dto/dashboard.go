package dto

/**
 * @Author: liu zw
 * @Date: 2022/10/18 19:54
 * @File: dashboard.go
 * @Description: //TODO $
 * @Version:
 */

type DashboardOutput struct {
	ServiceNum      int64 `json:"service_num"`
	APPNum          int64 `json:"app_num"`
	CurrentQPS      int64 `json:"current_qps"`
	TodayRequestNum int64 `json:"today_request_num"`
}

type DashboardFlowStatOutput struct {
	Today     []int64 `json:"today" form:"today" comment:"今日流量"`         // 今日流量
	Yesterday []int64 `json:"yesterday" form:"yesterday" comment:"昨日流量"` // 昨日流量
}

type DashboardServiceStatOutput struct {
	Legend []string                         `json:"legend" form:"legend" comment:"分类"` // 分类
	Data   []DashboardServiceStatItemOutput `json:"data" form:"data" comment:"数据"`     // 数据
}

type DashboardServiceStatItemOutput struct {
	Name     string `json:"name" form:"name" comment:"分类名"`           // 分类名
	LoadType int    `json:"load_type"  form:"load_type" comment:"类型"` // 类型
	Value    int64  `json:"value" form:"value" comment:"值"`           // 值
}
