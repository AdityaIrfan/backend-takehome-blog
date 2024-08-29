package contract

import (
	"github.com/labstack/echo/v4"
)

type IAuthHandler interface {
	Register(c echo.Context) error
	Login(c echo.Context) error
}

type IPostHandler interface {
	Create(c echo.Context) error
	Update(c echo.Context) error
	GetAllMine(c echo.Context) error
	GetAll(c echo.Context) error
	GetDetail(c echo.Context) error
	Delete(c echo.Context) error
}

type ICommentHandler interface {
	Create(c echo.Context) error
	GetAllByPost(c echo.Context) error
}
