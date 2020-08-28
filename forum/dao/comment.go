package dao

import (
	"fmt"
	"forum/model"
	orm "forum/respository"

	"github.com/jinzhu/gorm"
)

//GetCommetByID 通过id获得评论
func GetCommetByID(cid int) (*model.Comment, error) {
	comment := &model.Comment{}
	find := orm.DB.First(&comment, cid)
	if find.RowsAffected == 0 {
		return comment, nil
	}
	err:=orm.DB.Model(&comment).Related(&comment.Author,"author").Related(&comment.Touser,"touser").Error
	return comment, err
}
//GetCommentByLimit 分页查询评论
func GetCommentByLimit(rid int,pageindex int,pagesize int) (comments []model.Comment,count int,err error) {
	comment := &model.Comment{Raid: rid}
	commentdb:=orm.DB.Model(comment).Where("raid=?",rid).Preload("Author").Preload("Touser")
	commentdb.Count(&count)//总行数
	err=commentdb.Offset((pageindex-1)*pagesize).Limit(pagesize).Find(&comments).Error
	return
}
//GetUserCommentByLimit 分页查询用户发表的评论
func GetUserCommentByLimit(uid int,pageindex int,pagesize int) (comments []model.Comment,count int,err error) {
	comment := &model.Comment{UserID: uid}
	commentdb:=orm.DB.Model(comment).Where("user_id=?",uid).Preload("Author").Preload("Touser")
	commentdb.Count(&count)//总行数
	err=commentdb.Offset((pageindex-1)*pagesize).Limit(pagesize).Find(&comments).Error
	return
}
//GetCommentOrderByTime 按时间排序
func GetCommentOrderByTime(rid int,pageindex int,pagesize int) (comments []model.Comment,err error) {
	comment := &model.Comment{Raid: rid}
	commentdb:=orm.DB.Model(comment).Preload("Author")
	var count int
	commentdb.Count(&count)//总行数
	err=commentdb.Offset((pageindex-1)*pagesize).Order("id DESC").Limit(pagesize).Find(&comments).Error
	return
}
//GetCommentOrderByComnum 按评论数排序
func GetCommentOrderByComnum(rid int,pageindex int,pagesize int) (comments []model.Comment,err error) {
	comment := &model.Comment{}
	commentdb:=orm.DB.Model(comment).Preload("Author")
	var count int
	commentdb.Count(&count)//总行数
	err=commentdb.Offset((pageindex-1)*pagesize).Order("commentnum DESC").Limit(pagesize).Find(&comments).Error
	return
}
//GetCommentOrderByLikenum 按点赞数排序
func GetCommentOrderByLikenum(rid int,pageindex int,pagesize int) (comments []model.Comment,err error) {
	comment := &model.Comment{}
	commentdb:=orm.DB.Model(comment)
	var count int
	commentdb.Count(&count)//总行数
	err=commentdb.Offset((pageindex-1)*pagesize).Order("likenumnum DESC").Limit(pagesize).Related(&comment.Author,"author").Find(&comments).Error
	return
}
//AddComment 存入评论
func AddComment(comment *model.Comment)(*model.Comment,error){
	
	tx := orm.DB.Begin()
	err:=tx.Create(&comment).Error
	if err!=nil{
		tx.Rollback()
		return nil,err
	}
	if comment.Paid==0{
		err:=tx.Model(&model.Article{}).Where("id=?",comment.Raid).Update("commentnum",gorm.Expr("commentnum + ?",1)).Error
		if err!=nil{
			tx.Rollback()
			return nil,err
		}
	}else{
		err:=tx.Model(&model.Comment{}).Where("id=?",comment.Paid).Update("commentnum",gorm.Expr("commentnum + ?",1)).Error
		if err!=nil{
			tx.Rollback()
			return nil,err
		}
	}
	err = tx.Model(&comment).Related(&comment.Author,"author").Error
	if err!=nil{
		tx.Rollback()
		return nil,err
	}
	tx.Commit()
	return comment,nil
}
//DeleteCommentByID 通过id删除评论
func DeleteCommentByID(id int)(bool,error)  {
	tx:=orm.DB.Begin()
	var comment model.Comment
	if tx.First(&comment,id).RecordNotFound(){
		tx.Rollback()
		return false,gorm.ErrRecordNotFound
	}
	if err:=tx.Model(&comment).Delete(&comment).Error;err!=nil{
		tx.Rollback()
		return false,err
	}
	if comment.Paid==0{
		err:=tx.Model(&model.Article{}).Where("id=?",comment.Raid).Update("commentnum",gorm.Expr("commentnum - ?",1)).Error
		if err!=nil{
			tx.Rollback()
			return false,err
		}
	}else{
		err:=tx.Model(&model.Comment{}).Where("id=?",comment.Paid).Update("commentnum",gorm.Expr("commentnum - ?",1)).Error
		if err!=nil{
			tx.Rollback()
			return false,err
		}
	}
	tx.Commit()
	return true,nil
}

//CommentLikenumUp 点赞数+1
func CommentLikenumUp(cid int,uid int) (bool,error) {
	var comment model.Comment
	tx:=orm.DB.Begin()
	err:=tx.Model(&comment).Where("id=?",cid).Update("likenum",gorm.Expr("likenum + ?",1)).Error
	if err!=nil{
		tx.Rollback()
		return false,err
	}
	userclike:=model.UserClike{UserID: uid,CommentID: cid}
	re:=tx.Find(&userclike).RowsAffected
	if re!=0{
		tx.Rollback()
		return true,nil
	}
	err=tx.Create(&userclike).Error
	if err!=nil{
		tx.Rollback()
		return false,err
	}
	tx.Commit()
	return true,nil
}
//CommentLikenumDown 点赞数-1
func CommentLikenumDown(cid int,uid int) (bool,error) {
	var comment model.Comment
	tx:=orm.DB.Begin()
	err:=tx.Model(&comment).Where("id=?",cid).Update("likenum",gorm.Expr("likenum - ?",1)).Error
	if err!=nil{
		tx.Rollback()
		return false,err
	}
	userclike:=model.UserClike{UserID: uid,CommentID: cid}
	re:=tx.Find(&userclike).RowsAffected
	if re==0{
		tx.Rollback()
		return true,nil
	}
	err=tx.Delete(&userclike).Error
	if err!=nil{
		tx.Rollback()
		return false,err
	}
	tx.Commit()
	return true,nil
}
//CommentCountNum 评论数统计
func CommentCountNum(cid int)(int,error){
	var num int
	err:=orm.DB.Model(&model.Comment{}).Where("pid= ?",cid).Count(&num).Error
	return num,err
}
//UpdateCommentNum 更新评论数
func UpdateCommentNum(cid int,num int)(error){
	return orm.DB.Model(&model.Comment{}).Where("id = ?",cid).Update("commentnum",num).Error
}
//GetUserLikeComment 用户点赞评论列表
func GetUserLikeComment(uid int) (alist []int,err error) {
	err=orm.DB.Table("user_clikes").Where("user_id=?",uid).Pluck("comment_id",&alist).Order("comment_id ASC").Error
	fmt.Println("cuserlist",alist)
	return
}