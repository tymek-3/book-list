package books

import (
	"book-list/internal/data"
	"log"

	"github.com/gin-gonic/gin"
)

type BooksService struct {
	logger *log.Logger
	db     *data.Queries
}

type booksEndpoints struct {
	bs *BooksService
	r  *gin.RouterGroup
}

func AddBooksEndpoints(router *gin.RouterGroup, logger *log.Logger, db *data.Queries) {
	r := router.Group("/books")

	bs := &BooksService{logger, db}
	be := &booksEndpoints{bs, r}

	_ = be
}
