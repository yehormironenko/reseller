package repository

import (
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/uptrace/bun"

	"github.com/yehormironenko/reseller/internal/repository/entities"
	"github.com/yehormironenko/reseller/pkg/model"
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
	selectQuery := br.db.NewSelect().Model(books)

	query := br.constructSelect(selectQuery, bookname, author, genre)
	err := query.Scan(c.Request().Context())
	if err != nil {
		c.Logger().Errorf("Failed to scan books: %v", err)
		return nil, err
	}

	c.Logger().Info("found books", books)

	return *books, nil
}

func (br BookRepository) BuyBook(c echo.Context, books model.Books) (entities.Books, error) {

	c.Logger().Info("BuyBook repository")
	c.Logger().Info("books", books)

	purchasedBooks := entities.Books{}

	for _, book := range books {
		findBooks, err := br.GetBookByParams(c, book.Bookname, book.Author, book.Genre)
		if err != nil {
			c.Logger().Errorf("Update operation failing: %v", err)
			return nil, err
		}
		if findBooks != nil {
			purchasedBooks = append(purchasedBooks, findBooks[0])
		}
	}

	values := br.db.NewValues(&purchasedBooks)
	q := br.db.NewUpdate().
		With("_data", values).
		Model((*entities.Book)(nil)).
		Table("_data").
		Set("count = _data.count -1").
		Where("b.book_name = _data.book_name")

	if _, err := q.Exec(c.Request().Context()); err != nil {
		c.Logger().Errorf("Update operation failing: %v", err)
		return nil, err
	}

	return purchasedBooks, nil
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

	model.Where("count > 0")
	return model
}

func hasParam(parameter string) bool {
	return parameter != ""
}
