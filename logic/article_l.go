package logic

import (
	"fmt"
	"math"

	"github.com/Curricane/blogger/dal/db"
	"github.com/Curricane/blogger/model"
)

func getCategoryIds(articleInfoList []*model.ArticleInfo) (ids []int64) {

LABEL:
	for _, article := range articleInfoList {
		categoryId := article.Category
		for _, id := range ids { //去掉重复的分类
			if id == categoryId {
				continue LABEL
			}
		}

		ids = append(ids, categoryId)
	}
	return
}

//获取文章 + 类型
func GetArticleRecordList(pageNum, pageSize int) (articleRecordList []*model.ArticleRecord, err error) {
	//1,从数据库中获取数据
	articleInfoList, err := db.GetArtcleList(pageNum, pageSize)
	if err != nil {
		fmt.Printf("get artticle lsit failed, err:%v\n", err)
		return
	}
	if len(articleInfoList) == 0 {
		return
	}

	//2,获取文章对应的分类列表
	//为什么一次性获取全部文章集合，为了较少每个文章都通过数据库查询类型，减小数据库的访问？
	//是否可以用数据的联结查询获取到呢？
	categoryIds := getCategoryIds(articleInfoList)
	categoryList, err := db.GetCategoryList(categoryIds)
	if err != nil {
		fmt.Printf("get categoryList lsit failed, err:%v\n", err)
		return
	}

	//3，聚合数据
	for _, article := range articleInfoList {
		articleRecord := &model.ArticleRecord{
			ArticleInfo: *article,
		}
		categoryId := article.Category
		for _, category := range categoryList {
			if categoryId == category.CategoryId {
				articleRecord.Category = *category
				break
			}
		}
		articleRecordList = append(articleRecordList, articleRecord)
	}
	return
}

func GetArticleRecordListById(categoryId, pageNum, pageSize int) (articleRecordList []*model.ArticleRecord, err error) {
	//1,从数据库中获取数据
	articleInfoList, err := db.GetArtcleListByCategoryId(categoryId, pageNum, pageSize)
	if err != nil {
		fmt.Printf("get artticle lsit failed, err:%v\n", err)
		return
	}
	if len(articleInfoList) == 0 {
		return
	}

	//2,获取文章对应的分类列表
	//为什么一次性获取全部文章集合，为了较少每个文章都通过数据库查询类型，减小数据库的访问？
	//是否可以用数据的联结查询获取到呢？
	categoryIds := getCategoryIds(articleInfoList)
	categoryList, err := db.GetCategoryList(categoryIds)
	if err != nil {
		fmt.Printf("get categoryList lsit failed, err:%v\n", err)
		return
	}

	//3，聚合数据
	for _, article := range articleInfoList {
		articleRecord := &model.ArticleRecord{
			ArticleInfo: *article,
		}
		categoryId := article.Category
		for _, category := range categoryList {
			if categoryId == category.CategoryId {
				articleRecord.Category = *category
				break
			}
		}
		articleRecordList = append(articleRecordList, articleRecord)
	}
	return
}

func GetArticleDetail(articleId int64) (articleDetail *model.ArticleDetail, err error) {
	articleDetail, err = db.GetArticleDetail(articleId)
	if err != nil {
		return
	}

	category, err := db.GetCategoryById(articleDetail.ArticleInfo.Category)
	if err != nil {
		return
	}

	articleDetail.Category = *category
	return
}

func GetRelativeAricleList(articleId int64) (articleList []*model.RelativeArticle, err error) {

	articleList, err = db.GetRelativeArticle(articleId)
	return
}

func GetPrevAndNextArticleInfo(articleId int64) (prevArticle *model.RelativeArticle,
	nextArticle *model.RelativeArticle, err error) {

	prevArticle, err = db.GetPrevArticleById(articleId)
	if err != nil {
		//打印一个警告日志
	}

	nextArticle, err = db.GetNextArticleById(articleId)
	if err != nil {
		//打印一个警告日志
	}

	return
}

func InsertArticle(content, author, title string, categoryId int64) (err error) {

	//1. 构造一个ArticleDetail的类
	articleDetail := &model.ArticleDetail{}
	articleDetail.Content = content
	articleDetail.Username = author
	articleDetail.Title = title
	articleDetail.ArticleInfo.Category = categoryId

	contentUtf8 := []rune(content)
	minLength := int(math.Min(float64(len(contentUtf8)), 128.0))
	articleDetail.Summary = string([]rune(content)[:minLength])

	id, err := db.InsertArticle(articleDetail)
	fmt.Printf("insert article succ, id:%d, err:%v\n", id, err)
	return
}
