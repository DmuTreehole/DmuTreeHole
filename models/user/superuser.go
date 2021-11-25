package user
import(
	DB"main/db"
	"log"
)
type Baned struct{
	username string
	reason string
}
//通过树洞删除对应的用户
func BanUserByPostid(pid int,reason string) (int,bool){
	template:="select User_Id,User_Name from Post,User where Post_Id=? and User.User_Id=Post.User_Id"
	rows, err := DB.DB().Query(template, pid)
	if err != nil {
		log.Print(err)
	}
	rows.Next()
	var uid int 
	var username string
	err=rows.Scan(&uid,&username)
	if err != nil {
		log.Print(err)
	}
	template="insert Baneduser Set User_Id=?,User_Name=?,Reason=?"
	stmt, err := DB.DB().Prepare(template)
	if err != nil {
		log.Print(err)
	}
	result, err := stmt.Exec(uid,username,reason)
	if err != nil {
		log.Print(err)
	}
	id, _ := result.LastInsertId()
	return int(id), true
}
func ShowBannedUsers()([]Baned,error){
		template := "Select User_Name,Reason from Baneduser"
		rows, err := DB.DB().Query(template)
		if err != nil {
			log.Print(err)
		}
		banedlist := []Baned{}
		for rows.Next() {
			baned := Baned{}
			err = rows.Scan(&baned.username ,&baned.reason)
			if err != nil {
				log.Print(err)
			}
			banedlist=append(banedlist, baned)
		}
		return banedlist, nil
}