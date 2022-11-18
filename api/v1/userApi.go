package v1

import (
	"ginblog/model"
	"ginblog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var code int

//添加用户
func AddUser(c *gin.Context) {
	var user model.User
	c.ShouldBind(&user)
	code = model.CheckUser(user.Username)
	if code == errmsg.SUCCESS {
		model.CreateUser(&user)
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"data":    user,
			"message": errmsg.GetErrMsg(code),
		})
	} else {
		c.JSON(500, gin.H{
			"status":  code,
			"data":    user,
			"message": errmsg.GetErrMsg(code),
		})
	}

}

//查询单个用户
//查询用户列表
func GetUserList(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	/*if pageSize == 0 {
		pageSize = -1
	}*/
	if pageNum == 0 {
		pageNum = -1
	}
	list, total := model.GetUserList(pageSize, pageNum)
	code = errmsg.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    list,
		"total":   total,
		"message": errmsg.GetErrMsg(code),
	})
}

// 编辑用户

func EditeUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var data model.User
	c.ShouldBindJSON(&data)
	code = model.CheckUser(data.Username)
	if code == errmsg.SUCCESS {
		model.EditUser(id, &data)
	}
	if code == errmsg.ERROR_USERNAME_USED {
		c.Abort()
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

//删除用户
func DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code = model.DeleteUser(id)
	c.JSON(code, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
