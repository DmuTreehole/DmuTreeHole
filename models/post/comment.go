package post

import (
	"log"
	DB "main/db"
	Tools "main/utils"
)

type Comment struct {
	Uid     int    `json:"User_Id"`
	Pid     int    `json:"Post_Id"`
	Content string `json:"Content"`
}
type Comment_view struct{
	Id      int    `json:"Comment_Id"`
	Username string 
	Updated string
	Created string
	Uid     int    `json:"User_Id"`
	Pid     int    `json:"Post_Id"`
	Content string `json:"Content"`
}

//创建评论
func CreateComment(comment Comment) (int64,error){
	template := "Insert Comment Set Created=?,User_Id=?,Post_Id=?Updated=? Content=?"
	stmt, err := DB.DB().Prepare(template)
	if err != nil {
		log.Print(err)
	}
	created := Tools.GetDatetime()
	updated := Tools.GetDatetime()
	result, err := stmt.Exec(created,comment.Uid,comment.Pid, updated, comment.Content)
	if err != nil {
		log.Print(err)
	}
	id, _ := result.LastInsertId()
	return id, err
}

//查看评论
func ShowComment(pid int)([]Comment_view,error) {
	//创建更改时间，评论内容，发评论人,两表查询
	template := "Select Created,Updated,User_Name,Content from Comment,User Where Post_Id=? and Comment.User_Id=User.User_Id"
	rows, err := DB.DB().Query(template,pid)
	if err != nil {
		log.Print(err)
	}
	allcomment := []Comment_view{}
	for rows.Next() {
		comment_view := Comment_view{}
		err = rows.Scan(&comment_view.Created, &comment_view.Updated,  &comment_view.Username,&comment_view.Content)
		if err != nil {
			log.Print(err)
		}
		allcomment = append(allcomment, comment_view)
	}
	return allcomment, nil
}

//删除评论
func DeleteComment(Comment_id string)error {
	template := "DELETE From Comment Where Comment_Id=?"
	_, err := DB.DB().Query(template, Comment_id)
	if err != nil {
		log.Print(err)
		return err
	}
	return nil
}