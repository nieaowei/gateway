package dto

import (
	"gateway/dao"
	"gateway/lib"
	"gateway/public"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
)

type GetServiceListInput struct {
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

type GetServiceListOutput struct {
	Total int64             `json:"total" form:"total" validate:""`
	List  []ServiceListItem `json:"list" form:"list" validate:""`
}

func (p *GetServiceListInput) BindValidParam(c *gin.Context) (err error) {
	return public.DefaultGetValidParams(c, p)
}

func (p *GetServiceListInput) ErrorHandle(c *gin.Context, err error) {
	ResponseError(c, 1002, err)
}

func (p *GetServiceListInput) OutputHandle(c *gin.Context, outIn interface{}) (out interface{}) {
	return outIn
}

func (p *GetServiceListInput) Exec(c *gin.Context) (out interface{}, err error) {
	serviceInfo := &dao.ServiceInfo{}
	db := lib.GetDefaultDB()
	serviceInfos, count, err := serviceInfo.PageList(c, db, &dao.PageSize{
		Size: p.PageSize,
		No:   p.PageNo,
		Info: p.Info,
	})
	if err != nil {
		return
	}
	outE := &GetServiceListOutput{
		Total: count,
		List:  []ServiceListItem{},
	}
	conf := lib.GetDefaultConfBase()
	clusterIP := conf.Cluster.Ip
	clusterPort := conf.Cluster.Port
	clusterSSLPort := conf.Cluster.SslPort
	for _, info := range serviceInfos {
		serviceDetail, err := info.FindOneServiceDetail(c, db)
		if err != nil {
			return nil, err
		}
		serviceAddr := ""
		loadType := ""
		switch serviceDetail.LoadType {
		case dao.LoadTypeHttp:
			{
				loadType = "HTTP"
				service := serviceDetail.ServiceHTTPRuleExceptModel
				if service.RuleType == dao.HttpRuleTypePrefixURL && service.NeedHTTPs == 0 {
					serviceAddr = clusterIP + ":" + clusterPort + service.Rule

				}
				if service.RuleType == dao.HttpRuleTypePrefixURL && service.NeedHTTPs == 1 {
					serviceAddr = clusterIP + ":" + clusterSSLPort + service.Rule
				}
				if service.RuleType == dao.HttpRuleTypeDomain {
					serviceAddr = service.Rule
				}
				break
			}
		case dao.LoadTypeTcp:
			{
				loadType = "TCP"
				service := serviceDetail.ServiceTCPRuleExceptModel

				serviceAddr = clusterIP + ":" + strconv.Itoa(int(service.Port))
				break
			}
		case dao.LoadTypeGrpc:
			{
				loadType = "GRPC"
				service := serviceDetail.ServiceGrpcRuleExceptModel

				serviceAddr = clusterIP + ":" + strconv.Itoa(int(service.Port))
				break
			}
		}
		ipList := serviceDetail.ServiceLoadBalanceExceptModel.GetIPListByModel()

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

type DeleteServiceInput struct {
	ID uint `json:"id" form:"id" validate:"required"`
}

func (p *DeleteServiceInput) BindValidParam(c *gin.Context) (err error) {
	return public.DefaultGetValidParams(c, p)
}

func (p *DeleteServiceInput) ErrorHandle(c *gin.Context, err error) {
	ResponseError(c, 1002, err)
}

func (p *DeleteServiceInput) OutputHandle(c *gin.Context, outIn interface{}) (out interface{}) {
	return outIn
}

func (p *DeleteServiceInput) Exec(c *gin.Context) (out interface{}, err error) {
	db := lib.GetDefaultDB()
	serviceInfo := &dao.ServiceInfo{
		ID: p.ID,
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

type GetServiceDetailInput struct {
	ServiceID uint `json:"service_id" form:"service_id"`
}

type GetServiceDetailOutput struct {
	dao.ServiceDetail
}

func (p *GetServiceDetailInput) BindValidParam(c *gin.Context) (err error) {
	err = public.DefaultGetValidParams(c, p)
	return
}

func (p *GetServiceDetailInput) ErrorHandle(c *gin.Context, err error) {
	ResponseError(c, 1002, err)
}

func (p *GetServiceDetailInput) OutputHandle(c *gin.Context, outIn interface{}) (out interface{}) {
	return outIn
}

func (p *GetServiceDetailInput) Exec(c *gin.Context) (out interface{}, err error) {
	db := lib.GetDefaultDB()
	info := &dao.ServiceInfo{
		ID: p.ServiceID,
	}
	out, err = info.FindOneServiceDetail(c, db)
	if err != nil {
		return
	}
	return
}
