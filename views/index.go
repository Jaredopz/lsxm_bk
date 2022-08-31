package views

import (
	"errors"
	"log"
	"lsxm_bk/common"
	"lsxm_bk/service"
	"net/http"
	"strconv"
)

func (*HTMLApi) Index(w http.ResponseWriter, r *http.Request) {
	index := common.Template.Index
	// 执行文件
	// 页面上所有涉及到的数据必须有定义
	// 数据库查询
	err := r.ParseForm()
	if err != nil {
		log.Println("表单获取失败", err)
		index.WriteError(w, errors.New("系统错误，请联系管理员"))
		return
	}
	pageStr := r.Form.Get("page")
	page := 1
	if pageStr != "" {
		page, _ = strconv.Atoi(pageStr)
	}
	// 每页显示的数量
	pageSize := 10
	hr, err := service.GetAllIndexInfo(page, pageSize)
	if err != nil {
		log.Println("Index获取数据出错", err)
		index.WriteError(w, errors.New("系统错误，请联系管理员"))
	}
	index.WriteData(w, hr)
}
