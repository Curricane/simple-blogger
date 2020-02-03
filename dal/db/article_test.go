package db

import (
	"fmt"
	"testing"
)

func init() {
	//需要parseTime=true 才会把mysql中的time转化为go中的time
	dns := "root:cmc123456@tcp(localhost:3306)/blogger?parseTime=true"
	err := Init(dns)
	if err != nil {
		panic(err)
	}
}

// func TestInsertArticel(t *testing.T) {

// 	article := &model.ArticleDetail{}
// 	article.CategoryId = 1
// 	article.CommentCount = 0
// 	article.CreateTime = time.Now()
// 	article.Summary = `go中使用结构体嵌套来扩展类型
// 	嵌入到结构体中的字段，完全可以当作自己是自己的字段`
// 	article.Title = "结构体嵌套"
// 	article.Username = "Curricane"
// 	article.Category.CategoryId = 1
// 	article.ViewCount = 1
// 	article.Content = `匿名结构体就是在嵌入时，不指定名称，这样子会将匿名结果体的所有方法引入到该类型中；这样在使用时有很多便利`
// 	articleId, err := InsertArticle(article)
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	t.Log(articleId, err)
// }

// func TestGetArtcleList(t *testing.T) {

// 	result, err := GetArtcleList(7, 16)
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	fmt.Println("result is:", result, "len is:", len(result))
// }

func TestGetArtcleListByUsername(t *testing.T) {

	result, err := GetArtcleListByUsername("Curricane", 0, 16)
	if err != nil {
		fmt.Println("error")
		t.Error(err)
	}
	fmt.Println("result is:", result, "len is:", len(result))
}
