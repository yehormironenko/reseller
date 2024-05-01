package repository

import (
	"github.com/labstack/echo/v4"

	"reseller/internal/repository/entities"
)

type BookRepositoryInterface interface {
	GetBookByParams(c echo.Context, bookname, author, genre string) (entities.Book, error)
	// TODO buy
}
