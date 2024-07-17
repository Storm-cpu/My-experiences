package model

type Comment struct {
	ID      int    `json:"id" gorm:"primary_key"`
	BlogID  int    `json:"blog_id" gorm:"not null"`
	UserID  int    `json:"user_id" gorm:"not null"`
	Content string `json:"name" gorm:"type:varchar(255);unique_index;not null"`
	Base
}
