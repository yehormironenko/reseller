package handlers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

// /search?book=book-name&author=x-author&genre=fiction
func GetBook(c echo.Context) error {
	book := c.QueryParam("book")
	author := c.QueryParam("author")
	genre := c.QueryParam("genre")

	// to service -> to repo

	return c.String(http.StatusOK, fmt.Sprintf("book %s, author %s, genre %s", book, author, genre))
}
