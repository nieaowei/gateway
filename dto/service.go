package dto

import (
	"gateway/dao"
	"gateway/lib"
	"gateway/public"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
)

type ServiceListInput struct {
	Info     string `json:"info" form:"info"`
	PageNo   int    `json:"page_no" form:"page_no" validate:"required"`
	PageSize int    `json:"page_size" form:"page_size" validate:"required"`
}

type ServiceListItem struct {
	ID          uint   `json:"id" form:"id"`
	ServiceName string `json:"service_name" form:"service_name"`
	ServiceDesc string `json:"service_desc" form:"service_desc"`
	LoadType    string `json:"load_type" form:"load_type"`
	Address     string `json:"address" form:"Address"`
	Qps         uint   `json:"qps" form:"qps"`
	Qpd         uint   `json:"qpd" form:"qpd"`
	TotalNode   uint   `json:"total_node" form:"total_node"`
}

type ServiceListOutput struct {
	Total int64             `json:"total" form:"total" validate:""`
	List  []ServiceListItem `json:"list" form:"list" validate:""`
}

func (p *ServiceListInput) Exec(c *gin.Context) (out interface{}, err error) {
	serviceInfo := &dao.ServiceInfo{}
	db, err := lib.GetDefaultDB()
	if err != nil {
		return
	}
	serviceInfos, count, err := serviceInfo.PageList(c, db, &dao.PageSize{
		Size: p.PageSize,
		No:   p.PageNo,
		Info: p.Info,
	})
	if err != nil {
		return
	}
	outE := &ServiceListOutput{
		Total: count,
		List:  []ServiceListItem{},
	}
	clusterIP := lib.GetDefaultConfBase().Cluster.Ip
	clusterPort := lib.GetDefaultConfBase().Cluster.Port
	clusterSSLPort := lib.GetDefaultConfBase().Cluster.SslPort
	for _, info := range serviceInfos {
		serviceDetail, err := info.FindOneServiceDetail(c, db)
		if err != nil {
			return nil, err
		}
		serviceAddr := ""
		loadType := ""
		switch serviceDetail.Info.LoadType {
		case dao.LoadTypeHttp:
			{
				loadType = "HTTP"
				if serviceDetail.HTTP.RuleType == dao.HttpRuleTypePrefixURL && serviceDetail.HTTP.NeedHTTPs == 0 {
					serviceAddr = clusterIP + ":" + clusterPort + serviceDetail.HTTP.Rule

				}
				if serviceDetail.HTTP.RuleType == dao.HttpRuleTypePrefixURL && serviceDetail.HTTP.NeedHTTPs == 1 {
					serviceAddr = clusterIP + ":" + clusterSSLPort + serviceDetail.HTTP.Rule
				}
				if serviceDetail.HTTP.RuleType == dao.HttpRuleTypeDomain {
					serviceAddr = serviceDetail.HTTP.Rule
				}
				break
			}
		case dao.LoadTypeTcp:
			{
				loadType = "TCP"
				serviceAddr = clusterIP + ":" + strconv.Itoa(int(serviceDetail.TCP.Port))
				break
			}
		case dao.LoadTypeGrpc:
			{
				loadType = "GRPC"
				serviceAddr = clusterIP + ":" + strconv.Itoa(int(serviceDetail.GRPC.Port))
				break
			}
		}
		ipList := serviceDetail.LoadBalance.GetIPListByModel()

		item := ServiceListItem{
			ID:          info.ID,
			ServiceName: info.ServiceName,
			ServiceDesc: info.ServiceDesc,
			LoadType:    loadType,
			Address:     serviceAddr,
			Qps:         0,
			Qpd:         0,
			TotalNode:   uint(len(ipList)),
		}
		outE.List = append(outE.List, item)
	}
	return outE, nil
}

func (p *ServiceListInput) BindValidParam(c *gin.Context) (err error) {
	return public.DefaultGetValidParams(c, p)
}

func (p *ServiceListInput) ErrorHandle(c *gin.Context, err error) {
	ResponseError(c, 1002, err)
}

type DeleteServiceInput struct {
	ID uint `json:"id" form:"id" validate:"required"`
}

func (p *DeleteServiceInput) BindValidParam(c *gin.Context) (err error) {
	return public.DefaultGetValidParams(c, p)
}

func (p *DeleteServiceInput) ErrorHandle(c *gin.Context, err error) {
	ResponseError(c, 1002, err)
}

func (p *DeleteServiceInput) Exec(c *gin.Context) (out interface{}, err error) {
	db, err := lib.GetDefaultDB()
	if err != nil {
		return
	}
	serviceInfo := &dao.ServiceInfo{
		Model: gorm.Model{
			ID: p.ID,
		},
	}

	err = db.Transaction(
		func(tx *gorm.DB) (err error) {
			serviceInfo, err = serviceInfo.FindOne(c, tx)
			if err != nil {
				return
			}

			switch serviceInfo.LoadType {
			case dao.LoadTypeHttp:
				{
					http := dao.ServiceHTTPRule{
						ServiceID: serviceInfo.ID,
					}
					err = http.DeleteByID(c, tx)
					break
				}
			case dao.LoadTypeGrpc:
				{
					grpc := dao.ServiceGrpcRule{
						ServiceID: serviceInfo.ID,
					}
					err = grpc.DeleteByID(c, tx)
					break
				}
			case dao.LoadTypeTcp:
				{
					tcp := dao.ServiceTCPRule{
						ServiceID: serviceInfo.ID,
					}
					err = tcp.DeleteByID(c, tx)
					break
				}
			}
			if err != nil {
				return
			}
			slb := &dao.ServiceLoadBalance{
				ServiceID: serviceInfo.ID,
			}
			err = slb.DeleteByID(c, tx)
			if err != nil {
				return
			}

			sac := &dao.ServiceAccessControl{
				ServiceID: serviceInfo.ID,
			}
			err = sac.DeleteByID(c, tx)
			if err != nil {
				return
			}

			err = serviceInfo.DeleteByID(c, tx)
			if err != nil {
				return
			}
			return
		})
	return
}
