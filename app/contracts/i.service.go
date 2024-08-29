package contract

import (
	"backend-takehome-blog/helpers"
	"backend-takehome-blog/models"

	"github.com/labstack/echo/v4"
)

type IAuthService interface {
	Register(c echo.Context, in *models.Register) error
	Login(c echo.Context, in *models.Login) error
}

type IPostService interface {
	Create(c echo.Context, in *models.PostCreateRequest) error
	Update(c echo.Context, in *models.PostUpdateRequest) error
	GetAllMine(c echo.Context, authorId string, cursor *helpers.Cursor) error
	GetAll(c echo.Context, cursor *helpers.Cursor) error
	GetDetail(c echo.Context, postId string) error
	Delete(c echo.Context, postId string) error
}

type ICommentService interface {
	Create(c echo.Context, in *models.CommentWriteRequest) error
	GetAllByPost(c echo.Context, postId, parentId string, cursor *helpers.Cursor) error
}
