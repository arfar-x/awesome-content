package http

import (
	"fmt"

	"awesome-content/config"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Implements EngineInterface
type Server struct {
	Config *config.AppConfig
	DB     *gorm.DB
	engine *gin.Engine
}

func NewHttpServer(config *config.AppConfig, DB *gorm.DB) *Server {
	return &Server{
		Config: config,
		DB:     DB,
	}
}

func (s *Server) Boot() error {
	return nil
}

func (s *Server) Run() error {
	var mode string
	if s.Config.Mode == "production" || s.Config.Mode == "release" {
		mode = "release"
	} else {
		mode = "debug"
	}

	gin.SetMode(mode)
	s.engine = gin.New()

	s.InitRoutes()

	err := s.engine.Run(fmt.Sprintf("%s:%d", s.Config.Http.Host, s.Config.Http.Port))
	if err != nil {
		return err
	}

	return nil
}

func (s *Server) Shutdown() error {
	return nil
}
