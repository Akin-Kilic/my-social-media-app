package model

import (
	"time"

	"github.com/google/uuid"
)

type Post struct {
	Base
	UserID    uint       `json:"user_id"`
	Caption   string     `json:"caption"`
	ShortLink string     `gorm:"uniqueIndex;not null" json:"short_link"`
	ExpiresAt *time.Time `json:"expires_at"`
}

type Likes struct {
	Base
	EntityId   uuid.UUID `json:"entity_id"`
	LikeUserId uuid.UUID `json:"like_user_id"`
	LikeType   int       `json:"like_type"` // 1: Post, 2: Comment
}

type UserPost struct {
	Base
	UserId uuid.UUID `json:"user_id"`
	PostId uuid.UUID `json:"post_id"`
}
