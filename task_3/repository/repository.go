package repository

import "task3/model"

type UserRepository interface {
	GetById(uint) model.User
	GetByName(string) model.User
	GetTags(uint) []model.Tag
	GetPosts(uint) []model.Post
}

type PostRepository interface {
	AddOne(*model.Post) uint
	UpdateOne(*model.Post) uint
	SearchByTags([]model.Tag) []model.Post
	UpdateStatus(uint8, uint32) bool
}
