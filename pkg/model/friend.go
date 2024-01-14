package model

import (
	"social-media-app/pkg/dtos"

	"github.com/google/uuid"
)

type Friend struct {
	Base
	UserId   uuid.UUID `json:"user_id"`
	FriendId uuid.UUID `json:"friend_id"`
	Status   string    `json:"status" gorm:"default:1"` // 1:pending, 2:accepted, 3:rejected, 4:deleted
}

// TODO: error u kaldır
func (f *Friend) Mapper(req *dtos.AddFriendDTO) {
	f.FriendId = req.FriendId
	f.Status = "1"
}

func (f *Friend) AcceptFriendMapper(userId, friendId uuid.UUID) {
	f.UserId = userId
	f.FriendId = friendId
	f.Status = "2"
}
