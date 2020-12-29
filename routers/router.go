package routers

import (
	"go-gin-example/middleware/jwt"
	"go-gin-example/routers/api"
	v1 "go-gin-example/routers/api/v1"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	_ "go-gin-example/docs"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	// router.Use(gin.Logger())
	// router.Use(gin.Recovery())
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/auth", api.GetAuth)
	//定义v1的路由
	apiv1 := router.Group("/api/v1")
	apiv1.Use(jwt.JWT())
	{
		apiv1.GET("/tags", v1.GetTags)
		apiv1.POST("/tags", v1.AddTasg)
		apiv1.PUT("/tags/:id", v1.EditTag)
		apiv1.DELETE("/tags/:id", v1.DeleteTag)

		apiv1.GET("/articles", v1.GetArticles)
		apiv1.GET("/articles/:id", v1.GetArticle)
		apiv1.POST("/articles", v1.AddArticle)
		apiv1.PUT("/articles/:id", v1.EditArticle)
		apiv1.DELETE("/articles/:id", v1.DeleteArticle)
	}
	return router
}
