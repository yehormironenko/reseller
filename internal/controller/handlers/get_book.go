package handlers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"reseller/internal/service"
)

// /search?book=book-name&author=x-author&genre=fiction

func GetBook(c echo.Context) error {
	book := c.QueryParam("book")
	author := c.QueryParam("author")
	genre := c.QueryParam("genre")
	params := c.QueryParams()

	response, err := service.GetBook(c, params)
	if err != nil {
		return c.String(http.StatusInternalServerError, "internal error")
	}

	fmt.Println(response)

	return c.String(http.StatusOK, fmt.Sprintf("book %s, author %s, genre %s", book, author, genre))
}