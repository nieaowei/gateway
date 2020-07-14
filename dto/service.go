package dto

import (
	"gateway/dao"
	"gateway/public"
	"github.com/e421083458/golang_common/lib"
	"github.com/gin-gonic/gin"
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
	List  []ServiceListItem `json:"page_no" form:"page_no" validate:""`
}

func (p *ServiceListInput) GetServiceList(c *gin.Context) (out *ServiceListOutput, err error) {
	serviceInfo := &dao.ServiceInfo{}
	db, err := lib.GetGormPool("default")
	if err != nil {
		return
	}
	scan := serviceInfo.PageList(c, db, &dao.PageSize{
		Size: p.PageSize,
		No:   p.PageNo,
		Info: p.Info,
	})
	out = &ServiceListOutput{}
	err = scan.Scan(&out.List).Count(&out.Total).Error
	if err != nil {
		return
	}
	return
}

func (p *ServiceListInput) BindValidParam(c *gin.Context) (err error) {
	return public.DefaultGetValidParams(c, p)
}
