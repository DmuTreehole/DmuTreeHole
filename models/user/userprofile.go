package user

import (
	"log"
	DB "main/db"
)

//用户个人信息
type Userprofile struct {
	Id       int    `json:"UserId"`
	Sex      int    `json:"UserSex"`
	Nickname string `json:"UserNickname"`
	Addr     string `json:"UserAddr"`
}

//创建用户个人信息
func CreateUser(pro Userprofile) error {
	template := "Insert into Userprofile Set User_Id=?,User_NickName=?,User_Sex=?,User_Addr=?"
	//stmt, err := DB.DB().Prepare(template)
	//if err != nil {
	//	log.Print(err)
	//	return err
	//}
	//res, err := stmt.Exec(pro.Id, pro.Nickname, pro.Sex, pro.Addr)
	//if err != nil {
	//	log.Print(err)
	//	return err
	//}
	//return nil
	rows, err := DB.DB().Query(template, pro.Id, pro.Nickname, pro.Sex, pro.Addr)
	defer rows.Close()
	return err
}

//更改用户个人信息
func UpdateUser(pro Userprofile) error {

	template := "UPDATE Userprofile SET User_Nickname=?,User_Sex=?,User_Addr=? Where User_Id=?"
	rows, err := DB.DB().Query(template, pro.Nickname, pro.Sex, pro.Addr, pro.Id)
	defer rows.Close()
	//i, _ := res.LastInsertId() ?????? update 还查id?
	//return int(i), true
	return err
}

// 通过id 查询信息
func QueryUser(id int) (Userprofile, error) {
	var pro Userprofile
	template := "Select User_Id,User_Nickname,User_Sex,User_Addr from Userprofile where User_Id=?"
	rows, err := DB.DB().Query(template, id)
	defer rows.Close()
	if err != nil {
		log.Print(err)
		return pro, err
	}
	rows.Next()
	err = rows.Scan(&pro.Id, &pro.Nickname, &pro.Sex, &pro.Addr)
	if err != nil {
		log.Print(err)
		return pro, err
	}
	return pro, err
}
