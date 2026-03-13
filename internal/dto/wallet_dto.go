package dto

import (
	"time"

	"github.com/google/uuid"
	"github.com/rizqdwan/e-wallet/internal/domain"
)

type WalletResponse struct {
	ID           uuid.UUID `json:"id"`
	WalletNumber string    `json:"wallet_number"`
	Balance      int64     `json:"balance"`
	IsActive     bool      `json:"is_active"`
}

type TransactionResponse struct {
	ID              uuid.UUID              `json:"id"`
	TransactionType domain.TransactionType `json:"transaction_type"`
	Status          domain.Status          `json:"status"`
	Amount          int64                  `json:"amount"`
	ReferenceID     *string                `json:"reference_id"`
	Description     *string                `json:"description"`
	CreatedAt       time.Time              `json:"created_at"`
}

type TopUpRequest struct {
	Amount      int64  `json:"amount"       validate:"required,gt=0"`
	Description string `json:"description"  validate:"max=255"`
}

type TransferRequest struct {
	ToWalletNumber string `json:"to_wallet_number" validate:"required"`
	Amount         int64  `json:"amount"           validate:"required,gt=0"`
	Description    string `json:"description"      validate:"max=255"`
}
