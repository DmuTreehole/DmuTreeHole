package post

import (
	"log"
	DB "main/db"
	Tools "main/utils"
)

type Post struct {
	Id      int    `json:"PostId"`
	Uid     int    `json:"UserId"`
	Content string `json:"Content"`
}

type view struct {
	Created string `json:"created_time"`
	Updated string `json:"updated_time"`
	Content string `json:"content"`
	User    string `json:"user"`
}

//创建树洞
func CreatePost(post Post) (int64, error) {
	template := "Insert Post Set Created=?,User_Id=?,Updated=? Content=?"
	stmt, err := DB.DB().Prepare(template)
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
func ViewPost() ([]view, error) {
	template := "Select Created,Updated,Content,User_Name From Post,User where Post.User_Id=User.User_Id"
	rows, err := DB.DB().Query(template)
	if err != nil {
		log.Print(err)
	}
	//var created string
	//var updated string
	//var content string
	//var username string
	//	var buffer bytes.Buffer
	//	Isfirst := true
	//buffer.WriteString("[")
	allpost := []view{}
	for rows.Next() {
		post := view{}
		err = rows.Scan(&post.Created, &post.Updated, &post.Content, &post.User)
		if err != nil {
			log.Print(err)
		}
		allpost = append(allpost, post)
		//allpost = append(allpost, view{Created: created, Updated: updated, Content: content, User: username})
		/*
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
		*/

	}
	//	buffer.WriteString("]")
	//js, err := json.Marshal(allpost)
	//if err != nil {
	//	log.Print(err)
	//}
	//return string(js), nil
	return allpost, nil
}

//删除树洞
func DeletePost(post_id string) error {
	template := "DELETE From Post Where Post_id=?"
	_, err := DB.DB().Query(template, post_id)
	if err != nil {
		log.Print(err)
		return err
	}
	return nil
}
