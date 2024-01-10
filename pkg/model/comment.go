package model

type Comment struct {
	Base
	UserID  uint   `json:"user_id"`
	User    User   `json:"user"`
	PostID  uint   `json:"post_id"`
	Post    Post   `json:"post"`
	Text    string `json:"text"`
	Image   string `json:"image"`
	Likes   []User `gorm:"many2many:comment_likes;" json:"likes"`
	ReplyTo uint   `json:"reply_to,omitempty"`
}
