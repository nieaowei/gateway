package main

import (
	"flag"
	_ "gateway/docs"
	"gateway/lib"
	proxy_gprc "gateway/proxy/grpc"
	proxy_http "gateway/proxy/http"
	"gateway/proxy/manager"
	proxy_tcp "gateway/proxy/tcp"
	"gateway/router"
	"os"
	"os/signal"
	"syscall"
)

// @title 微服务网关接口文档
// @version 1.0
// @description This is a sample server celler server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8880
// @BasePath /
// @query.collection.format multi

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @securitydefinitions.oauth2.application OAuth2Application
// @tokenUrl https://example.com/oauth/token
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.implicit OAuth2Implicit
// @authorizationurl https://example.com/oauth/authorize
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.password OAuth2Password
// @tokenUrl https://example.com/oauth/token
// @scope.read Grants read access
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.accessCode OAuth2AccessCode
// @tokenUrl https://example.com/oauth/token
// @authorizationurl https://example.com/oauth/authorize
// @scope.admin Grants read and write access to administrative information
// @x-extension-openapi {"example": "value on a json format"}

func main() {
	//lib.InitModule("./conf/dev/", []string{"base", "mysql", "redis"})
	//defer lib.Destroy()
	conf := flag.String("conf", "dev", "dev or pro")
	swag := flag.Bool("swag", false, "true and false")
	endpoint := flag.String("endpoint", "dashboard", "dashboard,http,tcp and grpc")
	flag.Parse()

	if *conf == "pro" {
		lib.InitBaseConf("./conf/pro")
		lib.InitMysqlConf("./conf/pro")
		lib.InitRedisConf("./conf/pro")
		lib.InitProxyConf("./conf/pro")
		lib.InitDBPool()
	} else {
		lib.InitBaseConf("./conf/dev")
		lib.InitMysqlConf("./conf/dev")
		lib.InitRedisConf("./conf/dev")
		lib.InitProxyConf("./conf/dev")
		lib.InitRedis()
		lib.InitDBPool()
	}
	//db,_:=lib.GetDefaultDB()
	//db.Logger.LogMode(logger.Info)
	switch *endpoint {
	case "dashboard":
		{
			router.HttpServerRun(*swag)
			defer router.HttpServerStop()
		}
	case "http":
		{
			manager.InitManager()
			proxy_http.HttpServerRun()
			defer proxy_http.HttpServerStop()
		}
	case "tcp":
		{
			proxy_tcp.TcpServerRun()
			defer proxy_tcp.TcpServerStop()
		}
	case "grpc":
		{
			proxy_gprc.GrpcServerRun()
			defer proxy_gprc.GrpcServerStop()
		}
	}
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
	<-quit

}
