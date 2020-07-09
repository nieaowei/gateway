package dto

import "gateway/dao"

type AdminInfoOutput struct {
	*dao.AdminSessionInfo
	Avatar       string   `json:"avatar"`
	Introduction string   `json:"introduction"`
	Roles        []string `json:"roles"`
}

type AdminInfoInput struct {
}
