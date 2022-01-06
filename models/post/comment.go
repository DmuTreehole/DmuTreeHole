package post

import (
	"log"
	DB "main/db"
	Tools "main/utils"
)

type Comment struct {
	Pid     int    `json:"PostId"`
	Uid     int    `json:"UserId"`
	Content string `json:"Content"`
}
type Comment_view struct {
	Id       int    `json:"CommentId"`
	Uid      int    `json:"UserId"`
	Pid      int    `json:"PostId"`
	Username string `json:"Username"`
	Updated  string
	Created  string
	Content  string `json:"Content"`
}

//创建评论
func CreateComment(comment Comment) error {
	template := "Insert Comment Set Created=?,User_Id=?,Post_Id=?,Updated=?,Content=?"
	stmt, err := DB.DB().Prepare(template)
	defer stmt.Close()
	if err != nil {
		return err
	}
	created := Tools.GetDatetime()
	updated := created
	_, err = stmt.Exec(created, comment.Uid, comment.Pid, updated, comment.Content)
	if err != nil {
		return err
	}
	return nil
}

//查看评论
func ShowComment(comment Comment) ([]Comment_view, error) {
	//创建更改时间，评论内容，发评论人,两表查询
	template := "Select Created,Updated,User_Name,Content,Comment_Id,User.User_Id from Comment,User Where Post_Id=? And Comment.User_Id=User.User_Id And isDelete != 'false'"
	rows, err := DB.DB().Query(template, comment.Pid)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	allcomment := []Comment_view{}
	for rows.Next() {
		comment_view := Comment_view{}
		err = rows.Scan(&comment_view.Created, &comment_view.Updated, &comment_view.Username, &comment_view.Content, &comment_view.Id, &comment_view.Uid)
		if err != nil {
			return nil, err
		}
		comment_view.Content = Tools.Fuck(comment_view.Content)
		allcomment = append(allcomment, comment_view)
	}
	return allcomment, nil
}

//删除评论
func DeleteComment(commentId int) error {
	//template := "DELETE From Comment Where Comment_Id=?"
	template := `Update Comment Set Set isDelete = 'true', Etc = 'User Delete' Where Comment_Id = ?`
	rows, err := DB.DB().Query(template, commentId)
	rows.Close()
	if err != nil {
		log.Print(err)
		return err
	}
	return nil
}
