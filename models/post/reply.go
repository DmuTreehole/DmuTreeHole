package post

import (
	"main/db"
	Tools "main/utils"
)

type Reply struct {
	replyId    int    `json:"replyId"`
	userId     int    `json:"userId"`
	commentId  int    `json:"commentId"`
	replyProId int    `json:"replyProId"`
	created    string `json:"created"`
	updated    string `json:"updated"`
	content    string `json:"content"`
}

//创建回复
func CreateReply(reply Reply) bool {
	template := "Insert Into Reply Values User_Id=?,Comment_Id=?," +
		"Reply_proId=?,Created=?,Updated=?,Content=?"
	reply.created = Tools.GetDatetime()
	reply.updated = reply.created
	_, err := db.DB().Query(template, reply.userId, reply.commentId,
		reply.replyProId, reply.created, reply.updated, reply.created)
	return err == nil
}

//删除回复
func DeleteReply(reply Reply) bool {
	template := "Delete From Reply Where Reply_Id=?"
	_, err := db.DB().Query(template, reply.replyId)
	return err == nil
}

//查看回复
func ShowReply(reply Reply, page int) (allReply []Reply) {
	template := "Select * From Reply Where Comment_Id = ?,Reply_proId = ? Limit 5 Offset ? Order by Updated Desc"
	rows, err := db.DB().Query(template, reply.commentId, reply.replyProId, (page-1)*5)
	if err != nil {
		return nil
	}
	for rows.Next() {
		var newReply Reply
		rows.Scan()
		rows.Scan(&newReply.replyId, &newReply.userId, &newReply.commentId,
			&newReply.replyProId, &newReply.created, &reply.updated, &reply.content)
		allReply = append(allReply, newReply)
	}
	return
}
