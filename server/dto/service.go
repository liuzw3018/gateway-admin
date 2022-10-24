package dto

import (
	"github.com/gin-gonic/gin"
	"github.com/liuzw3018/gateway_admin/public"
)

/**
 * @Author: liu zw
 * @Date: 2022/10/18 09:33
 * @File: service_info.go
 * @Description: //TODO $
 * @Version:
 */

type ServiceListInput struct {
	Info  string `json:"info" form:"info" comment:"关键词"`                                     // 关键词
	Page  int    `json:"page" form:"page" comment:"页数" example:"1" validate:"required"`      // 页数
	Limit int    `json:"limit" form:"limit" comment:"每页条数" example:"20" validate:"required"` // 每页条数
}

func (sli *ServiceListInput) BindValidParam(c *gin.Context) error {
	return public.DefaultGetValidParams(c, sli)
}

type ServiceListOutput struct {
	Total int64                   `json:"total" form:"total" comment:"总数"` // 总数
	List  []ServiceListItemOutput `json:"list" form:"list" comment:"列表"`   // 列表
}

type ServiceListItemOutput struct {
	ID          int64  `json:"id" form:"id" comment:"id"`                               // id
	ServiceName string `json:"service_name" form:"service_name" comment:"service_name"` // 服务名称
	ServiceDesc string `json:"service_desc" form:"service_desc" comment:"service_desc"` // 服务描述
	LoadType    int    `json:"load_type" form:"load_type" comment:"load_type"`          // 服务类型
	ServiceAddr string `json:"service_addr" form:"service_addr" comment:"service_addr"` // 服务地址
	QPS         int64  `json:"qps" form:"qps" comment:"qps"`                            // qps
	QPD         int64  `json:"qpd" form:"qpd" comment:"qpd"`                            // qpd
	TotalNode   int    `json:"total_node" form:"total_node" comment:"total_node"`       // 节点数
}

type ServiceDeleteInput struct {
	ID int64 `json:"id" form:"id" comment:"服务ID" example:"56" validate:"required"` // 服务ID
}

func (sdi *ServiceDeleteInput) BindValidParam(c *gin.Context) error {
	return public.DefaultGetValidParams(c, sdi)
}

type ServiceAddHTTPInput struct {
	ServiceName            string `json:"service_name" form:"service_name" comment:"服务名称" example:"service_name" validate:"required,validServiceName"`         // 服务名称
	ServiceDesc            string `json:"service_desc" form:"service_desc" comment:"服务描述" example:"service_desc" validate:"required,max=255,min=1"`            // 服务描述
	RuleType               int    `json:"rule_type" form:"rule_type" comment:"接入类型" example:"0" validate:"max=1,min=0"`                                        // 接入类型
	Rule                   string `json:"rule" form:"rule" comment:"接入路径:域名或前缀" example:"rule" validate:"required,validRule"`                                  // 接入路径
	NeedHTTPS              int    `json:"need_https" form:"need_https" comment:"是否支持https" example:"0" validate:"max=1,min=0"`                                 // 是否支持https
	NeedStripURI           int    `json:"need_strip_uri" form:"need_strip_uri" comment:"启用strip_uri" example:"0" validate:"max=1,min=0"`                       // 启用strip_uri
	NeedWebsocket          int    `json:"need_websocket" form:"need_websocket" comment:"是否支持websocket" example:"0" validate:"max=1,min=0"`                     // 是否支持websocket
	URLRewrite             string `json:"url_rewrite" form:"url_rewrite" comment:"URL重写功能" example:"url_rewrite" validate:"validURLRewrite"`                   // URL重写功能
	HeaderTransfor         string `json:"header_transfor" form:"header_transfor" comment:"header头转换" example:"header_transfor" validate:"validHeaderTransfor"` // header头转换
	OpenAuth               int    `json:"open_auth" form:"open_auth" comment:"是否开启权限" example:"0" validate:"max=1,min=0"`                                      // 是否开启权限
	BlackList              string `json:"black_list" form:"black_list" comment:"黑名单" example:"black_list" validate:""`                                         // 黑名单
	WhiteList              string `json:"white_list" form:"white_list" comment:"白名单" example:"white_list" validate:""`                                         // 白名单
	ClientIPFlowLimit      int    `json:"clientip_flow_limit" form:"clientip_flow_limit" comment:"客户端IP限流" example:"0" validate:"min=0"`                       // 客户端IP限流
	ServiceFlowLimit       int    `json:"service_flow_limit" form:"service_flow_limit" comment:"服务端限流" example:"0" validate:"min=0"`                           // 服务端限流
	RoundType              int    `json:"round_type" form:"round_type" comment:"轮询方式" example:"0" validate:"max=3,min=0"`                                      // 轮询方式
	IpList                 string `json:"ip_list" form:"ip_list" comment:"服务IP列表" example:"ip_list" validate:"required,validIpList"`                           // 服务IP列表
	WeightList             string `json:"weight_list" form:"weight_list" comment:"权重列表" example:"weight_list" validate:"validWeightList"`                      // 权重列表
	UpStreamConnectTimeout int    `json:"upstream_connect_timeout" form:"upstream_connect_timeout" comment:"建立连接超时:秒" example:"0" validate:"min=0"`            // 建立连接超时：秒
	UpStreamHeaderTimeout  int    `json:"upstream_header_timeout" form:"upstream_header_timeout" comment:"获取header超时:秒" example:"0" validate:"min=0"`          // 获取header超时：秒
	UpStreamIdleTimeout    int    `json:"upstream_idle_timeout" form:"upstream_idle_timeout" comment:"链接最大空闲时间:秒" example:"0" validate:"min=0"`                // 链接最大空闲时间：秒
	UpStreamMaxIdle        int    `json:"upstream_max_idle" form:"upstream_max_idle" comment:"最大空闲链接数" example:"0" validate:"min=0"`                           // 最大空闲链接数
}

func (sahi *ServiceAddHTTPInput) BindValidParam(c *gin.Context) error {
	return public.DefaultGetValidParams(c, sahi)
}

type ServiceUpdateHTTPInput struct {
	ID                     int64  `json:"id" form:"id" comment:"服务ID" example:"64" validate:"required,min=1"`                                                  // 服务ID
	ServiceName            string `json:"service_name" form:"service_name" comment:"服务名称" example:"service_name" validate:"required,validServiceName"`         // 服务名称
	ServiceDesc            string `json:"service_desc" form:"service_desc" comment:"服务描述" example:"service_desc" validate:"required,max=255,min=1"`            // 服务描述
	RuleType               int    `json:"rule_type" form:"rule_type" comment:"接入类型" example:"0" validate:"max=1,min=0"`                                        // 接入类型
	Rule                   string `json:"rule" form:"rule" comment:"接入路径:域名或前缀" example:"rule" validate:"required,validRule"`                                  // 接入路径
	NeedHTTPS              int    `json:"need_https" form:"need_https" comment:"是否支持https" example:"0" validate:"max=1,min=0"`                                 // 是否支持https
	NeedStripURI           int    `json:"need_strip_uri" form:"need_strip_uri" comment:"启用strip_uri" example:"0" validate:"max=1,min=0"`                       // 启用strip_uri
	NeedWebsocket          int    `json:"need_websocket" form:"need_websocket" comment:"是否支持websocket" example:"0" validate:"max=1,min=0"`                     // 是否支持websocket
	URLRewrite             string `json:"url_rewrite" form:"url_rewrite" comment:"URL重写功能" example:"url_rewrite" validate:"validURLRewrite"`                   // URL重写功能
	HeaderTransfor         string `json:"header_transfor" form:"header_transfor" comment:"header头转换" example:"header_transfor" validate:"validHeaderTransfor"` // header头转换
	OpenAuth               int    `json:"open_auth" form:"open_auth" comment:"是否开启权限" example:"0" validate:"max=1,min=0"`                                      // 是否开启权限
	BlackList              string `json:"black_list" form:"black_list" comment:"黑名单" example:"black_list" validate:""`                                         // 黑名单
	WhiteList              string `json:"white_list" form:"white_list" comment:"白名单" example:"white_list" validate:""`                                         // 白名单
	ClientIPFlowLimit      int    `json:"clientip_flow_limit" form:"clientip_flow_limit" comment:"客户端IP限流" example:"0" validate:"min=0"`                       // 客户端IP限流
	ServiceFlowLimit       int    `json:"service_flow_limit" form:"service_flow_limit" comment:"服务端限流" example:"0" validate:"min=0"`                           // 服务端限流
	RoundType              int    `json:"round_type" form:"round_type" comment:"轮询方式" example:"0" validate:"max=3,min=0"`                                      // 轮询方式
	IpList                 string `json:"ip_list" form:"ip_list" comment:"服务IP列表" example:"ip_list" validate:"required,validIpList"`                           // 服务IP列表
	WeightList             string `json:"weight_list" form:"weight_list" comment:"权重列表" example:"weight_list" validate:"validWeightList"`                      // 权重列表
	UpStreamConnectTimeout int    `json:"upstream_connect_timeout" form:"upstream_connect_timeout" comment:"建立连接超时:秒" example:"0" validate:"min=0"`            // 建立连接超时：秒
	UpStreamHeaderTimeout  int    `json:"upstream_header_timeout" form:"upstream_header_timeout" comment:"获取header超时:秒" example:"0" validate:"min=0"`          // 获取header超时：秒
	UpStreamIdleTimeout    int    `json:"upstream_idle_timeout" form:"upstream_idle_timeout" comment:"链接最大空闲时间:秒" example:"0" validate:"min=0"`                // 链接最大空闲时间：秒
	UpStreamMaxIdle        int    `json:"upstream_max_idle" form:"upstream_max_idle" comment:"最大空闲链接数" example:"0" validate:"min=0"`                           // 最大空闲链接数
}

func (shhi *ServiceUpdateHTTPInput) BindValidParam(c *gin.Context) error {
	return public.DefaultGetValidParams(c, shhi)
}

type ServiceStatOutput struct {
	ServiceName string  `json:"service_name"  form:"service_name" comment:"服务名称"`
	Today       []int64 `json:"today" form:"today" comment:"今日流量"`         // 今日流量
	Yesterday   []int64 `json:"yesterday" form:"yesterday" comment:"昨日流量"` // 昨日流量
}

type ServiceAddGrpcInput struct {
	ServiceName       string `json:"service_name" form:"service_name" comment:"服务名称" validate:"required,validServiceName"`
	ServiceDesc       string `json:"service_desc" form:"service_desc" comment:"服务描述" validate:"required"`
	Port              int    `json:"port" form:"port" comment:"端口，需要设置8001-8999范围内" validate:"required,min=8001,max=8999"`
	HeaderTransfor    string `json:"header_transfor" form:"header_transfor" comment:"metadata转换" validate:"validHeaderTransfor"`
	OpenAuth          int    `json:"open_auth" form:"open_auth" comment:"是否开启权限验证" validate:""`
	BlackList         string `json:"black_list" form:"black_list" comment:"黑名单IP，以逗号间隔，白名单优先级高于黑名单" validate:"validIpList"`
	WhiteList         string `json:"white_list" form:"white_list" comment:"白名单IP，以逗号间隔，白名单优先级高于黑名单" validate:"validIpList"`
	WhiteHostName     string `json:"white_host_name" form:"white_host_name" comment:"白名单主机，以逗号间隔" validate:"validIpList"`
	ClientIPFlowLimit int    `json:"clientip_flow_limit" form:"clientip_flow_limit" comment:"客户端IP限流" validate:""`
	ServiceFlowLimit  int    `json:"service_flow_limit" form:"service_flow_limit" comment:"服务端限流" validate:""`
	RoundType         int    `json:"round_type" form:"round_type" comment:"轮询策略" validate:""`
	IpList            string `json:"ip_list" form:"ip_list" comment:"IP列表" validate:"required,validIpPortList"`
	WeightList        string `json:"weight_list" form:"weight_list" comment:"权重列表" validate:"required,validWeightList"`
	ForbidList        string `json:"forbid_list" form:"forbid_list" comment:"禁用IP列表" validate:"validIpList"`
}

func (params *ServiceAddGrpcInput) GetValidParams(c *gin.Context) error {
	return public.DefaultGetValidParams(c, params)
}

type ServiceUpdateGrpcInput struct {
	ID                int64  `json:"id" form:"id" comment:"服务ID" validate:"required"`
	ServiceName       string `json:"service_name" form:"service_name" comment:"服务名称" validate:"required,validServiceName"`
	ServiceDesc       string `json:"service_desc" form:"service_desc" comment:"服务描述" validate:"required"`
	Port              int    `json:"port" form:"port" comment:"端口，需要设置8001-8999范围内" validate:"required,min=8001,max=8999"`
	HeaderTransfor    string `json:"header_transfor" form:"header_transfor" comment:"metadata转换" validate:"validHeaderTransfor"`
	OpenAuth          int    `json:"open_auth" form:"open_auth" comment:"是否开启权限验证" validate:""`
	BlackList         string `json:"black_list" form:"black_list" comment:"黑名单IP，以逗号间隔，白名单优先级高于黑名单" validate:"validIpList"`
	WhiteList         string `json:"white_list" form:"white_list" comment:"白名单IP，以逗号间隔，白名单优先级高于黑名单" validate:"validIpList"`
	WhiteHostName     string `json:"white_host_name" form:"white_host_name" comment:"白名单主机，以逗号间隔" validate:"validIpList"`
	ClientIPFlowLimit int    `json:"clientip_flow_limit" form:"clientip_flow_limit" comment:"客户端IP限流" validate:""`
	ServiceFlowLimit  int    `json:"service_flow_limit" form:"service_flow_limit" comment:"服务端限流" validate:""`
	RoundType         int    `json:"round_type" form:"round_type" comment:"轮询策略" validate:""`
	IpList            string `json:"ip_list" form:"ip_list" comment:"IP列表" validate:"required,validIpPortList"`
	WeightList        string `json:"weight_list" form:"weight_list" comment:"权重列表" validate:"required,validWeightList"`
	ForbidList        string `json:"forbid_list" form:"forbid_list" comment:"禁用IP列表" validate:"validIpList"`
}

func (params *ServiceUpdateGrpcInput) GetValidParams(c *gin.Context) error {
	return public.DefaultGetValidParams(c, params)
}

type ServiceAddTcpInput struct {
	ServiceName       string `json:"service_name" form:"service_name" comment:"服务名称" validate:"required,validServiceName"`
	ServiceDesc       string `json:"service_desc" form:"service_desc" comment:"服务描述" validate:"required"`
	Port              int    `json:"port" form:"port" comment:"端口，需要设置8001-8999范围内" validate:"required,min=8001,max=8999"`
	HeaderTransfor    string `json:"header_transfor" form:"header_transfor" comment:"header头转换" validate:""`
	OpenAuth          int    `json:"open_auth" form:"open_auth" comment:"是否开启权限验证" validate:""`
	BlackList         string `json:"black_list" form:"black_list" comment:"黑名单IP，以逗号间隔，白名单优先级高于黑名单" validate:"validIpList"`
	WhiteList         string `json:"white_list" form:"white_list" comment:"白名单IP，以逗号间隔，白名单优先级高于黑名单" validate:"validIpList"`
	WhiteHostName     string `json:"white_host_name" form:"white_host_name" comment:"白名单主机，以逗号间隔" validate:"validIpList"`
	ClientIPFlowLimit int    `json:"clientip_flow_limit" form:"clientip_flow_limit" comment:"客户端IP限流" validate:""`
	ServiceFlowLimit  int    `json:"service_flow_limit" form:"service_flow_limit" comment:"服务端限流" validate:""`
	RoundType         int    `json:"round_type" form:"round_type" comment:"轮询策略" validate:""`
	IpList            string `json:"ip_list" form:"ip_list" comment:"IP列表" validate:"required,validIpPortList"`
	WeightList        string `json:"weight_list" form:"weight_list" comment:"权重列表" validate:"required,validWeightList"`
	ForbidList        string `json:"forbid_list" form:"forbid_list" comment:"禁用IP列表" validate:"validIpList"`
}

func (params *ServiceAddTcpInput) GetValidParams(c *gin.Context) error {
	return public.DefaultGetValidParams(c, params)
}

type ServiceUpdateTcpInput struct {
	ID                int64  `json:"id" form:"id" comment:"服务ID" validate:"required"`
	ServiceName       string `json:"service_name" form:"service_name" comment:"服务名称" validate:"required,validServiceName"`
	ServiceDesc       string `json:"service_desc" form:"service_desc" comment:"服务描述" validate:"required"`
	Port              int    `json:"port" form:"port" comment:"端口，需要设置8001-8999范围内" validate:"required,min=8001,max=8999"`
	OpenAuth          int    `json:"open_auth" form:"open_auth" comment:"是否开启权限验证" validate:""`
	BlackList         string `json:"black_list" form:"black_list" comment:"黑名单IP，以逗号间隔，白名单优先级高于黑名单" validate:"validIpList"`
	WhiteList         string `json:"white_list" form:"white_list" comment:"白名单IP，以逗号间隔，白名单优先级高于黑名单" validate:"validIpList"`
	WhiteHostName     string `json:"white_host_name" form:"white_host_name" comment:"白名单主机，以逗号间隔" validate:"validIpList"`
	ClientIPFlowLimit int    `json:"clientip_flow_limit" form:"clientip_flow_limit" comment:"客户端IP限流" validate:""`
	ServiceFlowLimit  int    `json:"service_flow_limit" form:"service_flow_limit" comment:"服务端限流" validate:""`
	RoundType         int    `json:"round_type" form:"round_type" comment:"轮询策略" validate:""`
	IpList            string `json:"ip_list" form:"ip_list" comment:"IP列表" validate:"required,validIpPortList"`
	WeightList        string `json:"weight_list" form:"weight_list" comment:"权重列表" validate:"required,validWeightList"`
	ForbidList        string `json:"forbid_list" form:"forbid_list" comment:"禁用IP列表" validate:"validIpList"`
}

func (params *ServiceUpdateTcpInput) GetValidParams(c *gin.Context) error {
	return public.DefaultGetValidParams(c, params)
}
