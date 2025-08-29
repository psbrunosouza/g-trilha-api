package entity

import (
	"time"

	"github.com/google/uuid"
)

type AccountEntity struct {
	ID        uuid.UUID
	Name      string
	Email     string
	Password  string
	Avatar    string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
