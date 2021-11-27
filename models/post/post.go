package post

import (
	"fmt"
	"log"
	DB "main/db"
	Tools "main/utils"
)

type Post struct {
	Id      int    `json:"PostId",form:"PostId"`
	Uid     int    `json:"UserId"`
	Content string `json:"Content",form:"Content"`
}

type PagePost struct {
	Page int `json:"page"`
}

type view struct {
	Id       string
	Created  string `json:"created_time"`
	Updated  string `json:"updated_time"`
	Content  string `json:"content"`
	Username string `json:"user"`
}

//创建树洞
func CreatePost(post Post) (int64, error) {
	template := "Insert Post Set Created=?,User_Id=?,Updated=?,Content=?"
	stmt, err := DB.DB().Prepare(template)
	if err != nil {
		log.Print(err)
	}
	created := Tools.GetDatetime()
	updated := created
	fmt.Println(post)
	result, err := stmt.Exec(created, post.Uid, updated, post.Content)
	if err != nil {
		log.Print(err)
	}
	id, _ := result.LastInsertId()
	return id, err
}

//查看树洞，采用分页查询,每次显示五条
func ViewPost(page int) ([]view, error) {
	template := "Select Post_Id,Created,Updated,Content,User_Name From Post,User " +
		"where Post.User_Id=User.User_Id Order By Created Desc Limit 5 Offset ?"
	rows, err := DB.DB().Query(template, (page-1)*5) // page 从1开始
	fmt.Println(template)
	if err != nil {
		log.Print(err)
	}
	allpost := []view{}
	for rows.Next() {
		post := view{}
		err = rows.Scan(&post.Id, &post.Created, &post.Updated, &post.Content, &post.Username)
		if err != nil {
			log.Print(err)
		}
		allpost = append(allpost, post)
	}
	return allpost, nil
}

//删除树洞
func DeletePost(post_id int) error {
	template := "DELETE From Post Where Post_id=?"
	_, err := DB.DB().Query(template, post_id)
	if err != nil {
		log.Print(err)
		return err
	}
	return nil
}
