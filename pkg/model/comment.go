package model

import (
	"social-media-app/pkg/dtos"

	"github.com/google/uuid"
)

type Comment struct {
	Base
	UserID  uuid.UUID `json:"user_id"`
	PostID  uuid.UUID `json:"post_id"`
	Text    string    `json:"text"`
	ReplyTo uuid.UUID `json:"reply_to,omitempty"`
}

func (c *Comment) Mapper(req *dtos.UpdateCommentDTO) {
	c.Text = req.Text
}
