package features

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddHealthCheck(router *gin.RouterGroup, logger *log.Logger) {
	router.GET("/health", getHealthCheck(logger))
}

func getHealthCheck(logger *log.Logger) func(c *gin.Context) {
	logger.Println("Healthy")
	return func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	}
}
