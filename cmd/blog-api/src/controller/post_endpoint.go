package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sum-project/ublog/cmd/blog-api/src/domain/post"
	"github.com/sum-project/ublog/cmd/blog-api/src/service/post_service"
	"github.com/sum-project/ublog/pkg/uerror"
	"net/http"
)

type PostEndpoint interface {
	Get(*gin.Context)
	GetAll(*gin.Context)
	Add(*gin.Context)
	Update(*gin.Context)
	Delete(*gin.Context)
}

func NewPostEndpoint(service post_service.PostService) PostEndpoint {
	return &postEndpoint{
		postService: service,
	}
}

type postEndpoint struct {
	postService post_service.PostService
}

func (endpoint *postEndpoint) Get(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, "Implement me!")
}
func (endpoint *postEndpoint) GetAll(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, "Implement me!")
}
func (endpoint *postEndpoint) Add(c *gin.Context) {
	var p post.Post
	if err := c.ShouldBindJSON(&p); err != nil {
		apiErr := uerror.NewBadRequestError("invalid json body")
		c.JSON(apiErr.Status(), apiErr)
	}

	newPost, err := endpoint.postService.AddPost(p)
	if err != nil {
		c.JSON(err.Status(), err)
	}

	c.JSON(http.StatusCreated, newPost)
}
func (endpoint *postEndpoint) Update(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, "Implement me!")
}
func (endpoint *postEndpoint) Delete(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, "Implement me!")
}
