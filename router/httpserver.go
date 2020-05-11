package router

import (
	"context"
	"log"
	"net/http"
	"tbwisk/public"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	//HTTPSrvHandler 服务
	HTTPSrvHandler *http.Server
)

//HTTPServerRun 服务启动
func HTTPServerRun() {
	gin.SetMode(public.DebugMode)
	r := InitRouter()
	HTTPSrvHandler = &http.Server{
		Addr:           public.GetStringConf("http", "addr"),
		Handler:        r,
		ReadTimeout:    time.Duration(public.GetIntConf("http", "read_timeout")) * time.Second,
		WriteTimeout:   time.Duration(public.GetIntConf("http", "write_timeout")) * time.Second,
		MaxHeaderBytes: 1 << uint(public.GetIntConf("http", "max_header_bytes")),
	}
	go func() {
		log.Printf(" [INFO] HttpServerRun:%s\n", public.GetStringConf("http", "addr"))
		if err := HTTPSrvHandler.ListenAndServe(); err != nil {
			log.Fatalf(" [ERROR] HttpServerRun:%s err:%v\n", public.GetStringConf("http", "addr"), err)
		}
	}()
}

//HTTPServerStop 服务停止
func HTTPServerStop() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := HTTPSrvHandler.Shutdown(ctx); err != nil {
		log.Fatalf(" [ERROR] HttpServerStop err:%v\n", err)
	}
	log.Printf(" [INFO] HttpServerStop stopped\n")
}
