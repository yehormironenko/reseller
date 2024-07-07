package repository

import (
	"github.com/labstack/echo/v4"

	"github.com/yehormironenko/reseller/internal/repository/entities"
	"github.com/yehormironenko/reseller/pkg/model"
)

type BookRepositoryInterface interface {
	GetBookByParams(c echo.Context, bookname, author, genre string) (entities.Books, error)
	BuyBook(c echo.Context, books model.Books) (entities.Books, error)
}
