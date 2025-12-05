package post

import (
	"task3/api"
	"task3/model"
	"task3/service"

	"github.com/gin-gonic/gin"
)

var userService service.IUserRepository = service.UserRepository{}
var postService service.IPostRepository = service.PostRepository{}

const (
	DRAFT = iota
	POSTED
)

type PostRequest struct {
	Title   string   `form:"title" json:"title"`
	Content string   `form:"content" json:"content"`
	Tags    []string `form:"tags" json:"tags"`
}

func AddPost(ctx *gin.Context) {
	userName := api.GetUserName(ctx)
	user := userService.GetByName(userName)

	param := PostRequest{}
	if err := ctx.ShouldBind(&param); err != nil {
		ctx.JSON(400, api.Result{
			Success: false,
			Message: "参数错误",
		})
	}

	post := model.Post{
		Title:   param.Title,
		Content: &param.Content,
		Status:  DRAFT,
		UserId:  user.ID,
	}

	tags := []model.Tag{}
	for _, v := range param.Tags {
		tags = append(tags, model.Tag{
			Name:   v,
			UserId: user.ID,
		})
	}
	post.Tags = tags

	postService.AddOne(&post)

	ctx.JSON(200, api.Result{
		Success: true,
		Message: "添加成功",
		Data:    post.ID,
	})

}
