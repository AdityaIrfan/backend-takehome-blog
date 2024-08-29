package models

import "time"

type Comment struct {
	ID        string    `gorm:"column:id;primary_key"`
	PostID    string    `gorm:"column:post_id;not null"`
	AuthorID  string    `gorm:"column:author_id;type:varchar(100);not null"`
	Content   string    `gorm:"column:content;type:text;not null"`
	ParentID  *string   `gorm:"column:parent_id;default null"`
	CreatedAt time.Time `gorm:"column:created_at;type:timestamp"`

	Author User `gorm:"foreignKey:AuthorID;references:ID"`
}

func (c *Comment) ToResponse() *CommentResponse {
	return &CommentResponse{
		ID:         c.ID,
		Content:    c.Content,
		AuthorID:   c.Author.ID,
		AuthorName: c.Author.Name,
	}
}

type CommentWriteRequest struct {
	PostID   string
	AuthorId string
	Content  string `json:"Content" form:"Content" validate:"required"`
	ParentID string `json:"ParentId" form:"ParentId"`
}

type CommentResponse struct {
	ID           string `json:"Id"`
	Content      string `json:"Content"`
	AuthorID     string `json:"AuthorId"`
	AuthorName   string `json:"AuthorName"`
	TotalComment uint32 `json:"TotalComment"`
}
