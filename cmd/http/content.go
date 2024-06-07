package http

import (
	"awesome-content/internal/adapters/driven/mysql"
	"awesome-content/internal/adapters/driver/http/handler"
	"awesome-content/internal/core/service"

	"github.com/gin-gonic/gin"
)

func (s *Server) DefineContentRoutes(router *gin.RouterGroup) {
	repo := mysql.NewContentRepository(s.DB)
	srv := service.NewContentService(repo)
	hdl := handler.NewContentHandler(&srv)

	contentRouter := router.Group("content")
	{
		contentRouter.GET("feed", hdl.Feed)
		contentRouter.GET("top", hdl.Top)
		contentRouter.GET("catalog", hdl.Catalog)
		contentRouter.GET("comments", hdl.GetComments)
		contentRouter.GET("preview", hdl.Preview)
	}
}
