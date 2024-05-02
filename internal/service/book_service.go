package service

import (
	"net/url"

	"github.com/labstack/echo/v4"

	"reseller/internal/repository"
	"reseller/internal/repository/entities"
	"reseller/pkg/model"
)

type BookService struct {
	bookrepository repository.BookRepository
}

func NewBookService(bookRepository repository.BookRepository) BookService {
	return BookService{
		bookrepository: bookRepository,
	}
}

func (bs BookService) GetBook(c echo.Context, arguments url.Values) (model.Books, error) {
	c.Logger().Info("getBook service")
	bookName, author, genre := getArgs(arguments)
	books, err := bs.bookrepository.GetBookByParams(c, bookName, author, genre)

	if err != nil {
		return model.Books{}, err
	}

	return convertBooksToClientModel(books), nil
}

func convertBooksToClientModel(books []entities.Book) []model.Book {
	clientBooks := make([]model.Book, 0, len(books)) // Pre-allocate slice with capacity equal to the number of books

	for _, book := range books {
		clientBook := model.Book{
			Bookname: book.Bookname,
			Author:   book.Author,
			Genre:    book.Genre,
			Price:    book.Price,
			Year:     book.Year,
		}
		clientBooks = append(clientBooks, clientBook)
	}
	return clientBooks
}

func getArgs(arguments url.Values) (string, string, string) {
	// TODO to constants this is temporary in the future another approach
	bookName := arguments.Get("book")
	author := arguments.Get("author")
	genre := arguments.Get("genre")

	return bookName, author, genre
}
