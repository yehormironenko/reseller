package repository

import (
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/uptrace/bun"

	"reseller/internal/repository/entities"
)

type BookRepository struct {
	db *bun.DB
}

func NewBookRepository(db *bun.DB) BookRepository {
	return BookRepository{
		db: db,
	}
}

func (br BookRepository) GetBookByParams(c echo.Context, bookname, author, genre string) (entities.Books, error) {
	c.Logger().Info("GetBookByParams repository")
	c.Logger().Info("bookname:", bookname)
	c.Logger().Info("author:", author)
	c.Logger().Info("genre:", genre)

	books := new(entities.Books)
	model := br.db.NewSelect().Model(books)

	query := br.constructSelect(model, bookname, author, genre)
	err := query.Scan(c.Request().Context())
	if err != nil {
		c.Logger().Errorf("Failed to scan books: %v", err)
		return nil, err
	}

	c.Logger().Info("found books", books)

	return *books, nil
}

func (br BookRepository) constructSelect(model *bun.SelectQuery, bookname, author, genre string) *bun.SelectQuery {

	if hasParam(bookname) {
		model.Where("lower(book_name) = ?", strings.ToLower(bookname))
	}

	if hasParam(author) {
		model.Where("lower(author) = ?", strings.ToLower(author))
	}

	if hasParam(genre) {
		model.Where("lower(genre) = ?", strings.ToLower(genre))
	}
	return model
}

func hasParam(parameter string) bool {
	return parameter != ""
}
