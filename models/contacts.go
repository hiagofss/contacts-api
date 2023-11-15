package models

type Contact struct {
	Id    string `gorm:"default:uuid_generate_v4()" json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Email string `json:"email"`
}

var Contacts []Contact
