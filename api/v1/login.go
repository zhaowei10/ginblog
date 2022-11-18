package v1

import (
	"ginblog/middleware"
	"ginblog/model"
	"ginblog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {
	var user model.User
	var token string
	var code int
	c.ShouldBindJSON(&user)
	code = model.CheckLogin(user.Username, user.Password)
	if code == errmsg.SUCCESS {
		token, code = middleware.SetToken(user.Username)
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
		"token":   token,
	})
}
