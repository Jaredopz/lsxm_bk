package main

import (
	"log"
	"lsxm_bk/common"
	"lsxm_bk/router"
	"net/http"
)

func init() {
	// 模板加载
	common.LoadTemplate()
}

func main() {
	// 程序入口，一个项目只能有一个入口
	// web程序，http协议 ip port
	server := http.Server{
		Addr: "127.0.0.1:8090",
	}
	// 路由
	router.Router()
	// 监听对应的端口并进行启动
	// 监听不为空
	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
