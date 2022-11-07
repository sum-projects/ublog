package router

import (
	"github.com/gin-gonic/gin"
	"github.com/sum-project/ublog/cmd/blog-api/src/controller"
)

func MapRouter(router *gin.Engine, endpoint controller.PostEndpoint) {
	v1 := router.Group("/api/v1")
	v1.GET("/user", endpoint.Get)
	v1.GET("/users", endpoint.GetAll)
	v1.POST("/user", endpoint.Add)
	v1.PUT("/user", endpoint.Update)
	v1.DELETE("/user", endpoint.Delete)
}
