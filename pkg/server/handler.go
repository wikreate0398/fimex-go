package server

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"wikreate/fimex/internal/config"
)

type Server struct {
	http *http.Server
}

func NewServer(router *gin.Engine, cfg *config.Config) *Server {
	return &Server{
		http: &http.Server{
			Addr:           fmt.Sprintf(":%d", cfg.Server.Port),
			Handler:        router,
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
			MaxHeaderBytes: 1 << 20,
		},
	}
}

func (s *Server) Start() error {
	return s.http.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	return s.http.Shutdown(ctx)
}
