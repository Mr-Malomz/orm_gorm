package models

type User struct {
	Id        int    `json:"id" gorm:"primaryKey"`
	FirstName string `json:"firstName,omitempty" validate:"required"`
	LastName  string `json:"lastName,omitempty" validate:"required"`
	Title     string `json:"title,omitempty" validate:"required"`
}

