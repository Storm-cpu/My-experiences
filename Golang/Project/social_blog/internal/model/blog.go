package model

type Blog struct {
	ID         int    `json:"id" gorm:"primary_key"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	AuthorID   int    `json:"author_id" gorm:"not null"`
	Status     string `json:"status" gorm:"type:varchar(50), not null"`
	Visibility string `json:"visibility" gorm:"type:varchar(50);not null"`
	Base
}

const (
	BlogVisibilityPublic      = "PUBLIC"
	BlogVisibilityPrivate     = "PRIVATE"
	BlogVisibilityFriendsOnly = "FRIENDS_ONLY"
)

const (
	UserStatusPending  = "PENDING"
	UserStatusApproved = "APPROVED"
	UserStatusRejected = "REJECTED"
)
