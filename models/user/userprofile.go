package user

import (
	"database/sql"
	"encoding/json"
	"log"
)

//用户个人信息
type Userprofile struct {
	Id       int    `json:User_id`
	Nickname string `json:User_Nickname`
	Sex      int    `json:User_Sex`
	Addr     string `User_Addr`
}

//创建用户个人信息
func CreateUser(pro Userprofile, db *sql.DB) (int, bool) {

	template := "Insert into UserProfile Values (?,?,?,?),"
	stmt, err := db.Prepare(template)
	if err != nil {
		log.Print(err)
		return -1, false
	}
	res, err := stmt.Exec(pro.Id, pro.Nickname, pro.Sex, pro.Addr)
	if err != nil {
		log.Print(err)
		return -1, false
	}
	id, _ := res.LastInsertId()
	return int(id), true

}

//更改用户个人信息
func UpdateUser(pro Userprofile, db *sql.DB) (int, bool) {

	template := "UPDATE UserProfile SET Nickname=?,Sex=?,Addr=? Where Id=?"
	stmt, err := db.Prepare(template)
	if err != nil {
		log.Print(err)
		return -1, false
	}
	res, err := stmt.Exec(pro.Nickname, pro.Sex, pro.Addr, pro.Id)
	if err != nil {
		log.Print(err)
		return -1, false
	}
	i, _ := res.LastInsertId()
	return int(i), true

}

// 通过id 查询信息
func QueryUser(id int, db *sql.DB) []byte {

	var pro Userprofile
	template := "Select * from UserProfile where Id=?,"
	rows, err := db.Query(template, id)
	if err != nil {
		log.Print(err)
		return nil
	}

	for rows.Next() {
		err = rows.Scan(pro.Id, pro.Nickname, pro.Sex, pro.Addr)
		if err != nil {
			log.Print(err)
			return nil
		}
	}

	res, err := json.Marshal(pro)
	if err != nil {
		log.Print(err)
		return nil
	}
	return res
}
