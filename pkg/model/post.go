package model

import (
	"social-media-app/pkg/dtos"

	"github.com/google/uuid"
)

type Post struct {
	Base
	UserID    uuid.UUID `json:"user_id"`
	Caption   string    `json:"caption"`
	ShortLink string    `gorm:"uniqueIndex;not null" json:"short_link"`
}

func (p *Post) Mapper(req *dtos.AddPostDTO) {
	p.Caption = req.Caption
	// p.ShortLink = req.ShortLink
}

func (p *Post) MapperForUpdate(req *dtos.UpdatePostDTO) {
	p.Caption = req.Caption
	// p.ShortLink = req.ShortLink
}
