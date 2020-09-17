package proxy

import (
	"context"
	"gateway/lib"
	"gateway/router"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

var (
	HttpServer *http.Server
)

func HttpServerRun() {
	gin.SetMode(lib.GetDefaultConfProxy().Base.DebugMode)

}

func HttpServerStop() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := HttpServer.Shutdown(ctx)
	if err != nil {
		log.Fatalf("[INFO] http proxy stop err:%v\n", err)
	}
	log.Printf("[INFO] http proxy %v stopped \n", lib.GetDefaultConfProxy().Http.Addr)
}

func InitHttpProxyRouter(middlewares ...gin.HandlerFunc) *gin.Engine {
	route := gin.New()
	route.Use(middlewares...)
	route.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.Register(route, &OauthController{})

	route.Use()

	return route
}
