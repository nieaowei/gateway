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
		serviceDetail, err := FindOneServiceDetail(c, &info)
		if err != nil {
			return nil, err
		}
		serviceAddr := ""
		loadType := ""
		switch serviceDetail.Info.LoadType {
		case public.LoadTypeHttp:
			{
				loadType = "HTTP"
				if serviceDetail.HTTP.RuleType == public.HttpRuleTypePrefixURL && serviceDetail.HTTP.NeedHttps == 0 {
					serviceAddr = clusterIP + ":" + clusterPort + serviceDetail.HTTP.Rule

				}
				if serviceDetail.HTTP.RuleType == public.HttpRuleTypePrefixURL && serviceDetail.HTTP.NeedHttps == 1 {
					serviceAddr = clusterIP + ":" + clusterSSLPort + serviceDetail.HTTP.Rule
				}
				if serviceDetail.HTTP.RuleType == public.HttpRuleTypeDomain {
					serviceAddr = serviceDetail.HTTP.Rule
				}
				break
			}
		case public.LoadTypeTcp:
			{
				loadType = "TCP"
				serviceAddr = clusterIP + ":" + strconv.Itoa(int(serviceDetail.TCP.Port))
				break
			}
		case public.LoadTypeGrpc:
			{
				loadType = "GRPC"
				serviceAddr = clusterIP + ":" + strconv.Itoa(int(serviceDetail.GRPC.Port))
				break
			}
		}
		item := ServiceListItem{
			ID:          info.ID,
			ServiceName: info.ServiceName,
			ServiceDesc: info.ServiceDesc,
			LoadType:    loadType,
			Address:     serviceAddr,
			Qps:         0,
			Qpd:         0,
			TotalNode:   0,
		}
		out.List = append(out.List, item)
	}

	return
}

func (p *ServiceListInput) BindValidParam(c *gin.Context) (err error) {
	return public.DefaultGetValidParams(c, p)
}

type ServiceDetail struct {
	Info          *dao.ServiceInfo          `json:"info"`
	HTTP          *dao.ServiceHttpRule      `json:"http"`
	GRPC          *dao.ServiceGrpcRule      `json:"grpc"`
	TCP           *dao.ServiceTcpRule       `json:"tcp"`
	LoadBalance   *dao.ServiceLoadBalance   `json:"load_balance"`
	AccessControl *dao.ServiceAccessControl `json:"access_control"`
}

func FindOneServiceDetail(c *gin.Context, info *dao.ServiceInfo) (out *ServiceDetail, err error) {
	db, err := lib.GetGormPool("default")
	if err != nil {
		return
	}
	//todo wait next step optimization.
	switch info.LoadType {
	case public.LoadTypeHttp:
		{
			break
		}
	case public.LoadTypeTcp:
		{
			break
		}
	case public.LoadTypeGrpc:
		{
			break
		}
	}

	httpRule := &dao.ServiceHttpRule{
		ServiceId: info.ID,
	}
	httpRule, err = httpRule.FindOne(c, db)
	if err != nil && err != gorm.ErrRecordNotFound {
		return
	}

	tcpRule := &dao.ServiceTcpRule{
		ServiceId: info.ID,
	}
	tcpRule, err = tcpRule.FindOne(c, db)
	if err != nil && err != gorm.ErrRecordNotFound {
		return
	}

	grpcRule := &dao.ServiceGrpcRule{
		ServiceId: info.ID,
	}
	grpcRule, err = grpcRule.FindOne(c, db)
	if err != nil && err != gorm.ErrRecordNotFound {
		return
	}

	accessControl := &dao.ServiceAccessControl{
		ServiceId: info.ID,
	}
	accessControl, err = accessControl.FindOne(c, db)
	if err != nil && err != gorm.ErrRecordNotFound {
		return
	}

	loadBalance := &dao.ServiceLoadBalance{
		ServiceId: info.ID,
	}
	loadBalance, err = loadBalance.FindOne(c, db)
	if err != nil && err != gorm.ErrRecordNotFound {
		return
	}

	out = &ServiceDetail{
		Info:          info,
		HTTP:          httpRule,
		GRPC:          grpcRule,
		TCP:           tcpRule,
		LoadBalance:   loadBalance,
		AccessControl: accessControl,
	}

	return
}
