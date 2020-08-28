package service

import (
	"forum/dao"
	"forum/model"
)

//AddCommentService 发表评论
func AddCommentService(comment *model.Comment)(*model.Comment,error){
	return dao.AddComment(comment)
}
//DeleteCommentByID 删除评论
func DeleteCommentByID(id int) (bool,error) {
	
	return dao.DeleteArticleByID(id)
}
//ComLikenumUpService 评论点赞数增加 
func ComLikenumUpService(aid int,uid int ) (bool,error) {
	return dao.CommentLikenumUp(aid,uid)
}

//ComLikenumDownService 评论点赞数减少
func ComLikenumDownService(aid int,uid int) (bool,error) {
	article,err:=dao.GetArticleByID(aid)
	if err!=nil{
		return false,err
	}
	if article.Likenum==0{
		return true,nil
	}
	return dao.CommentLikenumDown(aid,uid)
}
//GetCommentByRid 获取文章的所有评论
func GetCommentByRid(rid int,pageindex int,pagesize int) ( []model.Comment,int,error)  {
	return dao.GetCommentByLimit(rid,pageindex,pagesize)
}
//GetUserCommentByLimitService 获取用户发表的评论
func GetUserCommentByLimitService(uid int,pageindex int,pagesize int) ( []model.Comment,int,error) {
	return dao.GetUserCommentByLimit(uid,pageindex,pagesize)
}
//GetUserLikeCommentService 获取用户点赞列表
func GetUserLikeCommentService(uid int)([]int,error){
	return dao.GetUserLikeComment(uid)
}