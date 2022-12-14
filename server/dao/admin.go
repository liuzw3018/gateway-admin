package dao

import (
	"github.com/gin-gonic/gin"
	"github.com/liuzw3018/gateway_admin/dto"
	"github.com/liuzw3018/gateway_admin/public"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"time"
)

/**
 * @Author: liu zw
 * @Date: 2022/10/17 14:25
 * @File: admin.go
 * @Description: //TODO $
 * @Version:
 */

type Admin struct {
	Id        int       `json:"id" gorm:"primary_key" description:"自增主键"`
	Username  string    `json:"Username" gorm:"column:user_name" description:"管理员用户名"`
	Salt      string    `json:"Salt" gorm:"column:salt" description:"盐"`
	Password  string    `json:"Password" gorm:"column:password" description:"密码"`
	UpdatedAt time.Time `json:"update_at" gorm:"column:update_at" description:"更新时间"`
	CreatedAt time.Time `json:"create_at" gorm:"column:create_at" description:"创建时间"`
	IsDelete  int       `json:"is_delete" gorm:"column:is_delete" description:"是否删除"`
}

func (a *Admin) TableName() string {
	return "gateway_admin"
}

func (a *Admin) Find(c *gin.Context, tx *gorm.DB, search *Admin) (*Admin, error) {
	out := &Admin{}
	err := tx.WithContext(c).Where(search).Find(out).Error
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (a *Admin) LoginCheck(c *gin.Context, tx *gorm.DB, param *dto.AdminLoginInput) (*Admin, error) {
	adminInfo, err := a.Find(c, tx,
		&Admin{Username: param.Username, IsDelete: 0})
	if err != nil {
		return nil, errors.New("用户信息不存在")
	}
	saltPassword := public.GenSaltPassword(adminInfo.Salt, param.Password)
	if adminInfo.Password != saltPassword {
		return nil, errors.New("密码错误，请重新输入")
	}
	return adminInfo, nil
}

func (a *Admin) Save(c *gin.Context, tx *gorm.DB) error {
	return tx.WithContext(c).Save(a).Error
}
