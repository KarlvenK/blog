package routes

import (
	"blog/api/v1"
	"blog/utils"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()

	router := r.Group("api/v1")
	{
		//User model router api
		router.POST("user/add", v1.AddUser)
		router.GET("users", v1.GetUsers)
		router.PUT("user/:id", v1.EditUser)
		router.DELETE("user/:id", v1.DeleteUser)
		//category model router api
		router.POST("category/add", v1.AddCategory)
		router.GET("category", v1.GetCate)
		router.PUT("category/:id", v1.EditCate)
		router.DELETE("category/:id", v1.DeleteCate)
		//article model router api
		router.POST("article/add", v1.AddArticle)
		router.GET("article", v1.GetArt)
		router.PUT("article/:id", v1.EditArt)
		router.DELETE("article/:id", v1.DeleteArt)
	}

	r.Run(utils.HttpPort)
}
