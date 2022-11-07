package server

import (
	"context"
	"errors"
	"log"
	"net/http"
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

func (s *ApiServer) Start(stop <-chan struct{}) error {
	srv := &http.Server{
		Addr: s.addr,
	}

	go func() {
		log.Printf("starting server on addr: %s", s.addr)
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			log.Printf("listen: %s\n", err)
		}
	}()

	<-stop
	ctx, cancel := context.WithTimeout(context.Background(), defaultStopTimeout)
	defer cancel()

	log.Printf("stopping server timeout: %d", defaultStopTimeout)
	return srv.Shutdown(ctx)
}
