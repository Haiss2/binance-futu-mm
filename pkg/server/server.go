package server

import (
	"fmt"
	"net/http"

	"github.com/Haiss2/binance-futu-mm/pkg/hunter"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

type Server struct {
	s        *gin.Engine
	bindAddr string

	hunter *hunter.Hunter
}

func New(bindAddr string, hunter *hunter.Hunter) *Server {
	engine := gin.New()
	engine.Use(gin.Recovery())

	s := &Server{
		s:        engine,
		bindAddr: bindAddr,
		hunter:   hunter,
	}

	gin.SetMode(gin.ReleaseMode)

	api := engine.Group("/api")

	api.GET("/", s.helloWorld)
	s.register()

	return s
}

func (s *Server) Run() error {
	if err := s.s.Run(s.bindAddr); err != nil {
		return fmt.Errorf("run server: %w", err)
	}

	return nil
}

func (s *Server) register() {
	pprof.Register(s.s, "/debug")
}

func (s *Server) helloWorld(c *gin.Context) {
	c.JSON(http.StatusOK, "hello world")
}
