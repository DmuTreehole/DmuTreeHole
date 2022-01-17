package post

import (
	"main/db"
	"strings"
)

type Content struct {
	Content string `json:"Content"`
	Page    int    `json:"Page"`
}

func Query(Content Content) ([]view, error) {
	Content.Content = strings.ToUpper(Content.Content)
	Content.Content = `%` + Content.Content + `%`
	template := "Select Post_Id,Created,Updated,Content,User_Name From Post,User" +
		" Where Post.User_Id = User.User_Id And Upper(Content) Like ? And isDelete <> 'false' Order By Created Desc"
	rows, err := db.DB().Query(template, Content.Content)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	var allpost = []view{}
	var post = view{}
	for rows.Next() {
		err = rows.Scan(&post.Id, &post.Created, &post.Updated, &post.Content, &post.Username)
		if err != nil {
			continue
		}
		allpost = append(allpost, post)
	}
	return allpost, nil
}
