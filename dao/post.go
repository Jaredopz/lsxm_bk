package dao

import (
	"log"
	"lsxm_bk/models"
)

func GetPostsByCategoryId(cId int, page int, pageSize int) ([]models.Post, error) {
	page = (page - 1) * pageSize
	sqlStr := "select * from blog_post where category_id = ? limit ?, ?"
	rows, err := DB.Query(sqlStr, cId, page, pageSize)
	if err != nil {
		log.Println("查询行数失败", err)
		return nil, err
	}
	var posts []models.Post
	for rows.Next() {
		var post models.Post
		_ = rows.Scan(
			&post.Pid,
			&post.Title,
			&post.Content,
			&post.Markdown,
			&post.CategoryId,
			&post.UserId,
			&post.ViewCount,
			&post.Type,
			&post.Slug,
			&post.CreateAt,
			&post.UpdateAt)
		posts = append(posts, post)
	}
	return posts, err
}

func CountGetPostByCategoryId(cId int) (count int) {
	sqlStr := "select count(1) from blog_post where category_id = ?"
	row := DB.QueryRow(sqlStr, cId)
	_ = row.Scan(&count)
	return
}

func CountGetAllPost() (row int) {
	// QueryRow返回的是row，其调用方法Scan返回与行所匹配的值
	// 而Query返回的是rows，其调用方法Scan需要dest与rows中数据量相同
	// 因这里需要查询总条数，所以要用QueryRow
	sqlStr := "select count(1) from blog_post"
	rows := DB.QueryRow(sqlStr)
	_ = rows.Scan(&row)
	return
}

func GetPostPage(page int, pageSize int) ([]models.Post, error) {
	sqlStr := "select * from blog_post limit ?,?"
	page = (page - 1) * pageSize
	rows, err := DB.Query(sqlStr, page, pageSize)
	if err != nil {
		log.Println("查询行数失败", err)
		return nil, err
	}
	var posts []models.Post
	for rows.Next() {
		var post models.Post
		_ = rows.Scan(
			&post.Pid,
			&post.Title,
			&post.Content,
			&post.Markdown,
			&post.CategoryId,
			&post.UserId,
			&post.ViewCount,
			&post.Type,
			&post.Slug,
			&post.CreateAt,
			&post.UpdateAt)
		posts = append(posts, post)
	}
	return posts, err
}
