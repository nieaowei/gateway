package dto

import "github.com/gin-gonic/gin"

type DTO interface {
	BindValidParam(c *gin.Context) error
	Exec(c *gin.Context) (out interface{}, err error)
	ErrorHandle(c *gin.Context, err error)
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
	ResponseSuccess(c, out)
	return
}
