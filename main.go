package main

import (
	"gateway/lib"
	"gateway/router"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	//lib.InitModule("./conf/dev/", []string{"base", "mysql", "redis"})
	//defer lib.Destroy()
	lib.InitBaseConf("./conf/dev")
	lib.InitMysqlConf("./conf/dev")
	lib.InitRedisConf("./conf/dev")
	lib.InitDBPool()
	router.HttpServerRun()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	router.HttpServerStop()
}
