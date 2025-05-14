package models

type User struct {
	ID        uint     `gorm:"primaryKey" json:"id"`
	Email     string   `gorm:"unique;not null" json:"email"`
	Username  string   `gorm:"unique;not null" json:"username"`
	Password  string   `gorm:"not null" json:"-"`
	Latitude  *float64 `json:"latitude"`
	Longitude *float64 `json:"longitude"`
	Interests string   `json:"interests"` // 以逗號分隔
}
