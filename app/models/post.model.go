package models

import (
	"backend-takehome-blog/helpers"
	"time"

	"gorm.io/gorm"
)

type Post struct {
	ID        string          `gorm:"column:id;primary_key"`
	Title     string          `gorm:"column:title;type:varchar(255)"`
	Content   string          `gorm:"column:content;type:text"`
	AuthorID  string          `gorm:"column:author_id;not null"`
	CreatedAt time.Time       `gorm:"column:created_at;type:timestamp"`
	UpdatedAt *time.Time      `gorm:"column:updated_at;type:timestamp;<-:update"`
	DeletedAt *gorm.DeletedAt `gorm:"column:deleted_at;type:timestamp"`

	Author User `gorm:"foreignKey:AuthorID;references:ID"`
}

func (p *Post) ToResponse() *PostResponse {
	return &PostResponse{
		ID:        p.ID,
		Title:     p.Title,
		Content:   p.Content,
		CreatedAt: p.CreatedAt.Format(helpers.DefaultTimeFormat),
	}
}

type PostCreateRequest struct {
	ID       string
	Title    string `json:"Title" form:"Title" validate:"required,min=1,max=255"`
	Content  string `json:"Content" form:"Content" validate:"required"`
	AuthorID string
}

type PostUpdateRequest struct {
	ID       string
	Title    string `json:"Title" form:"Title" validate:"max=255"`
	Content  string `json:"Content" form:"Content"`
	AuthorID string
}

func (p *Post) ToResponseList() *PostResponse {
	return &PostResponse{
		ID:        p.ID,
		Title:     p.Title,
		CreatedAt: p.CreatedAt.Format(helpers.DefaultTimeFormat),
	}
}

type PostResponse struct {
	ID            string `json:"Id"`
	Title         string `json:"Title"`
	Content       string `json:"Content,omitempty"`
	CreatedAt     string `json:"CreatedAt"`
	AuthorID      string `json:"AuthorId,omitempty"`
	AuthorName    string `json:"Author,omitempty"`
	TotalComments uint32 `json:"TotalComments"`
}
