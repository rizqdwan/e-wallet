package domain

import (
	"time"

	"github.com/google/uuid"
)

type Wallet struct {
	ID           uuid.UUID `json:"id" db:"id"`
	UserID       uuid.UUID `json:"user_id" db:"user_id"`
	Balance      int64     `json:"balance" db:"balance"`
	WalletNumber string    `json:"wallet_number" db:"wallet_number"`
	IsActive     bool      `json:"is_active" db:"is_active"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
}
