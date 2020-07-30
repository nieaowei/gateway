package main

import (
	"gateway/lib"
	"gateway/router"
	"gorm.io/gorm"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	lib.InitModule("./conf/dev/", []string{"base", "mysql", "redis"})
	defer lib.Destroy()
	for _, val := range lib.GORMMapPool {
		val.SetLogger(gorm.Logger{log.New(os.Stdout, "\r\n", 0)})
		val.LogMode(true)
		val.SingularTable(true)
	}
	router.HttpServerRun()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	router.HttpServerStop()
}
