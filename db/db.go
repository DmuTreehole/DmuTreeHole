package db

import (
	"database/sql"
	"github.com/go-ini/ini"
	_ "github.com/go-sql-driver/mysql"
)

type mysql struct {
	ip       string
	port     string
	user     string
	password string
	database string
}

var DB *sql.DB

//初始化数据库
func init() {
	ctf, err := ini.Load("conf/app.ini")
	mysql := &mysql{
		ip:       ctf.Section("mysql").Key("ip").String(),
		port:     ctf.Section("mysql").Key("port").String(),
		user:     ctf.Section("mysql").Key("user").String(),
		password: ctf.Section("mysql").Key("password").String(),
		database: ctf.Section("mysql").Key("database").String(),
	}
	dsn := mysql.user + ":" + mysql.password + "@tcp(" + mysql.ip + ":" + mysql.port + ")/" + mysql.database
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
	}
}
