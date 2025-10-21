package model

type User struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	Email     string `gorm:"unique;not null" json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
