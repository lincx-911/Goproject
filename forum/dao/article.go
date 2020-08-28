package dao

import (
	"fmt"
	"forum/model"
	orm "forum/respository"

	"github.com/jinzhu/gorm"
)

//GetArticleByID 通过id获得文章
func GetArticleByID(aid int) (*model.Article, error) {
	article := &model.Article{}
	find := orm.DB.First(&article, aid)
	if find.RowsAffected == 0 {
		return article, nil
	}
	err:=orm.DB.Model(&article).Related(&article.Author,"author").Related(&article.Category,"category").Related(&article.Tag,"tag").Error
	return article, err
}
//GetArticleByUserID 通过作者id获取文章
func GetArticleByUserID(uid int,pageindex int,pagesize int)(articles []model.Article,count int,err error){
	article:=&model.Article{UserID: uid}
	articledb:=orm.DB.Model(article).Where("user_id=?",uid).Preload("Author").Preload("Category").Preload("Tag")
	articledb.Count(&count)//总行数
	err=articledb.Offset((pageindex-1)*pagesize).Limit(pagesize).Find(&articles).Error
	return
}
//GetUserFocueArticle 获取作者关注的文章列表
func GetUserFocueArticle(uid int,pageindex int,pagesize int)(articles []model.Article,err error)  {
	sub := orm.DB.Table("user_focus").Select("article_id").Where("user_id = ?", uid).SubQuery()
	article := &model.Article{}
	articledb:=orm.DB.Model(article).Where("id IN ?",sub).Preload("Author").Preload("Category").Preload("Tag")
	err=articledb.Offset((pageindex-1)*pagesize).Limit(pagesize).Find(&articles).Error
	return
}

//GetArticleByLimit 分页查询文章
func GetArticleByLimit(pageindex int,pagesize int) (articles []model.Article,count int,err error) {
	article := &model.Article{}
	articledb:=orm.DB.Model(article).Preload("Author").Preload("Category").Preload("Tag")
	articledb.Count(&count)//总行数
	err=articledb.Offset((pageindex-1)*pagesize).Limit(pagesize).Find(&articles).Error
	return
}
//GetArticleByComnum 按评论数排序文章
func GetArticleByComnum(pageindex int,pagesize int) (articles []model.Article,count int,err error) {
	article := &model.Article{}
	articledb:=orm.DB.Model(article).Preload("Author").Preload("Category").Preload("Tag")
	articledb.Count(&count)//总行数
	err=articledb.Offset((pageindex-1)*pagesize).Order("commentnum DESC").Limit(pagesize).Find(&articles).Error
	return
}
//GetArticleByTime 按时间排序
func GetArticleByTime(pageindex int,pagesize int) (articles []model.Article,count int,err error) {
	article := &model.Article{}
	articledb:=orm.DB.Model(article).Preload("Author").Preload("Category").Preload("Tag")
	articledb.Count(&count)//总行数
	err=articledb.Offset((pageindex-1)*pagesize).Order("id DESC").Limit(pagesize).Find(&articles).Error
	return
}
//GetArticleByLikenum 按点赞数排序文章
func GetArticleByLikenum(pageindex int,pagesize int) (articles []model.Article,count int,err error) {
	article := &model.Article{}
	articledb:=orm.DB.Model(article).Preload("Author").Preload("Category").Preload("Tag")
	articledb.Count(&count)//总行数
	err=articledb.Offset((pageindex-1)*pagesize).Order("likenum DESC").Limit(pagesize).Find(&articles).Error
	return
}
//GetCateArticleByLimit 分页查询文章
func GetCateArticleByLimit(cid int,pageindex int,pagesize int) (articles []model.Article,count int,err error) {
	article := &model.Article{}
	articledb:=orm.DB.Model(article).Where("category_id=?",cid).Preload("Author").Preload("Category").Preload("Tag")
	articledb.Count(&count)//总行数
	err=articledb.Offset((pageindex-1)*pagesize).Limit(pagesize).Find(&articles).Error
	return
}
//GetCateArticleByComnum 按评论数排序文章
func GetCateArticleByComnum(cid int,pageindex int,pagesize int) (articles []model.Article,count int,err error) {
	article := &model.Article{}
	articledb:=orm.DB.Model(article).Where("category_id=?",cid).Preload("Author").Preload("Category").Preload("Tag")
	articledb.Count(&count)//总行数
	err=articledb.Offset((pageindex-1)*pagesize).Order("commentnum DESC").Limit(pagesize).Find(&articles).Error
	return
}
//GetCateArticleByTime 按时间排序
func GetCateArticleByTime(cid int,pageindex int,pagesize int) (articles []model.Article,count int,err error) {
	article := &model.Article{}
	articledb:=orm.DB.Model(article).Where("category_id=?",cid).Preload("Author").Preload("Category").Preload("Tag")
	articledb.Count(&count)//总行数
	err=articledb.Offset((pageindex-1)*pagesize).Order("id DESC").Limit(pagesize).Find(&articles).Error
	return
}
//GetCateArticleByLikenum 按点赞数排序文章
func GetCateArticleByLikenum(cid int,pageindex int,pagesize int) (articles []model.Article,count int,err error) {
	article := &model.Article{}
	articledb:=orm.DB.Model(article).Where("category_id=?",cid).Preload("Author").Preload("Category").Preload("Tag")
	articledb.Count(&count)//总行数
	err=articledb.Offset((pageindex-1)*pagesize).Order("likenum DESC").Limit(pagesize).Find(&articles).Error
	return
}
//AddArticle 存入文章
func AddArticle(article *model.Article)(int,error){
	tx := orm.DB.Begin()
	var tags []model.Atag
	atag := article.Tag
	for i:=range atag{
		tag:=model.Atag{}
		if tx.Where("atagname=?",atag[i].Atagname).First(&tag).RecordNotFound(){
			tag = atag[i]
			if err:=tx.Create(&tag).Error;err!=nil{
				tx.Rollback()
				return -1,err
			}
		}
		tags = append(tags,tag)
	}
	article.Tag=tags
	err:=tx.Create(&article).Error
	if err!=nil{
		tx.Rollback()
		return -1,err
	}
	err = tx.Model(&article).Related(&article.Author,"author").Related(&article.Category,"category").Related(&article.Tag,"tag").Error
	if err!=nil{
		tx.Rollback()
		return -1,err
	}
	tx.Commit()
	return article.ID,nil
}

//UpdateArticle 更新文章
func UpdateArticle(article *model.Article)(bool,error){
	tx := orm.DB.Begin()
	var tags []model.Atag
	atag := article.Tag
	for i:=range atag{
		tag:=model.Atag{}
		if tx.Where("atagname=?",atag[i].Atagname).First(&tag).RecordNotFound(){
			tag = atag[i]
			if err:=tx.Create(&tag).Error;err!=nil{
				tx.Rollback()
				return false,err
			}
		}
		tags = append(tags,tag)
	}
	article.Tag=tags
	err:=tx.Model(&article).Association("tag").Replace(article.Tag).Error
	if err!=nil{
		tx.Rollback()
		return false,err
	}
	err =tx.Model(&article).Update(&article).Error
	if err!=nil{
		tx.Rollback()
		return false,err
	}
	tx.Commit()
	return true,nil
}
//DeleteArticleByID 通过id删除文章
func DeleteArticleByID(id int)(bool,error)  {
	tx:=orm.DB.Begin()
	var article model.Article
	if tx.First(&article,id).RecordNotFound(){
		tx.Rollback()
		return false,gorm.ErrRecordNotFound
	}
	if err:=tx.Model(&article).Association("tag").Clear().Error;err!=nil{
		tx.Rollback()
		return false,err
	}
	if err:=tx.Model(&article).Delete(&article).Error;err!=nil{
		tx.Rollback()
		return false,err
	}
	tx.Commit()
	return true,nil
}
//ArticleLikenumUp 点赞数+1
func ArticleLikenumUp(aid int,uid int) (bool,error) {
	var article model.Article
	tx:=orm.DB.Begin()
	err:=tx.Model(&article).Where("id=?",aid).Update("likenum",gorm.Expr("likenum + ?",1)).Error
	if err!=nil{
		tx.Rollback()
		return false,err
	}
	userlike:=model.UserLike{UserID: uid,ArticleID: aid}
	re:=tx.Find(&userlike).RowsAffected
	if re!=0{
		tx.Rollback()
		return true,nil
	}
	err=tx.Create(&userlike).Error
	if err!=nil{
		tx.Rollback()
		return false,err
	}
	tx.Commit()
	return true,nil
}
//ArticleLikenumDown 点赞数-1
func ArticleLikenumDown(aid int,uid int) (bool,error) {
	var article model.Article
	tx:=orm.DB.Begin()
	err:=tx.Model(&article).Where("id=?",aid).Update("likenum",gorm.Expr("likenum - ?",1)).Error
	if err!=nil{
		tx.Rollback()
		return false,err
	}
	userlike:=model.UserLike{UserID: uid,ArticleID: aid}
	re:=tx.Find(&userlike).RowsAffected
	if re==0{
		tx.Rollback()
		return true,nil
	}
	err=tx.Delete(&userlike).Error
	if err!=nil{
		tx.Rollback()
		return false,err
	}
	tx.Commit()
	fmt.Println("articlelikenum",article.Likenum)
	return true,nil
}
//UpdateArticleCommentNum 更新文章评论数
func UpdateArticleCommentNum(id int,num int)(error){
	return orm.DB.Model(&model.Article{}).Where("id = ?",id).Update("commentnum",num).Error
}
//ArticleCommentNum 文章评论数统计
func ArticleCommentNum(raid int)(int,error){
	var num int
	err:=orm.DB.Model(&model.Comment{}).Where("raid = ?",raid).Count(&num).Error
	return num,err
}

//AddFocusArticle 添加文章关注
func AddFocusArticle(ua model.UserFocus)(bool,error){
	err:=orm.DB.Create(&ua).Error
	if err!=nil{
		return false,err
	}
	return true,nil
}
//DeleteFocusArticle 取消关注
func DeleteFocusArticle(ua model.UserFocus)(bool,error){
	err:= orm.DB.First(&ua).Error
	if err!=nil{
		return false,err
	}
	err=orm.DB.Delete(&ua).Error
	if err!=nil{
		return false,err
	}
	return true,nil
}
//JudgeFocus 判断是否关注
func JudgeFocus(ua model.UserFocus)(error){
	err:= orm.DB.First(&ua).Error
	if err!=nil{
		return err
	}
	return nil
}
//GetUserLikeArticle 获取用户点赞文章id列表
func GetUserLikeArticle(uid int) (alist []int,err error) {
	//err=orm.DB.Table("user_likes").Select("article_id").Where("user_id=?",uid).Find(&alist).Error
	err=orm.DB.Table("user_likes").Where("user_id=?",uid).Pluck("article_id",&alist).Order("article_id ASC").Error
	return
}
