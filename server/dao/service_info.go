package dao

import (
	"github.com/gin-gonic/gin"
	"github.com/liuzw3018/gateway_admin/dto"
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

type ServiceInfo struct {
	ID          int64     `json:"id" gorm:"primary_key" description:"自增主键"`
	LoadType    int       `json:"load_type" gorm:"column:load_type" description:"负载类型 0=http 1=tcp 2=grpc"`
	ServiceName string    `json:"service_name" gorm:"column:service_name" description:"服务名称"`
	ServiceDesc string    `json:"service_desc" gorm:"column:service_desc" description:"服务描述"`
	UpdatedAt   time.Time `json:"update_at" gorm:"column:update_at" description:"更新时间"`
	CreatedAt   time.Time `json:"create_at" gorm:"column:create_at" description:"创建时间"`
	IsDelete    int       `json:"is_delete" gorm:"column:is_delete" description:"是否删除"`
}

func (s *ServiceInfo) TableName() string {
	return "gateway_service_info"
}

func (s *ServiceInfo) Find(c *gin.Context, tx *gorm.DB, search *ServiceInfo) (*ServiceInfo, error) {
	out := &ServiceInfo{}
	err := tx.WithContext(c).Where(search).Find(out).Error
	if err != nil {
		return nil, err
	}
	if out.ServiceName == "" {
		return nil, gorm.ErrRecordNotFound
	}
	return out, nil
}

func (s *ServiceInfo) Save(c *gin.Context, tx *gorm.DB) error {
	return tx.WithContext(c).Save(s).Error
}

func (s *ServiceInfo) PageList(c *gin.Context, tx *gorm.DB, param *dto.ServiceListInput) ([]ServiceInfo, int64, error) {
	var (
		total  int64
		list   []ServiceInfo
		offset = (param.PageNo - 1) * param.PageSize
	)
	query := tx.WithContext(c)
	query = query.Table(s.TableName()).Where("is_delete=0")
	if param.Info != "" {
		query = query.Where("service_name like ? or service_desc like ?", "%"+param.Info+"%", "%"+param.Info+"%")
	}
	if err := query.Limit(param.PageSize).
		Offset(offset).
		Order("id desc").
		Find(&list).
		Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0, err
	}

	query.Limit(param.PageSize).Offset(offset).Count(&total)
	return list, total, nil
}

func (s *ServiceInfo) Detail(c *gin.Context, tx *gorm.DB, search *ServiceInfo) (*ServiceDetail, error) {
	if search.ServiceName == "" {
		info, err := s.Find(c, tx, search)
		if err != nil {
			return nil, err
		}
		search = info
	}
	httpRule := &HttpRule{ServiceID: search.ID}
	httpRule, err := httpRule.Find(c, tx, httpRule)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	tcpRule := &TcpRule{ServiceID: search.ID}
	tcpRule, err = tcpRule.Find(c, tx, tcpRule)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	grpcRule := &GrpcRule{ServiceID: search.ID}
	grpcRule, err = grpcRule.Find(c, tx, grpcRule)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	loadBalance := &LoadBalance{ServiceID: search.ID}
	loadBalance, err = loadBalance.Find(c, tx, loadBalance)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	accessControl := &AccessControl{ServiceID: search.ID}
	accessControl, err = accessControl.Find(c, tx, accessControl)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	detail := &ServiceDetail{
		Info:          search,
		HttpRule:      httpRule,
		TCPRule:       tcpRule,
		GRPCRule:      grpcRule,
		LoadBalance:   loadBalance,
		AccessControl: accessControl,
	}
	return detail, nil
}

func (s *ServiceInfo) GroupByLoadType(c *gin.Context, tx *gorm.DB) ([]dto.DashboardServiceStatItemOutput, error) {
	var (
		list []dto.DashboardServiceStatItemOutput
	)
	query := tx.WithContext(c)
	if err := query.Table(s.TableName()).Where("is_delete=0").Select("load_type,count(*) as value").
		Group("load_type").Scan(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}
