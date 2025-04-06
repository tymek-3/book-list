package books

import (
	"book-list/internal/domain/entities"
	"time"

	"github.com/google/uuid"
)

type idName struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type bookResponse struct {
	ID              uuid.UUID `json:"id"`
	Name            string    `json:"name"`
	Score           float32   `json:"score"`
	PublicationDate time.Time `json:"publicationDate"`
	Type            idName    `json:"type"`
	Author          idName    `json:"author"`
	Publisher       idName    `json:"publisher"`
}

func BookResponseFromBook(b entities.Book) bookResponse {
	return bookResponse{
		b.ID(),
		b.Name,
		b.Score,
		b.PublicationDate,
		idName{b.Type.ID(), b.Type.Name},
		idName{b.Author.ID(), b.Author.FullName},
		idName{b.Publisher.ID(), b.Publisher.Name},
	}
}
