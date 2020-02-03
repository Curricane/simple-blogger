package db

import (
	"testing"

	"github.com/Curricane/blogger/model"
)

func TestGetCategoryList(t *testing.T) {
	var categoryIds []int64
	categoryIds = append(categoryIds, 1)
	categoryList, err := GetCategoryList(categoryIds)
	if err != nil {
		t.Errorf("get category list failed, err:%v\n", err)
		return
	}

	if len(categoryIds) != len(categoryIds) {
		t.Errorf("get category list failed, len of categorylist:%d, ids len:%d\n",
			len(categoryList), len(categoryIds))
	}

	for _, v := range categoryList {
		t.Logf("id: %d category:%#v\n", v.CategoryId, v.CategoryName)
	}
}

func TestInsertCategory(t *testing.T) {
	category := &model.Category{}
	category.CategoryName = "golang"
	category.CategoryNo = 1

	id, err := InsertCategory(category)
	if err != nil {
		t.Errorf("failed to inset category %v", err)
		return
	}
	t.Logf("categoryId is %d\n", id)
}
