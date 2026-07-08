package model

import (
	"time"

	"github.com/uptrace/bun"
)

type Car struct {
	bun.BaseModel `bun:"table:cars"`

	CarID              int64      `bun:"car_id,pk,autoincrement" json:"car_id"`
	RegistrationNumber string     `bun:"registration_number,notnull" json:"registration_number"`
	Brand              string     `bun:"brand,notnull" json:"brand"`
	Model              string     `bun:"model,notnull" json:"model"`
	Color              string     `bun:"color" json:"color"`
	Year               int        `bun:"year" json:"year"`
	Notes              string     `bun:"notes" json:"notes"`
	CreatedAt          time.Time  `bun:"created_at,notnull" json:"created_at"`
	UpdatedAt          time.Time  `bun:"updated_at,notnull" json:"updated_at"`
	DeletedAt          *time.Time `bun:"deleted_at" json:"deleted_at,omitempty"`
}
