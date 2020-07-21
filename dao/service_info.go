package dao

import (
	"gateway/public"
	"github.com/e421083458/gorm"
	"github.com/gin-gonic/gin"
)

type ServiceInfo struct {
	gorm.Model
	LoadType    uint   `json:"load_type"`
	ServiceName string `json:"service_name"`
	ServiceDesc string `json:"service_desc"`
}

func (p *ServiceInfo) TableName() string {
	return "service_info"
}

func (p *ServiceInfo) ServiceDetail() string {
	return "service_info"
}

func (p *ServiceInfo) PageList(c *gin.Context, tx *gorm.DB, params *PageSize) (list []ServiceInfo, count uint, err error) {
	offset := (params.No - 1) * params.Size
	query := tx.SetCtx(public.GetTraceContext(c)).Model(p)
	if params.Info != "" {
		query = query.Where("service_name like ? or service_desc like ?", "%"+params.Info+"%", "%"+params.Info+"%")
	}

	err = query.Limit(params.Size).Offset(offset).Order("id desc").Find(&list).Error

	if err != nil {
		return
	}
	query.Count(&count)
	return
}

func (p *ServiceInfo) DeleteOne(c *gin.Context, tx *gorm.DB) (err error) {
	serviceDeatail, err := p.FindOneServiceDetail(c, tx)
	if err != nil {
		return
	}
	switch p.LoadType {
	case LoadTypeHttp:
		{
			err = serviceDeatail.HTTP.Delete(c, tx)
			break
		}
	case LoadTypeGrpc:
		{
			err = serviceDeatail.GRPC.Delete(c, tx)
			break
		}
	case LoadTypeTcp:
		{
			err = serviceDeatail.TCP.Delete(c, tx)
			break
		}
	}
	return
}

type ServiceDetail struct {
	Info          *ServiceInfo          `json:"info"`
	HTTP          *ServiceHttpRule      `json:"http"`
	GRPC          *ServiceGrpcRule      `json:"grpc"`
	TCP           *ServiceTcpRule       `json:"tcp"`
	LoadBalance   *ServiceLoadBalance   `json:"load_balance"`
	AccessControl *ServiceAccessControl `json:"access_control"`
}

func (p *ServiceInfo) FindOneServiceDetail(c *gin.Context, db *gorm.DB) (out *ServiceDetail, err error) {

	//todo wait next step optimization.
	switch p.LoadType {
	case LoadTypeHttp:
		{
			break
		}
	case LoadTypeTcp:
		{
			break
		}
	case LoadTypeGrpc:
		{
			break
		}
	}

	httpRule := &ServiceHttpRule{
		ServiceId: p.ID,
	}
	httpRule, err = httpRule.FindOne(c, db)
	if err != nil && err != gorm.ErrRecordNotFound {
		return
	}

	tcpRule := &ServiceTcpRule{
		ServiceId: p.ID,
	}
	tcpRule, err = tcpRule.FindOne(c, db)
	if err != nil && err != gorm.ErrRecordNotFound {
		return
	}

	grpcRule := &ServiceGrpcRule{
		ServiceId: p.ID,
	}
	grpcRule, err = grpcRule.FindOne(c, db)
	if err != nil && err != gorm.ErrRecordNotFound {
		return
	}

	accessControl := &ServiceAccessControl{
		ServiceId: p.ID,
	}
	accessControl, err = accessControl.FindOne(c, db)
	if err != nil && err != gorm.ErrRecordNotFound {
		return
	}

	loadBalance := &ServiceLoadBalance{
		ServiceId: p.ID,
	}
	loadBalance, err = loadBalance.FindOne(c, db)
	if err != nil && err != gorm.ErrRecordNotFound {
		return
	}

	out = &ServiceDetail{
		Info:          p,
		HTTP:          httpRule,
		GRPC:          grpcRule,
		TCP:           tcpRule,
		LoadBalance:   loadBalance,
		AccessControl: accessControl,
	}
	return
}
