package dao

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"strings"
	"time"
)

var DB *sql.DB

var (
	driveName  = "mysql"
	dbUserName = "root"
	dbPassWord = "root"
	dbHostsIp  = "127.0.0.1:3306"
	dbName     = "goblog"
)

func init() {
	// 执行main之前，先执行init方法
	dbInfo := strings.Join([]string{dbUserName, ":", dbPassWord, "@tcp(", dbHostsIp, ")/",
		dbName, "?charset=utf8mb4&loc=Local&parseTime=true"}, "")
	db, err := sql.Open(driveName, dbInfo)
	if err != nil {
		log.Println("连接数据库异常")
		panic(err)
	}
	// 最大空闲连接数，默认不配置，是2个最大空闲连接
	db.SetMaxIdleConns(5)
	// 最大连接数，默认不配置，是不限制最大连接数
	db.SetMaxOpenConns(100)
	// 连接最大存活时间
	db.SetConnMaxLifetime(time.Minute * 3)
	// 空闲连接最大存活时间
	db.SetConnMaxIdleTime(time.Minute * 1)
	err = db.Ping()
	if err != nil {
		log.Println("数据库无法连接")
		_ = db.Close()
		panic(err)
	}
	DB = db
}
