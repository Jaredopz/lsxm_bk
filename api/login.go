package api

import (
	"lsxm_bk/common"
	"lsxm_bk/service"
	"net/http"
)

func (*Api) Login(w http.ResponseWriter, r *http.Request) {
	// 接收用户名和密码 返回对应的json数据
	// json不能通过r.ParseForm.Set()进行获取
	params := common.GetRequestJsonParam(r)
	username := params["username"].(string)
	passwd := params["passwd"].(string)
	loginRes, err := service.Login(username, passwd)
	if err != nil {
		common.Error(w, err)
		return
	}
	common.Success(w, loginRes)
}
