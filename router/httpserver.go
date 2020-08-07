package router

import (
	"context"
	"gateway/lib"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

var (
	HttpSrvHandler *http.Server
)

func HttpServerRun() {
	conf := lib.GetDefaultConfBase()
	gin.SetMode(conf.Base.DebugMode)
	r := InitRouter()
	HttpSrvHandler = &http.Server{
		Addr:           conf.Http.Addr,
		Handler:        r,
		ReadTimeout:    time.Duration(conf.Http.ReadTimeout) * time.Second,
		WriteTimeout:   time.Duration(conf.Http.WriteTimeout) * time.Second,
		MaxHeaderBytes: 1 << uint(conf.Http.MaxHeaderBytes),
	}
	go func() {

		log.Printf(" [INFO] HttpServerRun:%s\n", conf.Http.Addr)
		if err := HttpSrvHandler.ListenAndServe(); err != nil {
			log.Fatalf(" [ERROR] HttpServerRun:%s err:%v\n", conf.Http.Addr, err)
		}
	}()
}

func HttpServerStop() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := HttpSrvHandler.Shutdown(ctx); err != nil {
		log.Fatalf(" [ERROR] HttpServerStop err:%v\n", err)
	}
	log.Printf(" [INFO] HttpServerStop stopped\n")
}
