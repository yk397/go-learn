package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	UserName string `gorm:"column:user_name;type:varchar(30)"`
	Pass     string `gorm:"column:pass;type:varchar(100)"`
	Tags     []Tag  `gorm:"foreignKey:UserId"`
}

func (User) TableName() string {
	return "User"
}

type Comment struct {
	gorm.Model
	PostId  uint    `gorm:"column:post_id;type:bigint"`
	Content *string `gorm:"column:content:content;type:longtext"`
	UserId  uint    `gorm:"column:user_id;type:bigint"`
	User    User    `gorm:"references:UserId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (Comment) TableName() string {
	return "Comment"
}

type Post struct {
	gorm.Model
	UserId   uint      `gorm:"column:user_id;type:bigint"`
	Title    string    `gorm:"column:title;type:varchar(100)"`
	Content  *string   `gorm:"column:content;type:longtext"`
	Status   uint8     `gorm:"column:status:type:tinyint"`
	Comments []Comment `gorm:"foreignKey:PostId"`
	User     User      `gorm:"references:UserId"`
	Tags     []Tag     `gorm:"many2many:post_tags;references:ID"`
}

func (Post) TableName() string {
	return "posts"
}

type Tag struct {
	gorm.Model
	Name   string `gorm:"column:name;type:varchar(20)"`
	UserId uint   `gorm:"column:user_id;type:bigint"`
	Posts  []Post `gorm:"many2many:post_tags;references:ID"`
}

func (Tag) TableName() string {
	return "Tag"
}
