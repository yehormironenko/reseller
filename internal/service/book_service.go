package service

import (
	"fmt"
	"net/url"

	"github.com/labstack/echo/v4"

	"reseller/internal/model"
	"reseller/internal/repository"
)

type BookService struct {
	bookrepository repository.BookRepository
}

func NewBookService(bookRepository repository.BookRepository) BookService {
	return BookService{
		bookrepository: bookRepository,
	}
}

func (bs BookService) GetBook(c echo.Context, arguments url.Values) (model.Book, error) {
	c.Logger().Info("getBook service")

	//to repository
	fmt.Println(arguments)
	bs.bookrepository.GetBookByParams("1", "2")
	return model.Book{}, nil
}
