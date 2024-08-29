package contract

import (
	"backend-takehome-blog/helpers"
	"backend-takehome-blog/models"
)

type IAuthRepository interface {
}

type IUserRepository interface {
	GetByCustomAndSelectedFields(values map[string]interface{}, selectedFields string) (*models.User, error)
	Create(user *models.User) error
}

type IPostRepository interface {
	Create(post *models.Post) error
	GetByCustomAndSelectedFields(values map[string]interface{}, selectedFields string, preloads ...string) (*models.Post, error)
	GetAllByCursorAndSelectedFieldsPaginate(
		values map[string]interface{},
		cursor *helpers.Cursor,
		selectedFields string,
		preloads ...string) ([]*models.Post, *helpers.CursorPagination, error)
	Update(post *models.Post) error
	Delete(post *models.Post) error
}

type ICommentRepository interface {
	Create(comment *models.Comment) error
	GetByCustomAndSelectedFields(values map[string]interface{}, selectedFields string, preloads ...string) (*models.Comment, error)
	GetAllByCursorAndSelectedFieldsPaginate(
		values map[string]interface{},
		cursor *helpers.Cursor,
		selectedFields string,
		preloads ...string) ([]*models.Comment, *helpers.CursorPagination, error)
	GetTotalByCustom(values map[string]interface{}) (int64, error)
}
