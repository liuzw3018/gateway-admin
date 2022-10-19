package dao

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

/**
 * @Author: liu zw
 * @Date: 2022/10/18 12:23
 * @File: service_tcp_rule.go
 * @Description: //TODO $
 * @Version:
 */

type TcpRule struct {
	ID        int64 `json:"id" gorm:"primary_key"`
	ServiceID int64 `json:"service_id" gorm:"column:service_id" description:"服务id	"`
	Port      int   `json:"port" gorm:"column:port" description:"端口	"`
}

func (tr *TcpRule) TableName() string {
	return "gateway_service_tcp_rule"
}

func (tr *TcpRule) Find(c *gin.Context, tx *gorm.DB, search *TcpRule) (*TcpRule, error) {
	out := &TcpRule{}
	err := tx.WithContext(c).Where(search).Find(out).Error
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (tr *TcpRule) Save(c *gin.Context, tx *gorm.DB) error {
	return tx.WithContext(c).Save(tr).Error
}
