package models

import (
	"github.com/astaxie/beego/orm"
	"strconv"
	"strings"
	"time"
)

// 添加文章
func AddTopic(tType,labels, title, content string) error {
	labels = "$"+strings.Join(strings.Split(labels," "),"#$") + "#"
	db := orm.NewOrm()
	db.Begin() // 开启事务(不能使用全局的数据库连接，会报错)
	topic := &Topic{
		Type:      tType,
		Labels:labels,
		Title:     title,
		Content:   content,
		Created:   time.Now(),
		Updated:   time.Now(),
		ReplyTime: time.Now(),
	}
	cate := new(Category)
	err := db.QueryTable("category").Filter("title", tType).One(cate)
	if err == nil { // 分类名称存在
		cate.TopicCount = cate.TopicCount + 1
		cate.Updated = time.Now()
		cate.TopicTime = time.Now()
		_, err = db.Update(cate)
		if err != nil {
			db.Rollback()
			return err
		}
		_, err = db.Insert(topic)
		if err != nil {
			db.Rollback()
			return err
		}
		db.Commit()
		return nil
	} else { // 查询出错
		if err == orm.ErrNoRows { // 没有查询到分类名称
			cate = &Category{
				Title:      tType,
				Created:    time.Now(),
				Updated:    time.Now(),
				TopicTime:  time.Now(),
				TopicCount: 1,
			}
			_, err = db.Insert(cate) // 新建分类
			if err != nil {
				db.Rollback()
				return err
			}
		} else { // 其他错误 返回
			db.Rollback()
			return err
		}
	}

	_, err = db.Insert(topic) // 新建文章
	if err != nil {
		db.Rollback()
		return err
	}
	db.Commit()
	return nil
}

// 查询文章列表
func GetAllTopic(cate,label string,isDesc bool) ([]Topic, error) {
	topics := []Topic{}
	var err error
	qs := DB.QueryTable("topic")
	if isDesc { // 按时间倒序
		if len(cate) > 0 {
			qs = qs.Filter("type",cate)
		}
		if len(label) > 0 {
			qs = qs.Filter("labels__contains","$"+label+"#")
		}
		_, err = qs.OrderBy("-created").All(&topics)

	} else { // 按id主键顺序
		_, err = DB.QueryTable("topic").All(&topics)
	}
	return topics, err
}

// 查询文章
func GetTopic(tid string) (*Topic, error) {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return nil, err
	}

	topic := new(Topic)
	err = DB.QueryTable("topic").Filter("id", tidNum).One(topic)
	if err != nil {
		return nil, err
	}
	topic.Views++ // 浏览次数+1
	_, err = DB.Update(topic)
	topic.Labels = strings.Replace(strings.Replace(topic.Labels,"#","",-1),"$"," ",-1)
	return topic, err
}

// 修改文章
func ModifyTopic(tType, labels,tid, title, content string) error {
	labels = "$"+strings.Join(strings.Split(labels," "),"#$") + "#"
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return err
	}
	topic := &Topic{}
	err = DB.QueryTable("topic").Filter("id", tidNum).One(topic)
	if err == nil { // 修改的文章存在
		db := orm.NewOrm() // 新建数据库连接实例(开启事务不能使用全局变量，当多个请求进来时，一个请求的事务的操作会被锁住，其他事务无法进来)
		db.Begin()
		cate := new(Category)
		if topic.Type != "" { // 分类名称不为空
			err = db.QueryTable("category").Filter("title", topic.Type).One(cate) // 文章所属分类信息
			if err != nil {
				db.Rollback()
				return err
			}

			if cate.Title != tType { // 文章分类需要修改
				cate.TopicCount = cate.TopicCount - 1 // 原分类文章数量-1
				cate.Updated = time.Now()
				_, err := db.Update(cate)
				if err != nil {
					db.Rollback()
					return err
				}

				cate1 := new(Category)
				err = db.QueryTable("category").Filter("title", tType).One(cate1) // 新修改分类信息
				if err != nil {
					if err == orm.ErrNoRows { // 新分类不存在
						cateNew := &Category{
							Title:      tType,
							Created:    time.Now(),
							Updated:    time.Now(),
							TopicTime:  time.Now(),
							TopicCount: 0,
						}
						_, err = db.Insert(cateNew) // 新建分类信息
						if err != nil {
							db.Rollback()
							return err
						}
						err = db.QueryTable("category").Filter("title", tType).One(cate1) // 重新获取新修改分类信息
						if err != nil {
							db.Rollback()
							return err
						}
					} else { // 其他错误 返回
						db.Rollback()
						return err
					}
				}
				cate1.TopicCount = cate1.TopicCount + 1 // 新修改分类文章数量+1
				cate1.Updated = time.Now()
				_, err = db.Update(cate1)
				if err != nil {
					db.Rollback()
					return err
				}
			}
		} else { // 文章分类为空
			err = db.QueryTable("category").Filter("title", tType).One(cate) // 新修改分类信息
			if err != nil {
				if err == orm.ErrNoRows { // 新修改分类不存在
					cateNew := &Category{
						Title:      tType,
						Created:    time.Now(),
						Updated:    time.Now(),
						TopicTime:  time.Now(),
						TopicCount: 0,
					}
					_, err = db.Insert(cateNew) // 新建分类信息
					if err != nil {
						db.Rollback()
						return err
					}
					err = db.QueryTable("category").Filter("title", tType).One(cate) // 重新获取新修改分类信息
					if err != nil {
						db.Rollback()
						return err
					}
				} else { // 其他错误 返回
					db.Rollback()
					return err
				}
			}
			cate.TopicCount = cate.TopicCount + 1 // 新修改分类文章数量+1
			cate.Updated = time.Now()
			_, err = db.Update(cate)
			if err != nil {
				db.Rollback()
				return err
			}
		}

		topic.Type = tType
		topic.Labels = labels
		topic.Title = title
		topic.Content = content
		topic.Updated = time.Now()
		_, err := db.Update(topic) // 修改文章信息
		if err != nil {
			db.Rollback()
			return err
		}
		db.Commit()
		return nil
	}
	return nil
}

// 删除文章
func DeleteTopic(tid string) error {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return err
	}

	db := orm.NewOrm()
	db.Begin()

	topic := new(Topic)
	err = db.QueryTable("topic").Filter("id", tidNum).One(topic) // 查询文章信息
	if err != nil {
		db.Rollback()
		return err
	}

	cate := new(Category)
	err = db.QueryTable("category").Filter("title", topic.Type).One(cate) // 查询文章所属分类信息
	if err != nil {
		db.Rollback()
		return err
	}
	cate.Updated = time.Now()
	cate.TopicCount = cate.TopicCount - 1 // 文章所属分类文章数量-1
	_, err = db.Update(cate)
	if err != nil {
		db.Rollback()
		return err
	}

	//topic := &Topic{Id:tidNum}
	_, err = db.Delete(topic) // 删除文章
	if err != nil {
		db.Rollback()
		return err
	}
	db.Commit()
	return nil
}
