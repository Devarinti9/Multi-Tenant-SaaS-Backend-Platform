package models

import "time"

type Organization struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"uniqueIndex;not null" json:"name"`
	CreatedAt time.Time `json:"created_at"`
	Users     []User    `json:"users"`
}

type User struct {
	ID             uint      `gorm:"primaryKey" json:"id"`
	OrganizationID uint      `json:"organization_id"`
	Name           string    `json:"name"`
	Email          string    `gorm:"uniqueIndex;not null" json:"email"`
	PasswordHash   string    `json:"-"`
	Role           string    `json:"role"`
	CreatedAt      time.Time `json:"created_at"`
}
