package user

import (
	"net/http"
	"task3/api"
	"task3/jwt"
	"task3/model"
	"task3/service"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

var userService service.IUserRepository = service.UserRepository{}

type RegisterRequest struct {
	UserName    string `form:"userName" json:"user_name"`
	Pass        string `form:"pass" json:"pass"`
	ConfirmPass string `form:"confirmPass" json:"confirm_pass"`
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func Register(ctx *gin.Context) {
	var param RegisterRequest
	if error := ctx.ShouldBind(&param); error != nil {
		ctx.JSON(400, gin.H{
			"message": "参数错误",
		})
		return
	}

	if userService.GetByName(param.UserName).ID != 0 {
		ctx.JSON(400, gin.H{
			"message": "用户已存在",
		})
		return
	}

	hashedPassword, err := HashPassword(param.Pass)
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": "密码加密失败",
		})
		return
	}

	id := userService.AddOne(&model.User{UserName: param.UserName, Pass: hashedPassword})
	if id > 0 {
		ctx.JSON(http.StatusOK, api.Result{
			Success: true,
			Message: "注册成功",
			Data:    id,
		})
	} else {
		ctx.JSON(http.StatusOK, api.Result{
			Success: false,
			Message: "注册失败",
			Data:    nil,
		})
	}

}

func Login(ctx *gin.Context) {
	userName := ctx.Query("userName")
	pass := ctx.Query("pass")
	if userName == "" || pass == "" {
		ctx.JSON(400, gin.H{
			"message": "用户不存在或密码错误",
		})
	}

	user := userService.GetByName(userName)
	if user.ID == 0 || !CheckPasswordHash(pass, user.Pass) {
		ctx.JSON(400, gin.H{
			"message": "用户不存在或密码错误",
		})
		return
	}
	token, _ := jwt.CreateToken(user.UserName)

	ctx.JSON(http.StatusOK, api.Result{
		Success: true,
		Message: "登录成功",
		Data:    token,
	})

}
func Profile(ctx *gin.Context) {
	userName := api.GetUserName(ctx)
	user := userService.GetByName(userName)
	user.Tags = userService.GetTags(user.ID)
	user.Posts = userService.GetPosts(user.ID)
	user.Pass = ""
	ctx.JSON(http.StatusOK, api.Result{
		Success: true,
		Message: "查询成功",
		Data:    user,
	})
}
