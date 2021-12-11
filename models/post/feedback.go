package post

import (
	"log"
	DB "main/db"
	Tools "main/utils"
)

type Feedback struct {
	Uid     int    `json:"UserId"`
	Content string `json:"Content",form:"Content"`
}

func CreateFeedBack(feedback Feedback) (int64, error) {
	template := "Insert Feedback Set Created=?,User_Id=?,Updated=?,Content=?"
	stmt, err := DB.DB().Prepare(template)
	defer stmt.Close()
	if err != nil {
		log.Print(err)
	}
	created := Tools.GetDatetime()
	updated := created
	result, err := stmt.Exec(created, feedback.Uid, updated, feedback.Content)
	if err != nil {
		log.Print(err)
	}
	id, _ := result.LastInsertId()
	return id, err
}
