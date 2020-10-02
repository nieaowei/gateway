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
			c.Abort()
			return
		}
		service := data.(manager.HTTPService)
		whiteList := strings.Split(service.WhiteList, "\n")
		blackList := strings.Split(service.BlackList, "\n")
		if service.OpenAuth == 1 && len(whiteList) == 0 && len(blackList) > 0 {
			for _, s := range blackList {
				if s == c.ClientIP() {
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
			c.Abort()
			return
		}
		service := data.(manager.HTTPService)
		redisService, ok := manager.Default().GetRedisService(manager.ServicePrefix + service.ServiceName)
		if !ok {
			c.Abort()
			return
		}
		redisService.Exec()

		redisTotal, ok := manager.Default().GetRedisService(manager.TotalPrefix)
		if !ok {
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
			dto.ResponseError(c, 2002, Error_ServiceNotFound)
			c.Abort()
			return
		}
		service := data.(manager.HTTPService)
		//ipList := service.GetIPListByModel()
		// target, err := url.Parse("http://" + ipList[0])
		lb, ok := manager.Default().GetLoadBalancer(service.ServiceName)
		if !ok {
			dto.ResponseError(c, 2002, Error_LBNotFound)
			c.Abort()
			return
		}
		target, err := lb.Get(c.Request.Host)
		if err != nil {
			dto.ResponseError(c, 2002, Error_NoAvailableHost)
			c.Abort()
			return
		}
		log.Println("[LoadBalance] Proxy: " + target.String())
		proxy := NewHttpProxy(target, nil, func(url2 *url.URL) string {
			return strings.TrimPrefix(url2.Path, service.Rule)
		})
		proxy.ServeHTTP(c.Writer, c.Request)
		c.Abort()
		return
	}
}
