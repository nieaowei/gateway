package proxy_http

import (
	"gateway/dao"
	"gateway/dto"
	"gateway/proxy/manager"
	"github.com/gin-gonic/gin"
	"strings"
)

const Key_Http_Service = "Key_Http_Service"

func HTTPAccessModeMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		service, err := manager.Default().HTTPAccessMode(c)
		if err != nil {
			dto.ResponseError(c, Error_ServiceNotFound.Code, Error_ServiceNotFound)
			c.Abort()
			return
		}
		c.Set(Key_Http_Service, service)
		c.Next()
		return
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
		return
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
		target, err := lb.GetHost(c.Request.Host)
		if err != nil {
			dto.ResponseError(c, Error_NoAvailableHost.Code, Error_NoAvailableHost)
			c.Abort()
			return
		}
		//log.Println("[LoadBalance] Proxy: " + target.String())
		trans, ok := manager.Default().GetTransport(service.ServiceName)
		if !ok {
			dto.ResponseError(c, Error_NoAvailableTransport.Code, Error_NoAvailableTransport)
			c.Abort()
			return
		}
		proxy := NewHttpProxy(target.URL, trans.Transport, nil)
		//proxy := httputil.NewSingleHostReverseProxy(target.URL)
		proxy.ServeHTTP(c.Writer, c.Request)
		c.Abort()
		return
	}
}

func HTTPHeaderTransferMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		data, ok := c.Get(Key_Http_Service)
		if !ok {
			dto.ResponseError(c, Error_ServiceNotFound.Code, Error_ServiceNotFound)
			c.Abort()
			return
		}
		service := data.(manager.HTTPService)
		hdList := service.HeaderTransform
		for _, item := range hdList {
			switch item.Op {
			case dao.HeaderTransformOperation_Add, dao.HeaderTransformOperation_Edit:
				c.Request.Header.Set(item.Key, item.Val)
			case dao.HeaderTransformOperation_Del:
				c.Request.Header.Del(item.Key)
			}
		}
		c.Next()
		return
	}
}

func HTTPWhiteListMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		data, ok := c.Get(Key_Http_Service)
		if !ok {
			dto.ResponseError(c, Error_ServiceNotFound.Code, Error_ServiceNotFound)
			c.Abort()
			return
		}
		service := data.(manager.HTTPService)
		if service.OpenAuth {
			for _, ip := range service.WhiteList {
				if c.Request.URL.Host == ip.Host {
					c.Next()
					return
				}
			}
			dto.ResponseError(c, Error_WhiteListLimit.Code, Error_WhiteListLimit)
			c.Abort()
			return
		}
		c.Next()
		return
	}
}

func HTTPStripUriMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		data, ok := c.Get(Key_Http_Service)
		if !ok {
			dto.ResponseError(c, Error_ServiceNotFound.Code, Error_ServiceNotFound)
			c.Abort()
			return
		}
		service := data.(manager.HTTPService)
		if service.RuleType == dao.HttpRule_PrefixURL && service.NeedStripURI {
			c.Request.URL.Path = strings.TrimPrefix(c.Request.URL.Path, service.Rule)
		}
		c.Next()
		return
	}
}
