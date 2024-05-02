package service

import (
	"net/url"

	"github.com/labstack/echo/v4"

	"github.com/yehormironenko/reseller/pkg/model"
)

type Action interface {
	GetBook(c echo.Context, args url.Values) (model.Books, error)
}
