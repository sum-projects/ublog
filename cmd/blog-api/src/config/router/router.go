package router

import (
	"github.com/gin-gonic/gin"
	"github.com/sum-project/ublog/cmd/blog-api/src/controller"
)

func MapRouter(router *gin.Engine, endpoint controller.PostEndpoint) {
	v1 := router.Group("/api/v1")
	v1.GET("/post", endpoint.Get)
	v1.GET("/posts", endpoint.GetAll)
	v1.POST("/post", endpoint.Add)
	v1.PUT("/post", endpoint.Update)
	v1.DELETE("/post", endpoint.Delete)
}
