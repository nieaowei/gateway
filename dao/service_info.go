package dao

import (
	"gateway/public"
	"github.com/e421083458/gorm"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

type ServiceInfo struct {
	gorm.Model
	LoadType    uint   `json:"load_type" validate:"oneof=0 1 2"`
	ServiceName string `json:"service_name" validate:"required,alphanum,max=255,min=6"`
	ServiceDesc string `json:"service_desc" validate:"required,max=255,min=1"`
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

func (p *ServiceInfo) DeleteOneIncludeChild(c *gin.Context, tx *gorm.DB) (err error) {
	tx = tx.Begin()
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
	if err != nil {
		tx.Rollback()
		return
	}
	err = tx.Delete(p).Error
	if err != nil {
		tx.Rollback()
		return
	}
	tx.Commit()
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
	//switch p.LoadType {
	//case LoadTypeHttp:
	//	{
	//		break
	//	}
	//case LoadTypeTcp:
	//	{
	//		break
	//	}
	//case LoadTypeGrpc:
	//	{
	//		break
	//	}
	//}

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
	return out, nil
}

func (p *ServiceInfo) AddAfterCheck(c *gin.Context, db *gorm.DB) error {
	serviceInfo := &ServiceInfo{
		ServiceName: p.ServiceName,
	}
	// check unique ServiceName
	err := db.First(serviceInfo, serviceInfo).Error
	if err != gorm.ErrRecordNotFound {
		return errors.New("Violation of the uniqueness constraint #ServiceInfo.ServiceName")
	}
	// make sure insert
	p.ID = 0
	err = db.Create(p).Error
	return err
}
