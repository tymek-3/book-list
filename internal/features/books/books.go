package books

import (
	"book-list/internal/data"
	"book-list/internal/features/shared"
	"database/sql"
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

func AddBooks(router *gin.RouterGroup, logger *log.Logger, db *data.Queries) {
	r := router.Group("/books")

	bs := &BooksService{logger, db}
	be := &booksEndpoints{bs, r}

	r.GET("/", func(c *gin.Context) {
		bookRows, err := bs.db.BookFullAll(c)
		if err != nil {
			panic(err)
		}

		books := make([]bookResponse, 0, len(bookRows))
		for _, b := range bookRows {
			books = append(books, BookResponseFromBook(*b.ToEntity()))
		}

		c.JSON(200, shared.SliceResponse[bookResponse]{
			Data: books,
		})
	})

	r.GET("/search", func(c *gin.Context) {
		q := c.Query("q")
		bookRows, err := bs.db.BookFullSearch(c, sql.NullString{String: q, Valid: true})
		if err != nil {
			panic(err)
		}

		books := make([]bookResponse, 0, len(bookRows))
		for _, b := range bookRows {
			books = append(books, BookResponseFromBook(*b.ToEntity()))
		}

		c.JSON(200, shared.SliceResponse[bookResponse]{
			Data: books,
		})
	})
	_ = be
}
