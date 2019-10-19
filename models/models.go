package models

import (
	"github.com/Unknwon/com"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"path"
	"time"
)

const (
	_DB_NAME        = "data/beeblog.db"
	_SQLITE3_DRIVER = "sqlite3"
)

type Category struct {
	Id              int64
	Title           string    // 标题
	Created         time.Time `orm:"index"` // 分类创建时间
	Updated         time.Time // 分类更新时间
	Views           int64     `orm:"index"` // 浏览的次数
	TopicTime       time.Time `orm:"index"` // 文章发表时间
	TopicCount      int64     // 文章数量
	TopicLastUserId int64     // 最后一次更新分类的用户id
}

/*
CREATE TABLE `category` (
     `id` integer NOT NULL PRIMARY KEY AUTOINCREMENT,
     `title` varchar(255) NOT NULL DEFAULT '' ,
     `created` datetime NOT NULL DEFAULT current_timestamp,
     `views` integer NOT NULL DEFAULT 0 ,
     `topic_time` datetime NOT NULL DEFAULT current_timestamp,
     `topic_count` integer NOT NULL DEFAULT 0 ,
     `topic_last_user_id` integer NOT NULL DEFAULT 0
     );
CREATE INDEX `category_created` ON `category` (`created`);
CREATE INDEX `category_views` ON `category` (`views`);
CREATE INDEX `category_topic_time` ON `category` (`topic_time`);
*/
type Topic struct {
	Id              int64
	Uid             int64     // 用户id
	Type            string    // 文章分类
	Labels string // 标签
	Title           string    // 文章标题
	Content         string    `orm:"size(5000)"` // 文章内容
	Attachment      string    // 附件
	Created         time.Time `orm:"index"` // 文章创建时间
	Updated         time.Time `orm:"index"` // 文章更新时间
	Views           int64     `orm:"index"` // 浏览次数
	Author          string    // 作者姓名
	ReplyTime       time.Time `orm:"index"` // 评论回复时间
	ReplyCount      int64     // 评论数量
	ReplyLastUserId int64     // 最后评论的用户id
}

/*
 CREATE TABLE IF NOT EXISTS `topic` (
        `id` integer NOT NULL PRIMARY KEY AUTOINCREMENT,
        `uid` integer NOT NULL DEFAULT 0 ,
        `title` varchar(255) NOT NULL DEFAULT '' ,
        `content` varchar(5000) NOT NULL DEFAULT '' ,
        `attachment` varchar(255) NOT NULL DEFAULT '' ,
        `created` datetime NOT NULL DEFAULT current_timestamp,
        `updated` datetime NOT NULL DEFAULT current_timestampL,
        `views` integer NOT NULL DEFAULT 0 ,
        `author` varchar(255) NOT NULL DEFAULT '' ,
        `reply_time` datetime NOT NULL DEFAULT current_timestamp,
        `reply_count` integer NOT NULL DEFAULT 0 ,
        `reply_last_user_id` integer NOT NULL DEFAULT 0
    );
CREATE INDEX `topic_created` ON `topic` (`created`);
CREATE INDEX `topic_updated` ON `topic` (`updated`);
CREATE INDEX `topic_views` ON `topic` (`views`);
*/

type Comment struct {
	Id int64
	Tid int64
	Name string
	Content string `orm:"size(1000)"`
	Created time.Time
}

var DB orm.Ormer

func RegisterDB() {
	if !com.IsExist(_DB_NAME) {
		os.MkdirAll(path.Dir(_DB_NAME), os.ModePerm)
		os.Create(_DB_NAME)
	}

	orm.RegisterModel(new(Category), new(Topic),new(Comment))
	orm.RegisterDriver(_SQLITE3_DRIVER, orm.DRSqlite)
	orm.RegisterDataBase("default", _SQLITE3_DRIVER, _DB_NAME, 10)
	DB = orm.NewOrm()
}
