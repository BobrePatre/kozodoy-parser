package parser

import (
	"github.com/BobrePatre/kozodoy-parser/internal/delivery/http/datatransfers/requests"
	"github.com/gin-gonic/gin"
	"io"
	"log/slog"
	"mime/multipart"
	"net/http"
)

type Service interface {
	Parse(fileReader io.Reader, menuType string) error
}

type Handler struct {
	service Service
}

func NewHandler(s Service) *Handler {
	return &Handler{
		service: s,
	}
}

func (h *Handler) Parse(ctx *gin.Context) {
	var req requests.ParseMenu
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if req.File.Size == 0 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "file size is zero"})
		return
	}

	fileReader, err := req.File.Open()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "cannot open file",
		})
		return
	}
	defer func(file multipart.File) {
		err := file.Close()
		if err != nil {
			slog.Error("failed to close multipart file", slog.String("error", err.Error()))
		}
	}(fileReader)

	slog.Debug("file content type", "type", req.File.Header.Get("content-type"))

	if err := h.service.Parse(fileReader, req.MenuType); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(200)
}
