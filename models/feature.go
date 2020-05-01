package models

type Feature struct {
	ID        uint   `json:"id" gorm:"primary_key"`
	Title     string `json:"title"`
	Details   string `json:"details"`
	UpVotes   uint   `json:"upvotes"`
	CreatedAt int64  `json:"createdAt"`
}
