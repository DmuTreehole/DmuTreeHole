package post

import (
	"main/db"
	Tools "main/utils"
)

type Reply struct {
	ReplyId    int    `json:"replyId"`
	UserId     int    `json:"userId"`
	CommentId  int    `json:"commentId"`
	ReplyProId int    `json:"replyProId"`
	Created    string `json:"created"`
	Updated    string `json:"updated"`
	Content    string `json:"content"`
}

//创建回复
func CreateReply(reply Reply) bool {
	template := "Insert Into Reply Values User_Id=?,Comment_Id=?," +
		"Reply_proId=?,Created=?,Updated=?,Content=?"
	reply.Created = Tools.GetDatetime()
	reply.Updated = reply.Created
	_, err := db.DB().Query(template, reply.UserId, reply.CommentId,
		reply.ReplyProId, reply.Created, reply.Updated, reply.Created)
	return err == nil
}

//删除回复
func DeleteReply(reply Reply) bool {
	template := "Delete From Reply Where Reply_Id=?"
	_, err := db.DB().Query(template, reply.ReplyId)
	return err == nil
}

//查看回复
func ShowReply(reply Reply, page int) (allReply []Reply) {
	template := "Select * From Reply Where Comment_Id = ?,Reply_proId = ? Limit 5 Offset ? Order by Updated Desc"
	rows, err := db.DB().Query(template, reply.CommentId, reply.ReplyProId, (page-1)*5)
	if err != nil {
		return nil
	}
	for rows.Next() {
		var newReply Reply
		rows.Scan()
		rows.Scan(&newReply.ReplyId, &newReply.UserId, &newReply.CommentId,
			&newReply.ReplyProId, &newReply.Created, &reply.Updated, &reply.Content)
		allReply = append(allReply, newReply)
	}
	return
}
