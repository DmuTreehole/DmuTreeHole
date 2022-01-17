package db

import (
	"database/sql"
	"github.com/go-ini/ini"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var (
	dbs *sql.DB
	dsn string
	err error
)

type db struct {
	SourceIP string `ini:"ip"`
	Port     string `ini:"port"`
	Account  string `ini:"user"`
	PassWord string `ini:"password"`
	Database string `ini:"database"`
}

//初始化数据库
func init() {
	var mysql db
	cfg, _ := ini.Load("conf/app.ini")
	cfg.Section("mysql").MapTo(&mysql)
	//dsn = "root:dmutreehole@tcp(www.wonend.cn:3306)/Server"
	dsn = mysql.Account + `:` + mysql.PassWord + `@tcp(` + mysql.SourceIP + `:` + mysql.Port + `)/` + mysql.Database
	dbs, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("数据库连接失败")
	}
}

func DB() *sql.DB {
	return dbs
}
