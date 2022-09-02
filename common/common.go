package common

import (
	"encoding/json"
	"io"
	"log"
	"lsxm_bk/config"
	"lsxm_bk/models"
	"net/http"
	"sync"
)

// 通用文件

var Template models.HtmlTemplate

func LoadTemplate() {
	// 协程
	w := sync.WaitGroup{}
	w.Add(1)
	go func() {
		// 可能耗时，加快处理
		var err error
		Template, err = models.InitTemplate(config.Cfg.System.CurrentDir + "/template/")
		if err != nil {
			panic(err)
		}
		w.Done()
	}()
	w.Wait()
}

// json数据返回函数

func GetRequestJsonParam(r *http.Request) map[string]interface{} {
	var params map[string]interface{}
	body, _ := io.ReadAll(r.Body)
	_ = json.Unmarshal(body, &params)
	return params
}

// 通用的数据返回
// 成功返回

func Success(w http.ResponseWriter, data interface{}) {
	var res models.Result
	res.Code = 200
	res.Error = ""
	res.Data = data
	resultJson, _ := json.Marshal(res)
	w.Header().Set("Content-Type", "application/json")
	_, err := w.Write(resultJson)
	if err != nil {
		log.Println(err)
	}
}

// 错误返回

func Error(w http.ResponseWriter, err error) {
	var res models.Result
	res.Code = -999
	res.Error = err.Error()
	resultJson, _ := json.Marshal(res)
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(resultJson)
	if err != nil {
		log.Println(err)
	}
}
