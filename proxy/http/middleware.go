package proxy_http

import (
	"gateway/dto"
	"gateway/proxy/manager"
	"github.com/gin-gonic/gin"
	"log"
	"net/url"
	"strings"
)

const Key_Http_Service = "Key_Http_Service"

func HTTPAccessModeMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		service, err := manager.Default().HTTPAccessMode(c)
		if err != nil {
			dto.ResponseError(c, 2000, err)
			c.Abort()
			return
		}
		c.Set(Key_Http_Service, service)
		c.Next()
	}
}

func HTTPBlackListMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		data, ok := c.Get(Key_Http_Service)
		if !ok {
			dto.ResponseError(c, Error_ServiceNotFound.Code, Error_ServiceNotFound)
			c.Abort()
			return
		}
		service := data.(manager.HTTPService)
		//whiteList := service.WhiteList
		//blackList := service.BlackList
		if service.OpenAuth && len(service.WhiteList) == 0 && len(service.BlackList) > 0 {
			for _, s := range service.BlackList {
				if s.Host == c.ClientIP() {
					dto.ResponseError(c, Error_BlackListLimit.Code, Error_BlackListLimit)
					c.Abort()
					return
				}
			}
		}
		c.Next()
	}
}

func HTTPFlowStatisticMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		data, ok := c.Get(Key_Http_Service)
		if !ok {
			dto.ResponseError(c, Error_ServiceNotFound.Code, Error_ServiceNotFound)
			c.Abort()
			return
		}
		service := data.(manager.HTTPService)

		// statistic single service
		redisService, ok := manager.Default().GetRedisService(manager.ServicePrefix + service.ServiceName)
		if !ok {
			dto.ResponseError(c, Error_NoAvailableRedisService.Code, Error_NoAvailableRedisService)
			c.Abort()
			return
		}

		redisService.Exec()
		// statistic total
		redisTotal, ok := manager.Default().GetRedisService(manager.TotalPrefix)
		if !ok {
			dto.ResponseError(c, Error_NoAvailableRedisService.Code, Error_NoAvailableRedisService)
			c.Abort()
			return
		}
		redisTotal.Exec()

		c.Next()
		return
	}
}

func HTTPReverseProxyMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		data, ok := c.Get(Key_Http_Service)
		if !ok {
			dto.ResponseError(c, Error_ServiceNotFound.Code, Error_ServiceNotFound)
			c.Abort()
			return
		}
		service := data.(manager.HTTPService)

		lb, ok := manager.Default().GetLoadBalancer(service.ServiceName)
		if !ok {
			dto.ResponseError(c, Error_LBNotFound.Code, Error_LBNotFound)
			c.Abort()
			return
		}
		target, err := lb.Get(c.Request.Host)
		if err != nil {
			dto.ResponseError(c, Error_NoAvailableHost.Code, Error_NoAvailableHost)
			c.Abort()
			return
		}
		log.Println("[LoadBalance] Proxy: " + target.String())
		proxy := NewHttpProxy(target.URL, nil, nil)
		proxy.ServeHTTP(c.Writer, c.Request)
		c.Abort()
		return
	}
}
