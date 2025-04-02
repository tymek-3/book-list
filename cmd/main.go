package main

import (
	"book-list/config"
	"book-list/internal/data"
	"book-list/internal/features"
	"book-list/internal/features/auth"
	"book-list/internal/middleware"
	"book-list/internal/utils"
	"book-list/pkg/gintemplrenderer"

	"github.com/gin-gonic/gin"
	_ "modernc.org/sqlite"
)

func main() {
	err := config.Init()
	if err != nil {
		panic(err)
	}

	logger := utils.NewLogger()

	db, err := data.NewDB(logger)
	if err != nil {
		panic(err)
	}

	engine := gin.Default()
	engine.HTMLRender = &gintemplrenderer.HTMLTemplRenderer{}

	engine.Use(middleware.Errors(logger))

	apiV1 := engine.Group("/api/v1")

	features.AddHealthCheck(apiV1, logger)
	auth.AddAuth(apiV1, logger, db)

	engine.Run(":3000")
}
