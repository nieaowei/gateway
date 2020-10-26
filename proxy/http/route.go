package proxy_http

import (
	"gateway/router"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitHttpProxyRouter(middlewares ...gin.HandlerFunc) *gin.Engine {
	route := gin.Default()
	route.LoadHTMLGlob("./templates/*")
	route.Use(middlewares...)
	route.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	route.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{
			"msg": "not found",
		})
	})
	route.NoMethod(func(c *gin.Context) {
		c.JSON(404, gin.H{
			"msg": "Method not found",
		})
	})
	route.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "welcome.html", gin.H{
			"title":        "HTTP/HTTPS代理服务",
			"service_name": "Http/Https Proxy Service",
			"welcome_msg":  "The service is started.",
			"items": []gin.H{
				{
					"name": "API Document",
					"link": "swagger/index.html",
					"tag":  "Swagger",
				},
			},
		})
	})

	router.Register(route, &OauthController{})
	//root := route.Group("/")
	route.Use(
		HTTPAccessModeMiddleware(),
		HTTPFlowStatisticMiddleware(),
		HTTPJwtAuthTokenMiddleware(),
		HTTPBlackListMiddleware(),
		HTTPWhiteListMiddleware(),
		HTTPHeaderTransferMiddleware(),
		HTTPStripUriMiddleware(),
		HTTPReverseProxyMiddleware(),
	)

	return route
}
