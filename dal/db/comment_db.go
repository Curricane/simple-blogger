package db

import (
	"fmt"

	"github.com/Curricane/blogger/model"
	_ "github.com/go-sql-driver/mysql"
)

func InsertComment(comment *model.Comment) (err error) {
	if comment == nil {
		err = fmt.Errorf("nil parameter")
		return
	}

	tx, err := DB.Beginx()
	if err != nil {
		return
	}

	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
	}()

	sqlstr := `insert into comment(
			content, username, article_id)
		values (
			?, ?, ?
		)
		`
	_, err = tx.Exec(sqlstr, comment.Content, comment.Username, comment.ArtcleId)
	if err != nil {
		return
	}
	sqlstr = `update article set comment_count = comment_count + 1 where id = ?`
	_, err = tx.Exec(sqlstr, comment.ArtcleId)
	if err != nil {
		return
	}

	err = tx.Commit()
	return
}

func UpdateViewCount(articleId int64) (err error) {
	sqlstr := `update article set comment_count = comment_count + 1 where id = ?`
	_, err = DB.Exec(sqlstr, articleId)
	return
}

func DeleteComment(commentId int64) (err error) {
	sqlstr := `update comment set status=0 where id = ?`
	_, err = DB.Exec(sqlstr, commentId)
	return
}

func GetCommentList(articleId int64, pageNum, pageSize int) (commentList []*model.Comment, err error) {
	if pageNum < 0 || pageSize < 0 {
		err = fmt.Errorf("invalid parameter, page_num:%d, page_size:%d", pageNum, pageSize)
		return
	}
	sqlstr := `select id, content, username, create_time 
				from comment
				where article_id = ? and status=1
				order by create_time desc
				limit ?, ?`
	err = DB.Select(&commentList, sqlstr, articleId, pageNum, pageSize)
	return
}
