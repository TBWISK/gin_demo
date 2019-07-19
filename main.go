package main

import (
	"os"
	"os/signal"
	"syscall"
	"tbwisk/common/lib"
	"tbwisk/dao"
	"tbwisk/public"
	"tbwisk/router"
)

func main() {
	lib.InitModule("./conf/dev/", []string{"base", "mysql", "redis"})
	defer lib.Destroy()
	public.InitMysql()
	public.InitValidate()
	router.HttpServerRun()
	public.GormPool.AutoMigrate(&dao.Area{})
	public.GormPool.AutoMigrate(&dao.User{})

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	router.HttpServerStop()
}
