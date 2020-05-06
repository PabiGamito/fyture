package models

type User struct {
	ID           uint   `json:"id" gorm:"primary_key"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	Email        string `json:"email" gorm:"unique_index"`
	ProfileImage string `json:"profileImage"`
}
