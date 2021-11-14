package post

import (
	"database/sql"
	"log"
	Tools "main/utils"
)

type Post struct {
	Id      int    `json:Post_Id`
	Uid     int    `json:User_Id`
	Content string `json:Content`
}

//创建树洞
func CreatePost(post Post, db *sql.DB) (int64, bool) {
	template := "Insert Post Set Created=?,User_Id=?,Updated=? Content=?"
	stmt, err := db.Prepare(template)
	if err != nil {
		log.Print(err)
		return -1, false
	}
	created := Tools.GetDatetime()
	updated := Tools.GetDatetime()
	result, err := stmt.Exec(created, post.Uid, updated, post.Content)
	if err != nil {
		log.Print(err)
		return -1, false
	}
	id, _ := result.LastInsertId()
	return id, true
}

//查看树洞
func ViewPost(User_Id int, db *sql.DB) (string, string, string) {
	template := "Select Created,Updated,Content From Post Where User_Id=?"
	rows, err := db.Query(template, User_Id)
	if err != nil {
		log.Print(err)
	}
	var created string
	var updated string
	var content string
	rows.Next()
	err = rows.Scan(&created, &updated, &content)
	if err != nil {
		log.Print(err)
	}
	return created, updated, content
}

//删除树洞
func DeletePost(post_id string, db *sql.DB) error {
	template := "DELETE From Post Where Post_id=?"
	_, err := db.Query(template, post_id)
	if err != nil {
		log.Print(err)
		return err
	}
	return nil
}
