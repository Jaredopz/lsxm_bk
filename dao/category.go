package dao

import (
	"log"
	"lsxm_bk/models"
)

func GetCategoryNameById(categoryId int) string {
	sqlStr := "select name from blog_category where cid = ?"
	rows := DB.QueryRow(sqlStr, categoryId)
	if rows.Err() != nil {
		log.Println(rows.Err())
	}
	var categoryName string
	_ = rows.Scan(&categoryName)
	return categoryName
}

func GetAllCategory() ([]models.Category, error) {
	sqlStr := "select * from blog_category"
	// 如若上面不写详细，下面的Query也不需要id参数
	rows, err := DB.Query(sqlStr)
	if err != nil {
		log.Println("Category查询失败", err)
		panic(err)
	}
	var categorys []models.Category
	for rows.Next() {
		var category models.Category
		err = rows.Scan(&category.Cid, &category.Name, &category.CreateAt, &category.UpdateAt)
		if err != nil {
			log.Println("Category输出出错", err)
			return nil, err
		}
		categorys = append(categorys, category)
	}
	return categorys, nil
}
