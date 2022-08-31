package models

import (
	"html/template"
	"io"
	"log"
	"time"
)

type TemplateBlog struct {
	*template.Template
}
type HtmlTemplate struct {
	Index    TemplateBlog
	Category TemplateBlog
	// 自定义页面
	Custom TemplateBlog
	// 文章详情页面
	Detail TemplateBlog
	// 登录页面
	Login TemplateBlog
	// 文章归档页面
	Pigeonhole TemplateBlog
	// 文章编写、更新发布
	Writing TemplateBlog
}

func (t *TemplateBlog) WriteData(w io.Writer, data interface{}) {
	err := t.Execute(w, data)
	if err != nil {
		_, err = w.Write([]byte(err.Error()))
		if err != nil {
			log.Println(err)
			return
		}
	}
}

func (t *TemplateBlog) WriteError(w io.Writer, err error) {
	if err != nil {
		_, err = w.Write([]byte(err.Error()))
		if err != nil {
			log.Println(err)
			return
		}
	}
}

func InitTemplate( /*模板的路径*/ templateDir string) (HtmlTemplate, error) {
	tp, err := readTemplate(
		[]string{"index", "category", "custom",
			"detail", "login", "pigeonhole", "writing"},
		templateDir,
	)
	var htmlTemplate HtmlTemplate
	if err != nil {
		log.Println(err)
		return htmlTemplate, err
	}
	htmlTemplate.Index = tp[0]
	htmlTemplate.Category = tp[1]
	htmlTemplate.Custom = tp[2]
	htmlTemplate.Detail = tp[3]
	htmlTemplate.Login = tp[4]
	htmlTemplate.Pigeonhole = tp[5]
	htmlTemplate.Writing = tp[6]
	return htmlTemplate, nil
}

// IsODD header里的循环切片索引值
func IsODD(num int) bool {
	return num%2 == 0
}

func getNextName(Navigation []string, index int) string {
	return Navigation[index+1]
}

func Date(layout string) string {
	return time.Now().Format(layout)
}

func DateDay(date time.Time) string {
	return date.Format("2006-01-02 15:04:05")
}

func readTemplate(templates []string, templateDir string) ([]TemplateBlog, error) {
	var tbs []TemplateBlog
	for _, view := range templates {
		viewName := view + ".html"
		t := template.New(viewName)
		// 访问博客首页模板的时候，因为有多个模板的嵌套，解析文件的时候，需要将其涉及到的所有模板都进行解析
		header := templateDir + "layout/header.html"
		footer := templateDir + "layout/footer.html"
		home := templateDir + "home.html"
		postList := templateDir + "layout/post-list.html"
		pagination := templateDir + "layout/pagination.html"
		personal := templateDir + "layout/personal.html"
		// 补充解析页面里的各种未定义函数
		t.Funcs(template.FuncMap{"isODD": IsODD, "getNextName": getNextName,
			"date": Date, "dateDay": DateDay})
		// 分别解析页面
		t, err := t.ParseFiles(templateDir+viewName, header, footer, home, postList, pagination, personal)
		if err != nil {
			log.Println("解析模板出错：", err)
			return nil, err
		}
		var tb TemplateBlog
		tb.Template = t
		tbs = append(tbs, tb)
	}
	return tbs, nil
}
