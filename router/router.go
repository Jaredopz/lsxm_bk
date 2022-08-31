package router

import (
	"lsxm_bk/api"
	"lsxm_bk/views"
	"net/http"
)

func Router() {
	// 设置路由
	// 1.页面 views 2. api 数据（json) 3.静态资源
	// 1.页面
	http.HandleFunc("/", views.HTML.Index)
	http.HandleFunc("/c/", views.HTML.Category)
	// 2.api处理
	http.HandleFunc("/api/v1/post", api.API.SaveAndUpdatePost)
	// 路由与resource静态信息匹配
	// 3.静态资源
	http.Handle("/resource/", http.StripPrefix("/resource/", http.FileServer(http.Dir("public/resource/"))))
}
