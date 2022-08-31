package service

import (
	"html/template"
	"log"
	"lsxm_bk/config"
	"lsxm_bk/dao"
	"lsxm_bk/models"
)

// 业务逻辑函数

func GetAllIndexInfo(page int, pageSize int) (*models.HomeResponse, error) {
	categorys, err := dao.GetAllCategory()
	if err != nil {
		log.Println("获取category出错", err)
		return nil, err
	}
	posts, err := dao.GetPostPage(page, pageSize)
	var postMores []models.PostMore
	for _, post := range posts {
		var postMore models.PostMore
		categoryName := dao.GetCategoryNameById(post.CategoryId)
		userName := dao.GetUserNameById(post.UserId)
		content := template.HTML(post.Content)[0:200]
		postMore.Pid = post.Pid
		postMore.Title = post.Title
		postMore.Slug = post.Slug
		postMore.Content = content
		postMore.CategoryId = post.CategoryId
		postMore.UserId = post.UserId
		postMore.ViewCount = post.ViewCount
		postMore.Type = post.Type
		postMore.CreateAt = models.DateDay(post.CreateAt)
		postMore.UpdateAt = models.DateDay(post.UpdateAt)
		postMore.CategoryName = categoryName
		postMore.UserName = userName
		postMores = append(postMores, postMore)
	}
	total := dao.CountGetAllPost()
	pageInt := (total-1)/10 + 1
	pages := make([]int, pageInt)
	for i := 0; i < pageInt; i++ {
		pages[i] = i + 1
	}
	var hr = &models.HomeResponse{
		Viewer:    config.Cfg.Viewer,
		Categorys: categorys,
		Posts:     postMores,
		Total:     total,
		Page:      page,
		Pages:     pages,
		PageEnd:   page != pageSize,
	}
	return hr, err
}

func GetPostPageByCategoryId(cId int, page int, pageSize int) (*models.CategoryResponse, error) {
	categorys, err := dao.GetAllCategory()
	if err != nil {
		log.Println("获取category出错", err)
		return nil, err
	}
	categoryName := dao.GetCategoryNameById(cId)
	posts, err := dao.GetPostsByCategoryId(cId, page, pageSize)
	var postMores []models.PostMore
	for _, post := range posts {
		var postMore models.PostMore
		userName := dao.GetUserNameById(post.UserId)
		content := template.HTML(post.Content)[0:200]
		postMore.Pid = post.Pid
		postMore.Title = post.Title
		postMore.Slug = post.Slug
		postMore.Content = content
		postMore.CategoryId = post.CategoryId
		postMore.UserId = post.UserId
		postMore.ViewCount = post.ViewCount
		postMore.Type = post.Type
		postMore.CreateAt = models.DateDay(post.CreateAt)
		postMore.UpdateAt = models.DateDay(post.UpdateAt)
		postMore.CategoryName = categoryName
		postMore.UserName = userName
		postMores = append(postMores, postMore)
	}
	total := dao.CountGetPostByCategoryId(cId)
	pageInt := (total-1)/10 + 1
	pages := make([]int, pageInt)
	for i := 0; i < pageInt; i++ {
		pages[i] = i + 1
	}
	var hr = &models.HomeResponse{
		Viewer:    config.Cfg.Viewer,
		Categorys: categorys,
		Posts:     postMores,
		Total:     total,
		Page:      page,
		Pages:     pages,
		PageEnd:   page != pageSize,
	}
	var cr = &models.CategoryResponse{
		hr,
		categoryName,
	}
	return cr, err
}
