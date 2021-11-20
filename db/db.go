package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var (
	Dbs *sql.DB
	dsn string
	err error
)

//初始化数据库
func init() {
	dsn = "root:dmutreehole@tcp(www.wonend.cn:3306)/Server"
	Dbs, err = sql.Open("mysql", dsn)
	err = Dbs.Ping()
	if err != nil {
		log.Panic("数据库连接失败")
	}

}
