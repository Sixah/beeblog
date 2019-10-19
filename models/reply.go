package models

import (
	"github.com/astaxie/beego/orm"
	"strconv"
	"time"
)

func AddReply(tid,nickname,content string) error {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return err
	}
	topic := new(Topic)
	err = DB.QueryTable("topic").Filter("id",tidNum).One(topic)
	if err != nil {
		return err
	}
	db := orm.NewOrm()
	db.Begin()

	reply := &Comment{
		Tid:tidNum,
		Name:nickname,
		Content:content,
		Created:time.Now(),
	}
	_,err = db.Insert(reply)
	if err != nil {
		db.Rollback()
		return err
	}

	topic.ReplyCount = topic.ReplyCount+1
	topic.ReplyTime = time.Now()
	_,err = db.Update(topic)
	if err != nil {
		db.Rollback()
		return err
	}
	db.Commit()
	return nil
}

func GetAllReplies(tid string) ([]Comment,error) {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return nil,err
	}
	replies := []Comment{}
	_,err = DB.QueryTable("comment").Filter("tid",tidNum).All(&replies)
	return replies,err
}

func DeleteReply(rid string) error {
	ridNum, err := strconv.ParseInt(rid, 10, 64)
	if err != nil {
		return err
	}
	reply := &Comment{}
	err = DB.QueryTable("comment").Filter("id",ridNum).One(reply)
	if err != nil {
		return err
	}
	topic := new(Topic)
	err = DB.QueryTable("topic").Filter("id",reply.Tid).One(topic)
	if err != nil {
		return err
	}

	db := orm.NewOrm()
	db.Begin()

	_,err = DB.Delete(reply)
	if err != nil {
		db.Rollback()
		return err
	}
	lastReply := new(Comment)
	err = db.QueryTable("comment").Filter("tid",reply.Tid).OrderBy("-created").One(lastReply)
	if err != nil {
		db.Rollback()
		return err
	}
	topic.ReplyTime = lastReply.Created
	topic.ReplyCount = topic.ReplyCount-1
	_,err = db.Update(topic)
	if err != nil {
		db.Rollback()
		return err
	}
	db.Commit()
	return nil
}
