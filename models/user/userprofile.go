package user

import (
	"log"
	DB "main/db"
)

//用户个人信息
type Userprofile struct {
	Id       int    `json:User_id`
	Nickname string `json:User_Nickname`
	Sex      int    `json:User_Sex`
	Addr     string `User_Addr`
}

//创建用户个人信息
func CreateUser(pro Userprofile) (int, bool) {

	template := "Insert into UserProfile Values (?,?,?,?),"
	stmt, err := DB.DB().Prepare(template)
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
func UpdateUser(pro Userprofile) bool {

	template := "UPDATE UserProfile SET Nickname=?,Sex=?,Addr=? Where Id=?"
	stmt, err := DB.DB().Prepare(template)
	if err != nil {
		log.Print(err)
		//return -1, false
		return false
	}
	_, err = stmt.Exec(pro.Nickname, pro.Sex, pro.Addr, pro.Id)
	if err != nil {
		log.Print(err)
		//return -1, false
		return false
	}
	//i, _ := res.LastInsertId() ?????? update 还查id?
	//return int(i), true
	return true
}

// 通过id 查询信息
func QueryUser(id int) (Userprofile, bool) {

	var pro Userprofile
	template := "Select * from UserProfile where Id=?,"
	rows, err := DB.DB().Query(template, id)
	if err != nil {
		log.Print(err)
		return pro, false
	}
	rows.Next()
	err = rows.Scan(pro.Id, pro.Nickname, pro.Sex, pro.Addr)
	if err != nil {
		log.Print(err)
		return pro, false
	}
	return pro, true
}
