package models

import (
	"fmt"
	"strconv"
	"time"
)

func AddCategory(name string) error {
	fmt.Println("2222222222222",name)
	cate := &Category{Title:name,Created:time.Now(),TopicTime:time.Now()}
	qs := DB.QueryTable("category")
	err := qs.Filter("title",name).One(cate)
	fmt.Println("1111111111111111",cate)
	if err == nil {
		return err
	}

	_,err = DB.Insert(cate)
	if err != nil {
		return err
	}
	return nil
}

func GetAllCategories() ([]*Category,error) {
	//o := orm.NewOrm()
	cates := []*Category{}
	//qs := o.QueryTable("category")
	//_,err := qs.All(&cates)
	sql := "select * from category"
	fmt.Println(sql)
	_, err := DB.Raw(sql).QueryRows(&cates)
	return cates,err
}

func DelCategory(id string) error {
	cid,err := strconv.ParseInt(id,10,64)
	if err != nil {
		return err
	}
	cate := &Category{Id:cid}
	_,err = DB.Delete(cate)
	return err
}
