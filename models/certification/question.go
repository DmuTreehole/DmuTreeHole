package certification

import (
	DB "main/db"
)

type Question struct {
	QuestionId int    `json:"QuestionId"`
	QuestionName string `json:"QuestionName"`
	AnswerA    string `json:"AnswerA"`
	AnswerB    string `json:"AnswerB"`
	AnswerC    string `json:"AnswerC"`
	AnswerD    string `json:"AnswerD"`
	Correct    string `json:"Correct"`
}
// 拿到问题和答案
func GetThreeQuestions() ([]Question, error) {
	quelist := []Question{}
	que:=Question{}
	template := "select question,answer1,answer2,answer3,answer4,correct from Questions order by rand() limit 3"
	rows, err := DB.DB().Query(template)
	if err != nil {
		return quelist, err
	}
	for rows.Next(){
		err = rows.Scan(&que.QuestionName,&que.AnswerA, &que.AnswerB, &que.AnswerC, &que.AnswerD, &que.Correct)
		quelist=append(quelist, que)
		if err != nil {
			return quelist, err
		}
	}
	return quelist, nil
}
