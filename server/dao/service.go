package dao

/**
 * @Author: liu zw
 * @Date: 2022/10/18 12:18
 * @File: service.go
 * @Description: //TODO $
 * @Version:
 */

type ServiceDetail struct {
	Info          *ServiceInfo   `json:"info" description:"基本信息"`
	HttpRule      *HttpRule      `json:"http" description:"http_rule"`
	TCPRule       *TcpRule       `json:"tcp" description:"tcp_rule"`
	GRPCRule      *GrpcRule      `json:"grpc" description:"grpc_rule"`
	LoadBalance   *LoadBalance   `json:"load_balance" description:"load_balance_rule"`
	AccessControl *AccessControl `json:"access_control" description:"access_control_rule"`
}
