package admin

import (
	"book-list/internal/data"
	"book-list/internal/domain/entities"
	"book-list/internal/middleware"
	"log"

	"github.com/gin-gonic/gin"
)

type AdminService struct {
	logger *log.Logger
	db     *data.Queries
}

type adminEndpoints struct {
	as *AdminService
}

func AddAdmin(router *gin.RouterGroup, logger *log.Logger, db *data.Queries) {
	r := router.Group("/admin", middleware.Auth(logger, middleware.WithRolePolicy(entities.RoleAdmin)))

	as := &AdminService{logger, db}
	ae := &adminEndpoints{as}

	r.GET("/test", func(c *gin.Context) {
		r := c.Value("role").(entities.Role)
		logger.Println(r)
	})

	r.POST("/books", ae.AddBookHandler)

	r.PUT("/users/promote", ae.PromoteUserHandler)
}
