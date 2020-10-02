package proxy_http

import (
	"context"
	"gateway/lib"
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
	r := InitHttpProxyRouter()
	HttpServer = &http.Server{
		Addr:           lib.GetDefaultConfProxy().Http.Addr,
		Handler:        r,
		ReadTimeout:    time.Second * time.Duration(lib.GetDefaultConfProxy().Http.ReadTimeout),
		WriteTimeout:   time.Second * time.Duration(lib.GetDefaultConfProxy().Http.WriteTimeout),
		MaxHeaderBytes: 1 << uint(lib.GetDefaultConfProxy().Http.MaxHeaderBytes),
	}
	go func() {
		log.Printf("[INFO] http proxy %v running \n", lib.GetDefaultConfProxy().Http.Addr)
		if err := HttpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf(" [ERROR] http_proxy_run %s err:%v\n", lib.GetDefaultConfProxy().Http.Addr, err)
		}
	}()
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
