package dao

import (
	"log"
	"lsxm_bk/models"
)

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

func GetUser(userName string, passwd string) *models.User {
	sqlStr := "select * from blog_user where user_name = ? AND passwd = ?"
	rows := DB.QueryRow(sqlStr, userName, passwd)
	if rows.Err() != nil {
		log.Println(rows.Err())
		return nil
	}
	var user models.User
	err := rows.Scan(&user.Uid,
		&user.UserName,
		&user.Passwd,
		&user.Avatar,
		&user.CreateAt,
		&user.UpdateAt)
	if err != nil {
		log.Println(err)
		return nil
	}
	return &user
}
