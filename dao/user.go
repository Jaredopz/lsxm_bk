package dao

import "log"

func GetUserNameById(userId int) string {
	sqlStr := "select user_name from blog_user where uid = ?"
	rows := DB.QueryRow(sqlStr, userId)
	if rows.Err() != nil {
		log.Println(rows.Err())
	}
	var userName string
	_ = rows.Scan(&userName)
	return userName
}
