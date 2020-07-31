package dao

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

//type ServiceInfo struct {
//	gorm.Model
//	LoadType    uint   `json:"load_type" validate:"oneof=0 1 2"`
//	ServiceName string `json:"service_name" validate:"required,alphanum,max=255,min=6"`
//	ServiceDesc string `json:"service_desc" validate:"required,max=255,min=1"`
//}

func (p *ServiceInfo) ServiceDetail() string {
	return "service_info"
}

func (p *ServiceInfo) FindOne(c *gin.Context, tx *gorm.DB) (out *ServiceInfo, err error) {
	out = &ServiceInfo{}
	err = tx.Where(p).First(out).Error
	if err != nil {
		return nil, err
	}
	return
}

func (p *ServiceInfo) Delete(c *gin.Context, tx *gorm.DB) (err error) {
	return tx.Where(p).Delete(p).Error
}

func (p *ServiceInfo) PageList(c *gin.Context, tx *gorm.DB, params *PageSize) (list []ServiceInfo, count int64, err error) {
	offset := (params.No - 1) * params.Size
	query := tx.Model(p)
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

func (p *ServiceInfo) DeleteOneIncludeChild(c *gin.Context, db *gorm.DB) (err error) {
	err = db.Transaction(func(tx *gorm.DB) (err error) {
		p, err = p.FindOne(c, tx)
		if err != nil {
			return
		}

		switch p.LoadType {
		case LoadTypeHttp:
			{
				http := ServiceHTTPRule{
					ServiceID: p.ID,
				}
				err = http.Delete(c, tx)
				break
			}
		case LoadTypeGrpc:
			{
				grpc := ServiceGrpcRule{
					ServiceID: p.ID,
				}
				err = grpc.Delete(c, tx)
				break
			}
		case LoadTypeTcp:
			{
				tcp := ServiceTCPRule{
					ServiceID: p.ID,
				}
				err = tcp.Delete(c, tx)
				break
			}
		}
		if err != nil {
			return
		}
		slb := &ServiceLoadBalance{
			ServiceID: p.ID,
		}
		err = slb.Delete(c, tx)
		if err != nil {
			return
		}

		sac := &ServiceAccessControl{
			ServiceID: p.ID,
		}
		err = sac.Delete(c, tx)
		if err != nil {
			return
		}

		err = p.Delete(c, tx)
		if err != nil {
			return
		}
		return
	})

	return
}

type ServiceDetail struct {
	Info          *ServiceInfo          `json:"info"`
	HTTP          *ServiceHTTPRule      `json:"http"`
	GRPC          *ServiceGrpcRule      `json:"grpc"`
	TCP           *ServiceTCPRule       `json:"tcp"`
	LoadBalance   *ServiceLoadBalance   `json:"load_balance"`
	AccessControl *ServiceAccessControl `json:"access_control"`
}

func (p *ServiceInfo) FindOneServiceDetail(c *gin.Context, db *gorm.DB) (out *ServiceDetail, err error) {
	out = &ServiceDetail{}
	switch p.LoadType {
	case LoadTypeHttp:
		{
			httpRule := &ServiceHTTPRule{
				ServiceID: p.ID,
			}
			httpRule, err = httpRule.FindOne(c, db)
			if err != nil {
				return
			}
			out.HTTP = httpRule
			break
		}
	case LoadTypeTcp:
		{
			tcpRule := &ServiceTCPRule{
				ServiceID: p.ID,
			}
			tcpRule, err = tcpRule.FindOne(c, db)
			if err != nil {
				return
			}
			out.TCP = tcpRule
			break
		}
	case LoadTypeGrpc:
		{
			grpcRule := &ServiceGrpcRule{
				ServiceID: p.ID,
			}
			grpcRule, err = grpcRule.FindOne(c, db)
			if err != nil {
				return
			}
			out.GRPC = grpcRule
			break
		}
	}
	accessControl := &ServiceAccessControl{
		ServiceID: p.ID,
	}
	accessControl, err = accessControl.FindOne(c, db)
	if err != nil {
		return
	}

	loadBalance := &ServiceLoadBalance{
		ServiceID: p.ID,
	}
	loadBalance, err = loadBalance.FindOne(c, db)
	if err != nil {
		return
	}

	out.AccessControl = accessControl
	out.LoadBalance = loadBalance
	out.Info = p

	return
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
