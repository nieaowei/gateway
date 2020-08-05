package dto

import (
	"encoding/gob"
	"gateway/dao"
	"github.com/gin-gonic/gin"
)

func init() {
	gob.Register(&dao.AdminSessionInfo{})
	gob.Register(&GetServiceDetailOutput{})
}

type DTO interface {
	BindValidParam(c *gin.Context) (err error)
	Exec(c *gin.Context) (out interface{}, err error)
	ErrorHandle(c *gin.Context, err error)
	OutputHandle(c *gin.Context, outIn interface{}) (out interface{})
}

func Exec(dto DTO, c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			dto.ErrorHandle(c, err)
			return
		}
	}()
	if err = dto.BindValidParam(c); err != nil {
		//ResponseError(c, 1001, err)
		return
	}
	var out interface{}
	out, err = dto.Exec(c)
	if err != nil {
		//ResponseError(c, 1002, err)
		return
	}
	ResponseSuccess(c, dto.OutputHandle(c, out))
	return
}
