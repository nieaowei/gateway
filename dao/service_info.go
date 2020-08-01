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
	result := tx.Where(p).First(out)
	err = ErrorHandleForDB(result)
	return
}

func (p *ServiceInfo) BeforeUpdate(tx *gorm.DB) error {
	tx = tx.Statement.Where("deleted_at IS NULL").Omit("created_at")
	return nil
}

func (p *ServiceInfo) BeforeDelete(tx *gorm.DB) error {
	tx = tx.Statement.Where("deleted_at IS NULL")
	return nil
}
func (p *ServiceInfo) BeforeCreate(db *gorm.DB) error {
	db.Statement.Omit("id")
	return nil
}

func (p *ServiceInfo) UpdateAll(c *gin.Context, db *gorm.DB) (err error) {
	return db.Save(p).Error
}

func (p *ServiceInfo) DeleteByID(c *gin.Context, tx *gorm.DB) (err error) {
	result := tx.Delete(p)
	err = ErrorHandleForDB(result)
	return
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
	err = db.Transaction(
		func(tx *gorm.DB) (err error) {
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
					err = http.DeleteByID(c, tx)
					break
				}
			case LoadTypeGrpc:
				{
					grpc := ServiceGrpcRule{
						ServiceID: p.ID,
					}
					err = grpc.DeleteByID(c, tx)
					break
				}
			case LoadTypeTcp:
				{
					tcp := ServiceTCPRule{
						ServiceID: p.ID,
					}
					err = tcp.DeleteByID(c, tx)
					break
				}
			}
			if err != nil {
				return
			}
			slb := &ServiceLoadBalance{
				ServiceID: p.ID,
			}
			err = slb.DeleteByID(c, tx)
			if err != nil {
				return
			}

			sac := &ServiceAccessControl{
				ServiceID: p.ID,
			}
			err = sac.DeleteByID(c, tx)
			if err != nil {
				return
			}

			err = p.DeleteByID(c, tx)
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

func (p *ServiceInfo) InsertAfterCheck(c *gin.Context, db *gorm.DB, check bool) (err error) {
	if check {
		serviceInfo := &ServiceInfo{
			ServiceName: p.ServiceName,
		}
		// check unique ServiceName
		err = db.First(serviceInfo, serviceInfo).Error
		if err != gorm.ErrRecordNotFound {
			return errors.New("Violation of the uniqueness constraint #ServiceInfo.ServiceName")
		}
	}
	// make sure insert
	res := db.Create(p)
	err = ErrorHandleForDB(res)
	return
}
