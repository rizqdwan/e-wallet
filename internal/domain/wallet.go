package domain

import "time"

type Wallet struct {
	ID           int64 `json:"id" db:"id"`
	UserID       int64 `json:"user_id" db:"user_id"`
	Balance      int64 `json:"balance" db:"balance"`
	WalletNumber string `json:"wallet_number" db:"wallet_number"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
}