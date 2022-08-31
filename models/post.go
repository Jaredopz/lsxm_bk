package models

import (
	"html/template"
	"lsxm_bk/config"
	"time"
)

// Post 数据库映射相关的
type Post struct {
	Pid        int       `json:"pid"`        // 文章ID
	Title      string    `json:"title"`      // 文章ID
	Slug       string    `json:"slug"`       // 自定义页面 path
	Content    string    `json:"content"`    // 文章的html
	Markdown   string    `json:"markdown"`   // 文章的markdown
	CategoryId int       `json:"categoryId"` // 分类ID
	UserId     int       `json:"userId"`     // 用户ID
	ViewCount  int       `json:"viewCount"`  // 查看次数
	Type       int       `json:"type"`       // 文章类型 0 普通 1 自定义文章
	CreateAt   time.Time `json:"createAt"`   // 创建时间
	UpdateAt   time.Time `json:"updateAt"`   // 更新时间
}

// PostMore 便于页面进行展示，相比Post多了一些字段
type PostMore struct {
	Pid          int           `json:"pid"`          // 文章ID
	Title        string        `json:"title"`        // 文章ID
	Slug         string        `json:"slug"`         // 自定义页面 path
	Content      template.HTML `json:"content"`      // 文章的html
	CategoryId   int           `json:"categoryId"`   // 分类ID
	CategoryName string        `json:"categoryName"` // 分类名
	UserId       int           `json:"userId"`       // 用户ID
	UserName     string        `json:"userName"`     // 用户名
	ViewCount    int           `json:"viewCount"`    // 查看次数
	Type         int           `json:"type"`         // 文章类型 0 普通 1 自定义文章
	CreateAt     string        `json:"createAt"`
	UpdateAt     string        `json:"updateAt"`
}

// PostReq 文章相关的请求包装
type PostReq struct {
	Pid        int    `json:"pid"`
	Title      string `json:"title"`
	Slug       string `json:"slug"`
	Content    string `json:"content"`
	Markdown   string `json:"markdown"`
	CategoryId int    `json:"categoryId"`
	UserId     int    `json:"userId"`
	Type       int    `json:"type"`
}

// SearchResp 搜索相关的返回
type SearchResp struct {
	Pid   int    `orm:"pid" json:"pid"` //文章ID
	Title string `orm:"title" json:"title"`
}

// PostRes 文章相关的返回信息
type PostRes struct {
	config.Viewer
	config.SystemConfig
	Article PostMore
}
