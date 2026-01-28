package repository

import (
	"context"
	"database/sql"

	"github.com/rizqdwan/e-wallet/internal/domain"
)

type TransactionRepository interface {
	CreateTransaction(ctx context.Context, tx *sql.Tx, trx *domain.Transaction) error
	GetTransactionByID(ctx context.Context, tx *sql.Tx, id string) (*domain.Wallet, error)
	GetTransactionByReference(ctx context.Context, tx *sql.Tx, reference string) (*domain.Wallet, error)
	ListTransactionByWalletID(ctx context.Context, tx *sql.Tx, walletID string) ([]*domain.Transaction, error)
	UpdateTransactionStatus(ctx context.Context, tx *sql.Tx, trx *domain.Transaction) error
}


type transactionRepository struct {
	db *sql.DB
}


func (t *transactionRepository) CreateTransaction(ctx context.Context, tx *sql.Tx, trx *domain.Transaction) error {
	
	query := "INSERT INTO transactions (wallet_id, amount, reference, status) VALUES ($1, $2, $3, $4)"
	_, err := tx.ExecContext(ctx, query, trx.WalletID, trx.Amount, trx.ReferenceID, trx.Status)
	if err != nil {
		return err
	}
	
	return nil
}