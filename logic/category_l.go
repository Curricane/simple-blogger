package logic

import (
	"fmt"

	"github.com/Curricane/blogger/dal/db"
	"github.com/Curricane/blogger/model"
)

func GetAllCategoryList() (categoryList []*model.Category, err error) {
	categoryList, err = db.GetAllCategoryList()
	if err != nil {
		fmt.Printf("get all category failed, err:%v\n", err)
		return
	}
	return
}
