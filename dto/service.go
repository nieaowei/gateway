package dto

import (
	"gateway/dao"
	"gateway/public"
	"github.com/e421083458/golang_common/lib"
	"github.com/e421083458/gorm"
	"github.com/gin-gonic/gin"
	"strconv"
)

type ServiceListInput struct {
	Info     string `json:"info" form:"info"`
	PageNo   uint   `json:"page_no" form:"page_no" validate:"required"`
	PageSize uint   `json:"page_size" form:"page_size" validate:"required"`
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
	Total uint              `json:"total" form:"total" validate:""`
	List  []ServiceListItem `json:"list" form:"list" validate:""`
}

func (p *ServiceListInput) GetServiceList(c *gin.Context) (out *ServiceListOutput, err error) {
	serviceInfo := &dao.ServiceInfo{}
	db, err := lib.GetGormPool("default")
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
	out = &ServiceListOutput{
		Total: count,
		List:  []ServiceListItem{},
	}
	clusterIP := lib.GetStringConf("base.cluster.cluster_ip")
	clusterPort := lib.GetStringConf("base.cluster.cluster_port")
	clusterSSLPort := lib.GetStringConf("base.cluster.cluster_ssl_port")
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
				if serviceDetail.HTTP.RuleType == dao.HttpRuleTypePrefixURL && serviceDetail.HTTP.NeedHttps == 0 {
					serviceAddr = clusterIP + ":" + clusterPort + serviceDetail.HTTP.Rule

				}
				if serviceDetail.HTTP.RuleType == dao.HttpRuleTypePrefixURL && serviceDetail.HTTP.NeedHttps == 1 {
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
		out.List = append(out.List, item)
	}

	return
}

func (p *ServiceListInput) BindValidParam(c *gin.Context) (err error) {
	return public.DefaultGetValidParams(c, p)
}

type ServiceDeleteInput struct {
	ID uint `json:"id" form:"id" validate:"required"`
}

func (p *ServiceDeleteInput) Delete(c *gin.Context) (err error) {
	db, err := lib.GetGormPool("default")
	if err != nil {
		return
	}
	serviceInfo := &dao.ServiceInfo{
		Model: gorm.Model{
			ID: p.ID,
		},
	}
	return serviceInfo.DeleteOne(c, db)
}

func (p *ServiceDeleteInput) BindValidParam(c *gin.Context) (err error) {
	return public.DefaultGetValidParams(c, p)
}
