package dao

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

/**
 * @Author: liu zw
 * @Date: 2022/10/18 12:21
 * @File: service_grpc_rule.go
 * @Description: //TODO $
 * @Version:
 */

type GrpcRule struct {
	ID             int64  `json:"id" gorm:"primary_key"`
	ServiceID      int64  `json:"service_id" gorm:"column:service_id" description:"服务id	"`
	Port           int    `json:"port" gorm:"column:port" description:"端口	"`
	HeaderTransfor string `json:"header_transfor" gorm:"column:header_transfor" description:"header转换支持增加(add)、删除(del)、修改(edit) 格式: add headname headvalue"`
}

func (gr *GrpcRule) TableName() string {
	return "gateway_service_grpc_rule"
}

func (gr *GrpcRule) Find(c *gin.Context, tx *gorm.DB, search *GrpcRule) (*GrpcRule, error) {
	out := &GrpcRule{}
	err := tx.WithContext(c).Where(search).Find(out).Error
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (gr *GrpcRule) Save(c *gin.Context, tx *gorm.DB) error {
	return tx.WithContext(c).Save(gr).Error
}
