package utils

//一些信息代码

const (
	SuperUserId     = 1   //管理员UserId
	NotSuperUser    = 101 //不是管理员
	BannedSuccess   = 102 //封禁成功
	BannedDefault   = 103 //封禁失败
	BindDefault     = 200 //数据解析错误
	DatabaseDefault = 300 //数据库发生错误
)

//登录注册相关的代码
const (
	LoginSuccess      = 1000 //登录成功
	PasswordWrong     = 1001 //密码错误
	UserNameNotExists = 1002 //用户名不存在，登录时会出现
	UserIdNotFound    = 1003 //用户ID未找到
	RegisterSuccess   = 1010 //注册成功
	UserNameIsExists  = 1011 //用户名已存在，注册时会出现
	UserIconNotFound  = 1021 //用户头像查找错误
)

//Post相关代码
const (
	NotLogin          = 2001 //用户未登录
	PostNotExists     = 2002 //帖子不存在
	CreatePostSuccess = 2010 //创建Post成功
	CreatePostDefault = 2011 //创建Post失败
	RevisePostSuccess = 2020 //修改Post成功
	DeletePostSuccess = 2030 //删除Post成功
	DeletePostDefault = 2031 //删除Post失败
	CreateFeedBackSuccess=2040 //创建反馈失败
	CreateFeedBackDefault=2041 //创建反馈成功
)

//Comment相关代码
const (
	CreateCommentSuccess = 3010 //创建Comment成功
	ReviseCommentSuccess = 3020 //修改Comment成功
	DeleteCommentSuccess = 3030 //删除Comment成功
	DeleteCommentDefault = 3031 //删除Comment失败
)
