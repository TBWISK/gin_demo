package main

import (
	"os"
	"os/signal"
	"syscall"
	"tbwisk/dao"
	"tbwisk/public"
	"tbwisk/router"
)

func main() {
	public.InitModule("/Users/tbwisk/coding/gitee/gin_demo/")
	defer public.Destroy()
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
