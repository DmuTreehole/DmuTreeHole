package post

import (
	"fmt"
	"log"
	"main/db"
)

type Content struct {
	Content string `json:"Content"`
}

func Query(Content Content) ([]view, error) {
	// 减少使用 !=
	template := "Select Post_Id,Created,Updated,Content,User_Name From Post,User where Post.User_Id=User.User_Id And isDelete == 'false'And Match(Content) Against(?) Order By Created Desc"
	rows, err := db.DB().Query(template, Content.Content)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	var allPost = []view{}
	var post = view{}
	for rows.Next() {
		err = rows.Scan(&post.Id, &post.Created, &post.Updated, &post.Content, &post.Username)
		fmt.Println(post)
		if err != nil {
			log.Fatalln(err)
			continue
		}
		allPost = append(allPost, post)
	}
	return allPost, nil
}
