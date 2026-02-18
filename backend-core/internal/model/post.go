package model

import (
	"time"

	"gorm.io/gorm"
)

type Post struct {
	ID           uint   `gorm:"primaryKey;autoIncrement"`
	Title        string `gorm:"type:text;not null"`
	Content      string `gorm:"type:text;not null"`
	UserID       string `gorm:"not null"`
	User         User   `gorm:"foreignKey:UserID"`
	JobID        *int
	Job          Job `gorm:"foreignKey:JobID"`
	CompanyID    *int
	Company      Company        `gorm:"foreignKey:CompanyID"`
	Comments     []Comment      `gorm:"foreignKey:PostID"`
	Interactions []Interaction  `gorm:"foreignKey:PostID"`
	Images       []Image        `gorm:"many2many:post_images;"`
	CreatedAt    time.Time      `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt    time.Time      `gorm:"autoUpdateTime"`
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}

type Image struct {
	ID        int       `gorm:"primaryKey"`
	Name      string    `gorm:"type:text"`
	URL       string    `gorm:"not null"`
	Caption   string    `gorm:"type:text" json:"caption"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}

type PostImage struct {
	PostID    int       `gorm:"primaryKey"`
	ImageID   uint      `gorm:"primaryKey"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}
