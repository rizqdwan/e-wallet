package domain

import (
	"time"

	"github.com/google/uuid"
)

type TransactionType string
type Status string

const (
	Deposit    TransactionType = "deposit"
	Withdrawal TransactionType = "withdrawal"
	Transfer   TransactionType = "transfer"
)

const (
	Pending   Status = "pending"
	Completed Status = "completed"
)

type Transaction struct {
	ID              uuid.UUID 			`json:"id" db:"id"`
	WalletID        int64						`json:"wallet_id" db:"wallet_id"`
	TransactionType TransactionType	`json:"transaction_type" db:"transaction_type"`
	Amount          int64						`json:"amount" db:"amount"`
	Status          Status					`json:"status" db:"status"`
	ReferenceID     string					`json:"reference_id" db:"reference_id"`
	IdempotencyKey  string					`json:"idempotency_key" db:"idempotency_key"`
	Description     string					`json:"description" db:"description"`
	CreatedAt       time.Time				`json:"created_at" db:"created_at"`
}
