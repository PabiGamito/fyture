package models

type Feature struct {
	ID        uint   `json:"id" gorm:"primary_key"`
	Title     string `json:"title" gorm:"not null"`
	Details   string `json:"details" gorm:"not null"`
	UpVotes   uint   `json:"upvotes" gorm:"not null"`
	CreatedAt int64  `json:"createdAt" gorm:"not null"`
	CreatedBy User   `json:"createdBy" gorm:"foreign_key:UserID;not null"`
	UserID    uint   `json:"userId" gorm:"not null"`
}
