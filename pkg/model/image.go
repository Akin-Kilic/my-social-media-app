package model

import "github.com/google/uuid"

type Image struct {
	Base
	Data      string    `json:"image"`
	UserId    uuid.UUID `json:"user_id"`
	EnityId   uuid.UUID `json:"entity_id"`
	ImageType string    `json:"image_type"` // 1:profile, 2:post, 3:comment
}

func (i *Image) Mapper(data, imageType string, userId, entityId uuid.UUID) {
	i.Data = data
	i.UserId = userId
	i.EnityId = entityId
	i.ImageType = imageType
}
