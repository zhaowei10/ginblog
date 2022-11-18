package v1

import (
	"ginblog/model"
	"ginblog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//添加分类
func AddCategory(c *gin.Context) {
	var category model.Category
	c.ShouldBind(&category)
	code = model.CheckUser(category.Name)
	if code == errmsg.SUCCESS {
		model.CreateCategory(&category)
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"data":    category,
			"message": errmsg.GetErrMsg(code),
		})
	} else {
		c.JSON(500, gin.H{
			"status":  code,
			"data":    category,
			"message": errmsg.GetErrMsg(code),
		})
	}

}

//todo 查询分类下的所有文章
//查询分类列表
func GetCategoryList(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	if pageNum == 0 {
		pageNum = -1
	}
	list, total := model.GetCategoryList(pageSize, pageNum)
	code = errmsg.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    list,
		"total":   total,
		"message": errmsg.GetErrMsg(code),
	})
}

// 编辑分类名

func EditCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var data model.Category
	c.ShouldBindJSON(&data)
	code = model.CheckCategory(data.Name)
	if code == errmsg.SUCCESS {
		model.EditCategory(id, &data)
	}
	if code == errmsg.ERROR_CATEGORYNAME_USED {
		c.Abort()
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

//删除分类
func DeleteCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code = model.DeleteCategory(id)
	c.JSON(code, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
