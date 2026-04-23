package models

import "time"

type Organization struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"uniqueIndex;not null" json:"name"`
	Slug      string    `gorm:"uniqueIndex;not null" json:"slug"`
	Plan      string    `json:"plan"`
	CreatedAt time.Time `json:"created_at"`
	Users     []User    `json:"users,omitempty"`
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

type ProjectMembership struct {
	ID             uint      `gorm:"primaryKey" json:"id"`
	OrganizationID uint      `json:"organization_id"`
	ProjectName    string    `json:"project_name"`
	Environment    string    `json:"environment"`
	StoragePath    string    `json:"storage_path"`
	CreatedAt      time.Time `json:"created_at"`
}
