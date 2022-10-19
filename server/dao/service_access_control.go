package dao

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

/**
 * @Author: liu zw
 * @Date: 2022/10/18 12:20
 * @File: service_access_control.go
 * @Description: //TODO $
 * @Version:
 */

type AccessControl struct {
	ID                int64  `json:"id" gorm:"primary_key"`
	ServiceID         int64  `json:"service_id" gorm:"column:service_id" description:"服务id"`
	OpenAuth          int    `json:"open_auth" gorm:"column:open_auth" description:"是否开启权限 1=开启"`
	BlackList         string `json:"black_list" gorm:"column:black_list" description:"黑名单ip	"`
	WhiteList         string `json:"white_list" gorm:"column:white_list" description:"白名单ip	"`
	WhiteHostName     string `json:"white_host_name" gorm:"column:white_host_name" description:"白名单主机	"`
	ClientIPFlowLimit int    `json:"clientip_flow_limit" gorm:"column:clientip_flow_limit" description:"客户端ip限流	"`
	ServiceFlowLimit  int    `json:"service_flow_limit" gorm:"column:service_flow_limit" description:"服务端限流	"`
}

func (ac *AccessControl) TableName() string {
	return "gateway_service_access_control"
}

func (ac *AccessControl) Find(c *gin.Context, tx *gorm.DB, search *AccessControl) (*AccessControl, error) {
	out := &AccessControl{}
	err := tx.WithContext(c).Where(search).Find(out).Error
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (ac *AccessControl) Save(c *gin.Context, tx *gorm.DB) error {
	return tx.WithContext(c).Save(ac).Error
}
