package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var (
	dbs *sql.DB
	dsn string
	err error
)

//初始化数据库
func init() {
	dsn = "root:dmutreehole@tcp(www.wonend.cn:3306)/Server"
	dbs, err = sql.Open("mysql", dsn)
	err = dbs.Ping()
	if err != nil {
		log.Panic("数据库连接失败")
	}
}

func DB() *sql.DB {
	return dbs
}
