package server

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/sum-project/ublog/cmd/blog-api/src/config/router"
	"github.com/sum-project/ublog/cmd/blog-api/src/controller"
	"github.com/sum-project/ublog/cmd/blog-api/src/repository/post_repository"
	"github.com/sum-project/ublog/cmd/blog-api/src/service/post_service"
	"time"
)

var (
	defaultStopTimeout = time.Second * 30
)

type ApiServer struct {
	addr string
}

func NewApiServer(addr string) (*ApiServer, error) {
	if addr == "" {
		return nil, errors.New("addr cannot be blank")
	}

	return &ApiServer{
		addr: addr,
	}, nil
}

func (s *ApiServer) Start() error {
	engine := gin.Default()
	repository := post_repository.NewPostRepository()
	service := post_service.NewPostService(repository)
	userController := controller.NewPostEndpoint(service)
	router.MapRouter(engine, userController)
	return engine.Run(s.addr)
}
