package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h HandlerWithService) GetBook(c echo.Context) error {
	params := c.QueryParams()

	response, err := h.BookService.GetBook(c, params)
	if err != nil {
		return c.String(http.StatusInternalServerError, "internal error")
	}

	return c.JSON(http.StatusOK, response)
}
