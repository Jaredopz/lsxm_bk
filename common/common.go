package common

import (
	"lsxm_bk/config"
	"lsxm_bk/models"
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
