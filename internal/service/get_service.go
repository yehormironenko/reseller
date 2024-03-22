package service

import (
	"net/url"

	"github.com/labstack/echo/v4"
)

func GetBook(c echo.Context, args url.Values) (*string, error) {
	c.Logger().Info("getBook service")

	//to repository

	return nil, nil
}
