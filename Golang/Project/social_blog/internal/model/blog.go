package model

type Blog struct {
	ID         int    `json:"id" gorm:"primary_key"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	Published  bool   `json:"published"`
	Visibility string `json:"visibility"`
	Base
}

const (
	BlogVisibilityPublic  = "PUBLIC"
	BlogVisibilityPrivate = "PRIVATE"
)
