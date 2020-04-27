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
	lib.InitModule("/Users/tbwisk/coding/gitee/gin_demo/", []string{"base", "mysql", "redis"})
	defer lib.Destroy()
	public.InitMysql()
	public.InitValidate()
	router.HTTPServerRun()
	public.GormPool.AutoMigrate(&dao.Area{})
	public.GormPool.AutoMigrate(&dao.User{})

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	router.HTTPServerStop()
}
