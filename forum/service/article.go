package service

import (
	"forum/dao"
	"forum/model"

	"github.com/jinzhu/gorm"
)

//AddArticleService 发表文章
func AddArticleService(article *model.Article) (aid int, err error) {
	return dao.AddArticle(article)
}

//GetArticleByIDServie 通过aid获取文章
func GetArticleByIDServie(aid int) (*model.Article, error) {

	return dao.GetArticleByID(aid)
}
//GetArticleByUserIDService 获取作者的文章
func GetArticleByUserIDService(uid int,pi int,ps int)([]model.Article,int,error){
	return dao.GetArticleByUserID(uid,pi,ps)
}
//DeleteArticleByID 通过id删除文章
func DeleteArticleByID(id int) (bool, error) {
	return dao.DeleteArticleByID(id)
}

//UpdateArticleService 更新文章
func UpdateArticleService(article *model.Article) (bool, error) {
	return dao.UpdateArticle(article)
}

//LikenumUpService 文章点赞数增加
func LikenumUpService(aid int,uid int) (bool, error) {
	return dao.ArticleLikenumUp(aid,uid)
}

//LikenumDownService 文章点赞数减少
func LikenumDownService(aid int,uid int) (bool, error) {
	article, err := dao.GetArticleByID(aid)
	if err != nil {
		return false, err
	}
	if article.Likenum == 0 {
		return true, nil
	}
	return dao.ArticleLikenumDown(aid,uid)
}

//GetArticleListService 默认方式获取文章列表
func GetArticleListService(pi int, ps int) ([]model.Article, int, error) {
	return dao.GetArticleByLimit(pi, ps)
}

//GetArticleByTimeService 按最新发表时间获得列表
func GetArticleByTimeService(pi int, ps int) ([]model.Article, int, error) {
	return dao.GetArticleByTime(pi, ps)
}

//GetArticleByComService 按评论数获得列表
func GetArticleByComService(pi int, ps int) ([]model.Article, int, error) {
	return dao.GetArticleByComnum(pi, ps)
}

//GetCateArticleByLikeService 按点赞数获得列表
func GetCateArticleByLikeService(cid int,pi int, ps int) ([]model.Article, int, error) {
	return dao.GetCateArticleByLikenum(cid,pi, ps)
}
//GetCateArticleListService 默认方式获取文章列表
func GetCateArticleListService(cid int,pi int, ps int) ([]model.Article, int, error) {
	return dao.GetCateArticleByLimit(cid,pi, ps)
}

//GetCateArticleByTimeService 按最新发表时间获得列表
func GetCateArticleByTimeService(cid int,pi int, ps int) ([]model.Article, int, error) {
	return dao.GetCateArticleByTime(cid,pi, ps)
}

//GetCateArticleByComService 按评论数获得列表
func GetCateArticleByComService(cid int,pi int, ps int) ([]model.Article, int, error) {
	return dao.GetCateArticleByComnum(cid,pi, ps)
}

//GetArticleByLikeService 按点赞数获得列表
func GetArticleByLikeService(pi int, ps int) ([]model.Article, int, error) {
	return dao.GetArticleByLikenum(pi, ps)
}
//AddFocusArticleService 添加关注
func AddFocusArticleService(uid int,aid int) (bool,error) {
	userarticle:=model.UserFocus{UserID: uid,ArticleID:aid}
	return dao.AddFocusArticle(userarticle)
}
//DeleteFocusArticleService 添加关注
func DeleteFocusArticleService(uid int,aid int) (bool,error) {
	userarticle:=model.UserFocus{UserID: uid,ArticleID:aid}
	return dao.DeleteFocusArticle(userarticle)
}
//JudgeFocusService 判断关注
func JudgeFocusService(uid int,aid int)(bool,error){
	userarticle:=model.UserFocus{UserID: uid,ArticleID:aid}
	err:=dao.JudgeFocus(userarticle)
	if err!=nil{
		if err==gorm.ErrRecordNotFound{
			return false,nil
		}
		return false,err
	}
	return true,nil
}
//GetUserFocueArticleService 获取用户关注文章列表
func GetUserFocueArticleService(uid int, pageindex int, pagesize int) ( []model.Article, int, error){
	alist,err:=dao.GetUserFocueArticle(uid, pageindex, pagesize)
	if err!=nil{
		return nil,-1,err
	}
	num:=len(alist)
	return alist,num,nil
}
//GetUserLikeArticleService 获取用户点赞列表
func GetUserLikeArticleService(uid int)([]int,error){
	return dao.GetUserLikeArticle(uid)
}