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

func (p *ServiceInfo) FindOneScan(c *gin.Context, tx *gorm.DB, out interface{}) (err error) {
	//out = &ServiceInfo{}
	result := tx.Model(p).Where(p).Limit(1).Scan(out)
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

func (p *ServiceInfo) UpdateAllByID(c *gin.Context, db *gorm.DB) (err error) {
	result := db.Select(GetFields(p)).Where("id=?", p.ID).Updates(p)
	err = ErrorHandleForDB(result)
	return
}

func (p *ServiceInfo) DeleteByID(c *gin.Context, tx *gorm.DB) (err error) {
	result := tx.Where(p).Delete(p)
	err = ErrorHandleForDB(result)
	return
}

func (p *ServiceInfo) PageList(c *gin.Context, tx *gorm.DB, params *PageSize) (list []ServiceInfo, count int64, err error) {
	offset := (params.No - 1) * params.Size
	query := tx.Model(p)
	if params.Info != "" {
		query = query.Where("service_name like ? or service_desc like ?", "%"+params.Info+"%", "%"+params.Info+"%")
	}
	query.Count(&count)

	err = query.Limit(params.Size).Offset(offset).Order("id desc").Find(&list).Error

	if err != nil {
		return
	}
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
			case LoadType_HTTP:
				{
					http := ServiceHTTPRule{
						ServiceID: p.ID,
					}
					err = http.DeleteByID(c, tx)
					break
				}
			case LoadType_GRPC:
				{
					grpc := ServiceGrpcRule{
						ServiceID: p.ID,
					}
					err = grpc.DeleteByID(c, tx)
					break
				}
			case LoadType_TCP:
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
	*ServiceInfoExceptModel
	*ServiceHTTPRuleExceptModel
	*ServiceGrpcRuleExceptModel
	*ServiceTCPRuleExceptModel
	*ServiceLoadBalanceExceptModel
	*ServiceAccessControlExceptModel
}

func (p *ServiceInfo) FindOneServiceDetail(c *gin.Context, db *gorm.DB) (out *ServiceDetail, err error) {

	out = &ServiceDetail{}
	out.ServiceInfoExceptModel = &ServiceInfoExceptModel{}
	err = p.FindOneScan(c, db, out.ServiceInfoExceptModel)
	if err != nil {
		return
	}
	switch out.ServiceInfoExceptModel.LoadType {
	case LoadType_HTTP:
		{
			httpRule := &ServiceHTTPRule{
				ServiceID: p.ID,
			}
			httpOut := &ServiceHTTPRuleExceptModel{}
			err = httpRule.FindOneScan(c, db, httpOut)
			if err != nil {
				return
			}
			out.ServiceHTTPRuleExceptModel = httpOut
			break
		}
	case LoadType_TCP:
		{
			tcpRule := &ServiceTCPRule{
				ServiceID: p.ID,
			}
			tcpOut := &ServiceTCPRuleExceptModel{}
			err = tcpRule.FindOneScan(c, db, tcpOut)
			if err != nil {
				return
			}
			out.ServiceTCPRuleExceptModel = tcpOut

			break
		}
	case LoadType_GRPC:
		{
			grpcRule := &ServiceGrpcRule{
				ServiceID: p.ID,
			}
			grpcOut := &ServiceGrpcRuleExceptModel{}

			err = grpcRule.FindOneScan(c, db, grpcOut)
			if err != nil {
				return
			}
			out.ServiceGrpcRuleExceptModel = grpcOut

			break
		}
	}
	accessControl := &ServiceAccessControl{
		ServiceID: p.ID,
	}
	out.ServiceAccessControlExceptModel = &ServiceAccessControlExceptModel{}
	err = accessControl.FindOneScan(c, db, out.ServiceAccessControlExceptModel)
	if err != nil {
		return
	}

	loadBalance := &ServiceLoadBalance{
		ServiceID: p.ID,
	}
	out.ServiceLoadBalanceExceptModel = &ServiceLoadBalanceExceptModel{}

	err = loadBalance.FindOneScan(c, db, out.ServiceLoadBalanceExceptModel)
	if err != nil {
		return
	}
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
