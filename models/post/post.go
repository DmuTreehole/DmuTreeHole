package post

import (
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
	Id   int `json:"UserId"`
	Page int `json:"Page"`
}

type view struct {
	Id       string `json:"PostId"`
	UserId   int    `json:"UserId"`
	Created  string `json:"created_time"`
	Updated  string `json:"updated_time"`
	Content  string `json:"Content"`
	Username string `json:"Username"`
}

//创建树洞
func CreatePost(post Post) (int64, error) {
	template := "Insert Post Set Created=?,User_Id=?,Updated=?,Content=?"
	stmt, err := DB.DB().Prepare(template)
	defer stmt.Close()
	if err != nil {
		log.Print(err)
		return -1, err
	}
	created := Tools.GetDatetime()
	updated := created
	result, err := stmt.Exec(created, post.Uid, updated, Tools.Fuck(post.Content))
	if err != nil {
		log.Print(err)
		return -1, err
	}
	id, _ := result.LastInsertId()
	return id, err
}

//查看树洞，采用分页查询,每次显示五条
func ViewPost(page int) ([]view, error) {
	template := "Select Post_Id,Created,Updated,Content,User_Name,User.User_Id From Post,User " +
		"where Post.User_Id=User.User_Id Order By Created Desc Limit 5 Offset ?"
	rows, err := DB.DB().Query(template, (page-1)*5) // page 从1开始
	defer rows.Close()
	if err != nil {
		log.Print(err)
		return nil, err
	}
	allpost := []view{}
	for rows.Next() {
		post := view{}
		err = rows.Scan(&post.Id, &post.Created, &post.Updated, &post.Content, &post.Username, &post.UserId)
		if err != nil {
			log.Print(err)
			return nil, err
		}
		allpost = append(allpost, post)
	}
	return allpost, nil
}

//删除树洞
func DeletePost(post_id int) error {
	template := "DELETE From Comment Where Post_Id=?"
	rows, err := DB.DB().Query(template, post_id)
	defer rows.Close()
	if err != nil {
		log.Print(err)
		return err
	}
	template = "DELETE From Post Where Post_Id=?"
	rows1, err := DB.DB().Query(template, post_id)
	rows1.Close()
	if err != nil {
		log.Print(err)
		return err
	}
	return nil
}

//按userid查帖子
func QueryPostById(info PagePost) ([]view, error) {
	template := "Select Post_Id,Created,Updated,Content,User_Name From Post,User " +
		"where Post.User_Id=User.User_Id And Post.User_Id = ? Order By Created Desc Limit 5 Offset ?"
	rows, err := DB.DB().Query(template, info.Id, (info.Page-1)*5)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	var allpost = []view{}
	var post = view{}
	for rows.Next() {
		err = rows.Scan(&post.Id, &post.Created, &post.Updated, &post.Content, &post.Username)
		if err != nil {
			return nil, err
		}
		allpost = append(allpost, post)
	}
	return allpost, nil
}
