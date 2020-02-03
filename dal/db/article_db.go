package db

import (
	"database/sql"
	"fmt"

	"github.com/Curricane/blogger/model"
	_ "github.com/go-sql-driver/mysql"
)

func InsertArticle(article *model.ArticleDetail) (articleId int64, err error) {
	if article == nil {
		err = fmt.Errorf("invalid article parameter")
		return
	}
	sqlstr := "insert into article(content, summary, title, username, category_id)values(?,?,?,?,?)"
	result, err := DB.Exec(sqlstr, article.Content, article.Summary,
		article.Title, article.Username, article.Category.CategoryId)
	if err != nil {
		fmt.Println("数据库执行失败！")
		fmt.Println(err)
		return
	}

	articleId, err = result.LastInsertId()
	return
}

//在所有文章中，从第pageNum，开始最多还去pageSize篇文章
func GetArtcleList(pageNum, pageSize int) (articleList []*model.ArticleInfo, err error) {
	if pageNum < 0 || pageSize < 0 {
		err = fmt.Errorf("invalid parameter, page_num:%d, page_size:%d", pageNum, pageSize)
		return
	}
	sqlstr := `select 
					id, summary, title, view_count, create_time, comment_count, username
				from
					article
				where
					status = 1
				order by create_time desc
				limit ?, ?`
	err = DB.Select(&articleList, sqlstr, pageNum, pageSize)
	return
}

func GetArtcleListByUsername(username string, pageNum, pageSize int) (articleList []*model.ArticleInfo, err error) {
	if pageNum < 0 || pageSize < 0 || len(username) <= 0 {
		err = fmt.Errorf("invalid parameter, page_num:%d, page_size:%d, username:%s", pageNum, pageSize, username)
		return
	}
	sqlstr := `select 
					id, summary, title, view_count, create_time, comment_count, username
				from
					article
				where
					status = 1 and username = ?
				order by create_time desc
				limit ?, ?`
	err = DB.Select(&articleList, sqlstr, username, pageNum, pageSize)
	return
}

func GetArtcleListByCategoryId(categoryId, pageNum, pageSize int) (articleList []*model.ArticleInfo, err error) {
	if pageNum < 0 || pageSize < 0 {
		err = fmt.Errorf("invalid parameter, page_num:%d, page_size:%d, categoryId:%d", pageNum, pageSize, categoryId)
		return
	}
	sqlstr := `select 
					id, summary, title, view_count, create_time, comment_count, username
				from
					article
				where
					status = 1 and category_id = ?
				order by create_time desc
				limit ?, ?`
	err = DB.Select(&articleList, sqlstr, categoryId, pageNum, pageSize)
	return
}

//获取文章所有信息
func GetArticleDetail(articleId int64) (articleDetail *model.ArticleDetail, err error) {
	if articleId < 0 {
		err = fmt.Errorf("invalid parameter, article_id:%d", articleId)
		return
	}
	articleDetail = &model.ArticleDetail{}
	sqlstr := `select
					id, summary, title, view_count, content,
					create_time, comment_count, username, category_id
				from
					article
				where
					id = ?
				and 
					status = 1				
				`
	err = DB.Get(articleDetail, sqlstr, articleId)
	fmt.Printf("articleDetail:%#v\n", articleDetail)
	return
}

//获取类型相关的文章
func GetRelativeArticle(articleId int64) (articleList []*model.RelativeArticle, err error) {
	var categoryId int64
	sqlstr := "select category_id from article where id=?"
	err = DB.Get(&categoryId, sqlstr, articleId)
	if err != nil {
		return
	}

	sqlstr = "select id, title from article where category_id=? and id != ? limit 10"
	err = DB.Select(&articleList, sqlstr, categoryId, articleId)
	return
}

//获取上一篇文章
func GetPrevArticleById(articleId int64) (info *model.RelativeArticle, err error) {
	info = &model.RelativeArticle{
		ArticleId: -1,
	}

	sqlstr := "select id, title from article where id < ? order by id desc limit 1"
	err = DB.Get(info, sqlstr, articleId)
	if err != nil {
		return
	}
	return
}

func GetNextArticleById(articleId int64) (info *model.RelativeArticle, err error) {
	info = &model.RelativeArticle{
		ArticleId: -1,
	}

	sqlstr := "select id, title from article where id > ? order by id asc limit 1"
	err = DB.Get(info, sqlstr, articleId)
	if err != nil {
		return
	}
	return
}

func IsArticleExist(articleId int64) (exists bool, err error) {
	var id int64
	sqlstr := "select id from article where id=?"
	err = DB.Get(&id, sqlstr, articleId)
	if err == sql.ErrNoRows {
		exists = false
		return
	}
	if err != nil {
		return
	}
	exists = true
	return
}

//删除文章
func DeleteArticle(articleId int64) (err error) {
	sqlstr := `update article set status=0 where id = ?`
	_, err = DB.Exec(sqlstr, articleId)
	return
}
