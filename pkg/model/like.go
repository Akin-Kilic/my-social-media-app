package model

import "github.com/google/uuid"

type Likes struct {
	Base
	UserId    uuid.UUID `json:"like_user_id"`
	PostId    uuid.UUID `json:"post_id"`
	CommentId uuid.UUID `json:"comment_id"`
	Type      string    `json:"like_type"` // 1: Post, 2: Comment
}
