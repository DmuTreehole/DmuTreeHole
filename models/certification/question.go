package certification

import (
	DB "main/db"
	"math/rand"
)

type Question struct {
	questionId int    `json:"question_id"`
	answerA    string `json:"question_A"`
	answerB    string `json:"question_B"`
	answerC    string `json:"question_C"`
	answerD    string `json:"question_D"`
	correct    string `json:"question_correct"`
}

func Getonequestion() (Question, error) {
	que := Question{}
	var num int
	template := "select question_Id from Questions Order by question_Id"
	stmt, _ := DB.DB().Query(template)
	stmt.Scan(&num)
	que.questionId = rand.Intn(num-1) + 1
	template = "select answer1,answer2,answer3,answer4,correct from Questions Where question_Id = ?"
	stmt, err := DB.DB().Query(template, que.questionId)
	if err != nil {
		return que, err
	}
	stmt.Next()
	err = stmt.Scan(&que.answerA, &que.answerB, &que.answerC, &que.answerD, &que.correct)
	if err != nil {
		return que, err
	}
	return que, nil
}
