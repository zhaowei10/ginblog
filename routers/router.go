package routers

import (
	v1 "ginblog/api/v1"
	"ginblog/middleware"
	"ginblog/utils"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()
	authRouter := r.Group("api/v1")
	authRouter.Use(middleware.JwtToken())
	{
		//用户模块路由接口
		authRouter.PUT("/user/:id", v1.EditeUser)
		authRouter.DELETE("user/:id", v1.DeleteUser)

		//文章模块路由接口
		authRouter.POST("/article/add", v1.AddArticle)
		authRouter.PUT("/article/:id", v1.EditArticle)
		authRouter.DELETE("article/:id", v1.DeleteArticle)

		//分类模块路由接口
		authRouter.POST("/category/add", v1.AddCategory)
		authRouter.PUT("/category/:id", v1.EditCategory)
		authRouter.DELETE("category/:id", v1.DeleteCategory)
		//上传文件
		authRouter.POST("/upload", v1.Upload)
	}
	router := r.Group("api/v1")
	{
		router.GET("/user/list", v1.GetUserList)
		router.GET("/article/:id", v1.GetArticleInfo)
		router.GET("/article/list", v1.GetArticleList) //查询所有文章
		router.GET("/article/categorylist/:id", v1.GetCateArticleList)
		router.GET("/category/list", v1.GetCategoryList)
		router.POST("/login", v1.Login)
		router.POST("/user/add", v1.AddUser)
	}

	r.Run(utils.HttpPort)

}
