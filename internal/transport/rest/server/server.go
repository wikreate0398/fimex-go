package server

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"wikreate/fimex/internal/config"
	"wikreate/fimex/internal/domain/interfaces"
)

type Server struct {
	http   *http.Server
	logger interfaces.Logger
}

func NewServer(router *gin.Engine, cfg *config.Config, logger interfaces.Logger) *Server {
	return &Server{
		logger: logger,
		http: &http.Server{
			Addr:           fmt.Sprintf(":%d", cfg.Server.Port),
			Handler:        router,
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
			MaxHeaderBytes: 1 << 20,
		},
	}
}

func (s *Server) Start() {
	go func() {
		if err := s.http.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			s.logger.OnError(err, "Failed to start server")
		}
	}()
}

func (s *Server) Stop(ctx context.Context) {
	err := s.http.Shutdown(ctx)
	s.logger.OnError(err, "Failed to stop server")
}
