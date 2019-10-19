package models

import (
	"fmt"
	"strconv"
	"time"
)

// 查询分类信息
func AddCategory(name string) error {
	cate := &Category{Title: name, Created: time.Now(), Updated: time.Now(), TopicTime: time.Now()}
	qs := DB.QueryTable("category")
	err := qs.Filter("title", name).One(cate)
	if err == nil {
		return err
	}

	_, err = DB.Insert(cate)
	if err != nil {
		return err
	}
	return nil
}

// 查询分类列表
func GetAllCategories() ([]Category, error) {
	cates := []Category{}
	sql := "select * from category"
	fmt.Println(sql)
	_, err := DB.Raw(sql).QueryRows(&cates)
	return cates, err
}

// 删除分类
// TODO: 删除分类前，需要检查该分类是否有文章，没有文章才能删除
func DelCategory(id string) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	cate := &Category{Id: cid}
	_, err = DB.Delete(cate)
	return err
}
