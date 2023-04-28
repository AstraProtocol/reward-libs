package model

import "time"

// Base entity only created time
type BaseEntity struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time `json:"createdAt"`
}

// Base entity with updated time
type BasicEntity struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
