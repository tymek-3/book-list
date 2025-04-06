package data

import (
	"book-list/internal/domain/entities"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

func (u User) ToEntity() *entities.User {
	return entities.ConstructUser(
		entities.Email(u.Email),
		u.Name,
		u.PasswordHash,
		entities.RoleFromString(u.Role),
		// TODO:
		[]uuid.UUID{},
	)
}

func (u BookGetFullByIdRow) ToEntity() *entities.Book {
	return constructFullBook(u.ID, u.Name, u.Score, u.PublicationDate, u.TypeID, u.AuthorID, u.PublisherID, u.TypeName, u.AuthorFullName, u.PublisherName)
}

func (u BookFullSearchRow) ToEntity() *entities.Book {
	return constructFullBook(u.ID, u.Name, u.Score, u.PublicationDate, u.TypeID, u.AuthorID, u.PublisherID, u.TypeName, u.AuthorFullName, u.PublisherName)
}

func (u BookFullAllRow) ToEntity() *entities.Book {
	return constructFullBook(u.ID, u.Name, u.Score, u.PublicationDate, u.TypeID, u.AuthorID, u.PublisherID, u.TypeName, u.AuthorFullName, u.PublisherName)
}

func constructFullBook(id uuid.UUID, name string, score sql.NullFloat64, publicationDate sql.NullString, typeID, authorID, publisherID any, typeName, authorFullName, publisherName sql.NullString) *entities.Book {
	pubDate, err := time.Parse(time.DateOnly, publicationDate.String)
	if err != nil {
		pubDate = time.Unix(0, 0)
	}
	return entities.ConstructBook(
		id,
		name,
		float32(score.Float64),
		pubDate,
		*entities.ConstructBookType(uuid.MustParse(typeID.(string)), typeName.String),
		*entities.ConstructAuthor(uuid.MustParse(authorID.(string)), authorFullName.String),
		*entities.ConstructPublisher(uuid.MustParse(publisherID.(string)), publisherName.String),
	)
}
