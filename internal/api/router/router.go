package router

import (
	"github.com/gin-gonic/gin"

	"github.com/Banh-Canh/salaryman/internal/api/handlers"
	"github.com/Banh-Canh/salaryman/internal/services"
)

func Init(service *services.ResumeService) *gin.Engine {
	router := gin.New()

	router.GET("/status", handlers.Status())
	router.POST("/pdf", handlers.GetPdf(service))

	return router
}
