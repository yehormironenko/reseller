package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/yehormironenko/reseller/pkg/model"
)

func (h HandlerWithService) BuyBook(c echo.Context) error {

	books := model.Books{}

	if err := c.Bind(&books); err != nil {
		return err
	}

	response, err := h.BookService.BuyBook(c, books)
	if err != nil {
		return c.String(http.StatusInternalServerError, "internal error")
	}

	return c.JSON(http.StatusOK, response)
}
