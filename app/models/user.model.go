package models

import "time"

type User struct {
	ID           string     `gorm:"column:id;primary_key"`
	Name         string     `gorm:"column:name;type:varchar(100)"`
	Email        string     `gorm:"column:email;type:varchar(100);unique_index"`
	PasswordHash string     `gorm:"column:password_hash;type:varchar(255)"`
	PasswordSalt string     `gorm:"column:password_salt;type:varchar(255)"`
	CreatedAt    time.Time  `gorm:"column:created_at;type:timestamp"`
	UpdatedAt    *time.Time `gorm:"column:updated_at;type:timestamp;<-:update"`
}
