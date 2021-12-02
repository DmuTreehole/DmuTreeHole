package certification

import (
	DB "main/db"
	"math/rand"
)

type Question struct {
	QuestionId int    `json:"QuestionIid"`
	QuestionName string `json:"QuestionName"`
	AnswerA    string `json:"AnswerA"`
	AnswerB    string `json:"AnswerB"`
	AnswerC    string `json:"AnswerC "`
	AnswerD    string `json:"AnswerD "`
	Correct    string `json:"Correct"`
}
// 拿到问题和答案
func Getonequestion() (Question, error) {
	que := Question{}
	var num int
	template := "select question_Id from Questions Order by question_Id"
	stmt, _ := DB.DB().Query(template)
	stmt.Scan(&num)
	que.QuestionId = rand.Intn(num-1) + 1
	template = "select question,answer1,answer2,answer3,answer4,correct from Questions Where question_Id = ?"
	stmt, err := DB.DB().Query(template, que.QuestionId)
	if err != nil {
		return que, err
	}
	stmt.Next()
	err = stmt.Scan(&que.QuestionName,&que.AnswerA, &que.AnswerB, &que.AnswerC, &que.AnswerD, &que.Correct)
	if err != nil {
		return que, err
	}
	return que, nil
}
