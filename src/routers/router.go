package routers

import (
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := r.Group("/api")
	{
		//获取标签列表
		api.GET("/tags", v1.GetTags)
		//新建标签
		api.POST("/tags", v1.AddTag)
		//更新指定标签
		api.PUT("/tags/:id", v1.EditTag)
		//删除指定标签
		api.DELETE("/tags/:id", v1.DeleteTag)
	}

	return r
}
