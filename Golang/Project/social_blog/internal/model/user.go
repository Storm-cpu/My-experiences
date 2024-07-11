package model

type User struct {
	ID          int    `json:"id" gorm:"primary_key"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	PhoneNumber string `json:"phone_number"`
	Address     string `json:"address"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	RoleID      int    `json:"role_id"`
	Status      string `json:"status"`
	Base
}

const (
	UserStatusDeletedByAdmin = "DELETE_BY_ADMIN"
)
