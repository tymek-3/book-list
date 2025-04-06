package admin

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (ae *adminEndpoints) AddBookHandler(c *gin.Context) {
}

type addBookRequest struct {
	name             string
	publication_date string
	type_id          uuid.UUID
	author_id        uuid.UUID
	publisher_id     uuid.UUID
}

func (as *AdminService) AddBook(ctx context.Context, request addBookRequest) {
}
