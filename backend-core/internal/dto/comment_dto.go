package dto

import "time"

type CreateCommentReq struct {
    Content string `json:"content" binding:"required"`
    PostID  uint   `json:"post_id" binding:"required"` // <--- AGREGADO
}

type UpdateCommentReq struct {
    Content string `json:"content" binding:"required"`
}

type CommentResponse struct {
    ID        uint      `json:"id"`
    Content   string    `json:"content"`
    PostID    uint      `json:"post_id"`
    CreatedAt time.Time `json:"created_at"`
    User      UserShortResponse `json:"user"`
}

type UserShortResponse struct {
    ID       uint   `json:"id"`
    FullName string `json:"full_name"`
    Avatar   string `json:"avatar,omitempty"`
}