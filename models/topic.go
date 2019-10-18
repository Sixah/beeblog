package models

import (
	"strconv"
	"time"
)

func AddTopic(title,content string) error {
	topic := &Topic{
		Title:title,
		Content:content,
		Created:time.Now(),
		Updated:time.Now(),
		ReplyTime:time.Now(),
	}

	_,err := DB.Insert(topic)
	return err
}

func GetAllTopic(isDesc bool) ([]*Topic,error) {
	topics := []*Topic{}
	var err error
	if isDesc {
		_,err = DB.QueryTable("topic").OrderBy("-created").All(&topics)
	} else {
		_,err = DB.QueryTable("topic").All(&topics)
	}
	return topics,err
}

func GetTopic(tid string) (*Topic,error) {
	tidNum,err := strconv.ParseInt(tid,10,64)
	if err != nil {
		return nil,err
	}

	topic := new(Topic)
	err = DB.QueryTable("topic").Filter("id",tidNum).One(topic)
	if err != nil {
		return nil,err
	}
	topic.Views++
	_,err = DB.Update(topic)
	return topic,err
}

func ModifyTopic(tid,title,content string) error {
	tidNum,err := strconv.ParseInt(tid,10,64)
	if err != nil {
		return err
	}
	topic := &Topic{Id:tidNum}
	if DB.Read(topic) == nil {
		topic.Title = title
		topic.Content = content
		topic.Updated = time.Now()
		_,err := DB.Update(topic)
		return err
	}
	return nil
}

func DeleteTopic(tid string) error {
	tidNum,err := strconv.ParseInt(tid,10,64)
	if err != nil {
		return err
	}
	topic := &Topic{Id:tidNum}
	_,err = DB.Delete(topic)
	return err
}
