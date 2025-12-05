package main

import (
	"task3/api/user"

	"task3/api/post"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	r.POST("/register", user.Register)

	r.POST("/login", user.Login)

	r.GET("/profile", user.Profile)

	postRoute := r.Group("/post")
	postRoute.POST("save", post.AddPost)

	r.Run()
}
