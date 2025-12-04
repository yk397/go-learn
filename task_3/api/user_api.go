package api

import (
	"crypto"
	"net/http"
	"task3/model"
	"task3/service"

	"github.com/gin-gonic/gin"
)

type RegisterRequest struct {
	UserName    string `form:"userName" json:"user_name"`
	Pass        string `form:"pass" json:"pass"`
	ConfirmPass string `form:"confirmPass" json:"confirm_pass"`
}

type Result struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Register(ctx *gin.Context) {
	var userService service.IUserRepository = service.UserRepository{}

	var param RegisterRequest
	if error := ctx.ShouldBind(&param); error != nil {
		ctx.JSON(400, gin.H{
			"message": "参数错误",
		})
	}
	if userService.GetByName(param.UserName).ID != 0 {
		ctx.JSON(400, gin.H{
			"message": "用户已存在",
		})
	}
	id := userService.AddOne(&model.User{UserName: param.UserName, Pass: string(crypto.MD5.New().Sum([]byte(param.Pass)))})
	ctx.JSON(http.StatusOK, Result{
		Success: true,
		Message: "注册成功",
		Data:    id,
	})
}

func Login(ctx *gin.Context) {
	var userService service.IUserRepository = service.UserRepository{}
	var param RegisterRequest
	if error := ctx.ShouldBind(&param); error != nil {
		ctx.JSON(400, gin.H{
			"message": "参数错误",
		})
	}
	user := userService.GetByName(param.UserName)
	if user.ID == 0 || user.Pass != string(crypto.MD5.New().Sum([]byte(param.Pass))) {
		ctx.JSON(400, gin.H{
			"message": "用户不存在或密码错误",
		})
	}
	ctx.JSON(http.StatusOK, Result{
		Success: true,
		Message: "登录成功",
		Data:    user.ID,
	})

}
