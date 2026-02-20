package dto

import (
	"time"
)

type CreatePostRequest struct {
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	UserID    string `json:"user_id"`
	JobID     *int   `json:"job_id,omitempty"`
	CompanyID *int   `json:"company_id,omitempty"`
	Images    []string `json:"images"`
}

type PostResponseDTO struct {
	ID           uint                  `json:"id"`
	Title        string                `json:"title"`
	Content      string                `json:"content"`
	CreatedAt    time.Time             `json:"created_at"`
	User         UserShortDTO          `json:"user"`
	Job          *JobShortDTO          `json:"job,omitempty"`
	Company      *CompanyShortDTO      `json:"company,omitempty"`
	Images       []string              `json:"images"` // Solo las URLs
	Interactions []InteractionShortDTO `json:"interactions"`
	Comments     []CommentShortDTO     `json:"comments"`
}

//Todos estos son DTO para mapear y poder mostrar de mejor manera los post en json 
type UserShortDTO struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type JobShortDTO struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

type CompanyShortDTO struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Location string `json:"location"`
}

type InteractionShortDTO struct {
	ID       int    `json:"id"`
	UserName string `json:"user_name"`
	Reaction string `json:"reaction"`
}

type CommentShortDTO struct {
	ID        uint   `json:"id"`
	Content   string `json:"content"`
	UserName  string `json:"user_name"`
	CreatedAt int64  `json:"created_at"`
}
