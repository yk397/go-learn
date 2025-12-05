package service

import (
	"fmt"
	"log"
	"task3/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var dst = "root:123456@tcp(127.0.0.1:3306)/blog?charset=utf8mb4&parseTime=True&loc=Local"
var db *gorm.DB

func init() {
	db1, err := gorm.Open(mysql.Open(dst), &gorm.Config{
		Logger: logger.New(
			log.Default(),
			logger.Config{
				LogLevel: logger.Info,
				Colorful: true,
			},
		),
	})
	if err != nil {
		panic(err.Error())
	}
	db = db1
	dbConnections, err := db1.DB()
	if err != nil {
		panic(err.Error())
	}
	dbConnections.SetMaxOpenConns(100)
	dbConnections.SetMaxIdleConns(10)
}

func Migration() {
	db.AutoMigrate(&model.User{}, &model.Post{}, &model.Comment{}, &model.Tag{})
}

type IUserRepository interface {
	GetById(uint) model.User
	GetByName(string) model.User
	GetTags(uint) []model.Tag
	GetPosts(uint) []model.Post
	AddOne(*model.User) uint
}

type IPostRepository interface {
	AddOne(*model.Post) uint
	UpdateOne(*model.Post) uint
	SearchByTags([]model.Tag) []model.Post
	UpdateStatus(uint8, uint) bool
	AddComment(comment model.Comment, postid uint) bool
	GetById(uint) model.Post
}

type UserRepository struct{}

func (UserRepository) GetById(id uint) model.User {
	var user model.User
	db.First(&user, id)
	return user
}
func (UserRepository) GetByName(name string) model.User {
	var user model.User
	db.Where("user_name = ?", name).First(&user)
	return user
}
func (UserRepository) GetTags(id uint) []model.Tag {
	var tags []model.Tag
	db.Model(&model.Tag{}).Where("user_id = ?", id).Find(&tags)
	return tags
}
func (UserRepository) GetPosts(id uint) []model.Post {
	var posts []model.Post
	db.Model(&model.Post{}).Where("user_id = ?", id).Find(&posts)
	return posts
}
func (UserRepository) AddOne(user *model.User) uint {
	result := db.Create(user)
	if result.Error != nil {
		fmt.Printf("insert user eroor :%s", result.Error.Error())
	}
	return user.ID
}

type PostRepository struct{}

func (PostRepository) AddOne(post *model.Post) uint {
	db.Create(post)
	return post.ID
}
func (PostRepository) UpdateOne(post *model.Post) uint {
	db.Save(post)
	return post.ID
}

func (PostRepository) SearchByTags(tags []model.Tag) []model.Post {
	var posts []model.Post
	result := db.Model(&model.Post{}).Joins("inner join post_tags as pt on post.id = pt.post_id").Where("pt.tag_id in (?)", tags).Select("post.*").Find(&posts)
	if result.Error != nil {
		return nil
	}
	return posts
}

func (PostRepository) UpdateStatus(status uint8, id uint) bool {
	var post model.Post
	db.First(&post, id)
	post.Status = status
	result := db.Save(&post)
	return result.Error != nil
}

func (PostRepository) AddComment(comment model.Comment, postid uint) bool {
	var post model.Post
	check := db.First(&post, postid)
	if check.Error == nil || post.ID == 0 {
		return false
	}
	comment.PostId = postid
	result := db.Create(&comment)
	return result.Error != nil
}

func (PostRepository) GetById(postId uint) model.Post {
	var post model.Post
	result := db.Preload("Comments").Preload("Tags").Preload("User").First(&post, postId)
	if result.Error != nil {
		log.Fatal("查询Post失败,参数=", postId)
	}
	return post
}
