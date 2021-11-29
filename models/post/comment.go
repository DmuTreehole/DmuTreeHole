package post

import (
	"fmt"
	"log"
	DB "main/db"
	Tools "main/utils"
)

type Comment struct {
	Pid     int    `json:"PostId",form:"PostId"`
	Uid     int    `json:"UserId"`
	Content string `json:"Content",form:"Content"`
}
type Comment_view struct {
	Id       int `json:"Comment_Id"`
	Uid      int
	Pid      int    `json:"Post_Id"`
	Username string `json:"User_Id"`
	Updated  string
	Created  string
	Content  string `json:"Content"`
}

//创建评论
func CreateComment(comment Comment) (int64, error) {
	template := "Insert Comment Set Created=?,User_Id=?,Post_Id=?,Updated=?,Content=?"
	stmt, err := DB.DB().Prepare(template)
	if err != nil {
		log.Print(err)
	}
	created := Tools.GetDatetime()
	updated := created
	fmt.Println(comment)
	result, err := stmt.Exec(created, comment.Uid, comment.Pid, updated, comment.Content)
	if err != nil {
		log.Print(err)
	}
	id, _ := result.LastInsertId()
	return id, err
}

//查看评论
func ShowComment(pid int) ([]Comment_view, error) {
	//创建更改时间，评论内容，发评论人,两表查询
	template := "Select Created,Updated,User_Name,Content from Comment,User Where Post_Id=? and Comment.User_Id=User.User_Id"
	rows, err := DB.DB().Query(template, pid)
	if err != nil {
		log.Print(err)
	}
	allcomment := []Comment_view{}
	for rows.Next() {
		comment_view := Comment_view{}
		err = rows.Scan(&comment_view.Created, &comment_view.Updated, &comment_view.Username, &comment_view.Content)
		if err != nil {
			log.Print(err)
		}
		allcomment = append(allcomment, comment_view)
	}
	return allcomment, nil
}

//删除评论
func DeleteComment(commentId int) error {
	template := "DELETE From Comment Where Comment_Id=?"
	_, err := DB.DB().Query(template, commentId)
	if err != nil {
		log.Print(err)
		return err
	}
	return nil
}
