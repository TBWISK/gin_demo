package router

import (
	"context"
	"log"
	"net/http"
	"tbwisk/common/lib"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	//HTTPSrvHandler 服务
	HTTPSrvHandler *http.Server
)

//HTTPServerRun 服务启动
func HTTPServerRun() {
	gin.SetMode(lib.ConfBase.DebugMode)
	r := InitRouter()
	HTTPSrvHandler = &http.Server{
		Addr:           lib.GetStringConf("http", "addr"),
		Handler:        r,
		ReadTimeout:    time.Duration(lib.GetIntConf("http", "read_timeout")) * time.Second,
		WriteTimeout:   time.Duration(lib.GetIntConf("http", "write_timeout")) * time.Second,
		MaxHeaderBytes: 1 << uint(lib.GetIntConf("http", "max_header_bytes")),
	}
	go func() {
		log.Printf(" [INFO] HttpServerRun:%s\n", lib.GetStringConf("http", "addr"))
		if err := HTTPSrvHandler.ListenAndServe(); err != nil {
			log.Fatalf(" [ERROR] HttpServerRun:%s err:%v\n", lib.GetStringConf("http", "addr"), err)
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
