package model

import "time"

type Pastel struct {
	ID          string       `gorm:"primaryKey;size:50"`
	Name        string       `gorm:"not null;unique;size:50" sql:"index"`
	Price       float32      `gorm:"not null"`
	Ingredients []Ingredient `gorm:"many2many:pasteis_ingredients;"`
	CreatedAt   *time.Time   `gorm:"default:current_timestamp"`
	UpdatedAt   *time.Time
}

func (Pastel) TableName() string {
	return "pasteis"
}
