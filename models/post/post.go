package post

import (
	"bytes"
	"database/sql"
	"log"
	Tools "main/utils"
)

type Post struct {
	Id      int    `json:"PostId"`
	Uid     int    `json:"UserId"`
	Content string `json:"Content"`
}

//创建树洞
func CreatePost(post Post, db *sql.DB) (int64, error) {
	template := "Insert Post Set Created=?,User_Id=?,Updated=? Content=?"
	stmt, err := db.Prepare(template)
	if err != nil {
		log.Print(err)
	}
	created := Tools.GetDatetime()
	updated := Tools.GetDatetime()
	result, err := stmt.Exec(created, post.Uid, updated, post.Content)
	if err != nil {
		log.Print(err)
	}
	id, _ := result.LastInsertId()
	return id, err
}

//查看树洞，采用分页查询
func ViewPost(db *sql.DB) (string, error) {
	template := "Select Created,Updated,Content,User_Name From Post,User where Post.User_Id=User.User_Id"
	rows, err := db.Query(template)
	if err != nil {
		log.Print(err)
	}
	var created string
	var updated string
	var content string
	var username string
	var buffer bytes.Buffer
        Isfirst := true 
	buffer.WriteString("[")
        for rows.Next() {
		err = rows.Scan(&created, &updated, &content, &username)
		if err != nil {
			log.Print(err)
		}
                if !Isfirst {
                  buffer.WriteString(",")  
                }
                Isfirst = false
		buffer.WriteString(`{"created":`)
		buffer.WriteString(`"`)
		buffer.WriteString(created)
		buffer.WriteString(`"`)
		// buffer.WriteString("\"")
		buffer.WriteString(",")
		buffer.WriteString(`"updated":`)
		// buffer.WriteString("\"")
		buffer.WriteString(`"`)
		buffer.WriteString(updated)
		buffer.WriteString(`"`)
		// buffer.WriteString("\"")
		buffer.WriteString(",")
		buffer.WriteString(`"content":`)
		// buffer.WriteString("\"")
		buffer.WriteString(`"`)
		buffer.WriteString(content)
		buffer.WriteString(`"`)
		// buffer.WriteString("\"")
		buffer.WriteString(",")
		buffer.WriteString(`"username":`)
		// buffer.WriteString("\"")
		buffer.WriteString(`"`)
		buffer.WriteString(username)
		buffer.WriteString(`"`)
		// buffer.WriteString("\"")
		buffer.WriteString("}")
	}
        buffer.WriteString("]")
	return buffer.String(), nil
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
