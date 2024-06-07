package handler

import (
	"awesome-content/internal/ports/driver/service"

	"github.com/gin-gonic/gin"
)

type ContentHandler struct {
	service service.ContentServicePort
}

func NewContentHandler(srv *service.ContentServicePort) *ContentHandler {
	return &ContentHandler{
		service: *srv,
	}
}

func (hdl *ContentHandler) Feed(ctx *gin.Context) {

	hdl.service.Feed(ctx)

	ctx.JSON(200, map[string]any{
		"message": "feed content",
	})
}

func (hdl *ContentHandler) Top(ctx *gin.Context) {

	hdl.service.Top(ctx)

	ctx.JSON(200, map[string]any{
		"message": "list of top content",
	})
}

func (hdl *ContentHandler) Catalog(ctx *gin.Context) {

	hdl.service.Catalog(ctx)

	ctx.JSON(200, map[string]any{
		"message": "show catalog",
	})
}

func (hdl *ContentHandler) GetComments(ctx *gin.Context) {

	hdl.service.GetComments(ctx)

	ctx.JSON(200, map[string]any{
		"message": "comments",
	})
}

func (hdl *ContentHandler) Preview(ctx *gin.Context) {

	hdl.service.Preview(ctx)

	ctx.JSON(200, map[string]any{
		"message": "previewing",
	})
}
