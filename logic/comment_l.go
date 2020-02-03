package logic

import (
	"fmt"
	"time"

	"github.com/Curricane/blogger/dal/db"
	"github.com/Curricane/blogger/model"
)

func InsertComment(comment, author, email string, articleId int64) (err error) {
	//检查id是否合法
	exist, err := db.IsArticleExist(articleId)
	if err != nil {
		fmt.Printf("query database failed, err: %v\n", err)
		return
	}

	if exist == false {
		err = fmt.Errorf("article id:%d not found\n", articleId)
		return
	}

	//评论
	var c model.Comment
	c.ArtcleId = articleId
	c.Content = comment
	c.Username = author
	c.CreateTime = time.Now()
	c.Status = 1

	err = db.InsertComment(&c)
	return
}

//缺少对评论的翻页
func GetCommentList(articleId int64) (commentList []*model.Comment, err error) {

	//1. 首先，要验证article_id是否合法
	exist, err := db.IsArticleExist(articleId)
	if err != nil {
		fmt.Printf("query database failed, err:%v\n", err)
		return
	}

	if exist == false {
		err = fmt.Errorf("article id:%d not found", articleId)
		return
	}

	//2. 调用dal GetCommentList获取评论列表
	commentList, err = db.GetCommentList(articleId, 0, 100)
	return
}
