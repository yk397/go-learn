package api

import (
	"net/http"
	"strings"
	"task3/jwt"

	"github.com/gin-gonic/gin"
)

type Result struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func GetUserName(ctx *gin.Context) string {
	tokenStr := ctx.GetHeader("token")
	if tokenStr == "" {
		ctx.JSON(http.StatusForbidden, Result{
			Success: false,
			Message: "用户未登录",
		})
	}
	claim := jwt.MyCustomClaims{}
	token := strings.Replace(tokenStr, "Barer ", "", 1)
	if t := jwt.ParseCustomClaims(token, &claim); t == nil || claim.Foo == "" {
		ctx.JSON(http.StatusForbidden, Result{
			Success: false,
			Message: "用户未登录",
		})
	}
	return claim.Foo
}
