package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/yehormironenko/reseller/internal/service"
)

// /search?book=book-name&author=x-author&genre=fiction
type HandlerGetBook struct {
	BookService service.BookService
}

func (h HandlerGetBook) GetBook(c echo.Context) error {
	params := c.QueryParams()

	response, err := h.BookService.GetBook(c, params)
	if err != nil {
		return c.String(http.StatusInternalServerError, "internal error")
	}

	return c.JSON(http.StatusOK, response)
}
