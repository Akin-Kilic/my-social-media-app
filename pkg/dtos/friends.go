package dtos

import "github.com/google/uuid"

type AddFriendDTO struct {
	FriendId uuid.UUID `json:"friend_id"`
}

type ChangeStatusDTO struct {
	FriendId     uuid.UUID `json:"friend_id"`
	ChangeStatus string    `json:"change_status"`
}
