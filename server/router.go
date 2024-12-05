package server

import (
	"file-registry/config"
	"file-registry/handlers"
	"file-registry/service"

	"github.com/gin-gonic/gin"
)

func InitRouter(config config.Config) (router *gin.Engine) {
	router = gin.Default()

	serviceInstance := service.NewService(config)
	handler := handlers.NewHandler(serviceInstance)

	router.POST("/v1/files", handler.UploadFileHandler)
	router.GET("/v1/files", handler.GetFileHandler)
	return router
}
