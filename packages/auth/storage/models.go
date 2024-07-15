package storage

import (
	"time"
)

type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Username  string    `gorm:"uniqueIndex:idx_user_username;not null" json:"username"`
	Email     string    `gorm:"unique;not null" json:"email"`
	Password  string    `gorm:"not null" json:"password"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP;not null" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP;not null" json:"updated_at"`
	URLs      []URL     `gorm:"foreignKey:UserUsername;references:Username" json:"urls"`
}

func (User) TableName() string {
	return "user"
}

type URL struct {
	ID           uint   `gorm:"primaryKey" json:"id"`
	UserUsername string `gorm:"index:idx_url_username;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"user_username"`
	Short        string `gorm:"uniqueIndex:idx_url_short;not null" json:"short"`
	Long         string `gorm:"not null" json:"long"`
}

func (URL) TableName() string {
	return "url"
}
