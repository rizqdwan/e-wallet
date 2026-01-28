package repository

import (
	"context"
	"database/sql"

	"github.com/rizqdwan/e-wallet/internal/domain"
)


type WalletRepostory interface {
	GetWalletByID(ctx context.Context, tx *sql.Tx, walletID int64) (*domain.Wallet, error)
	GetWalletByUserID(ctx context.Context, tx *sql.Tx, userID int64) (*domain.Wallet, error)

	GetWalletByIDForUpdate(ctx context.Context, tx *sql.Tx, id int64) (*domain.Wallet,  error)
	UpdateWalletBalance(ctx context.Context, tx *sql.Tx, walletID int64, newBalance int64) error

	CreateWallet(ctx context.Context, tx *sql.Tx, wallet *domain.Wallet) error
}

type walletRepository struct {
	db *sql.DB
}


func (r *walletRepository) GetWalletByID(ctx context.Context, tx *sql.Tx, walletID int64) (*domain.Wallet, error) {
	w := domain.Wallet{}

	query := "SELECT id, user_id, balance, wallet_number, created_at FROM wallet WHERE id = $1"
	row := tx.QueryRowContext(ctx, query, walletID)
	if err := row.Scan(
		&w.ID,
		&w.UserID,
		&w.Balance,
		&w.WalletNumber,
		&w.CreatedAt,
	); err != nil {
		return nil, err
	}

	return &w, nil
}


func (r *walletRepository) GetWalletByUserID(ctx context.Context, tx *sql.Tx, userID int64) (*domain.Wallet, error) {

	query := "SELECT id, user_id, balance, wallet_number, created_at FROM wallet WHERE user_id = $1"
	row := tx.QueryRowContext(ctx, query, userID)
	var w domain.Wallet
	if err := row.Scan(
		&w.ID,
		&w.UserID,
		&w.Balance,
		&w.WalletNumber,
		&w.CreatedAt,
	); err != nil {
		return nil, err
	}

	return &w, nil
}


func (r *walletRepository) GetWalletByIDForUpdate(ctx context.Context, tx *sql.Tx, id int64) (*domain.Wallet,  error) {

	query := "SELECT id, user_id, balance, wallet_number, created_at FROM wallet WHERE id = $1 FOR UPDATE"
	row := tx.QueryRowContext(ctx, query, id)
	var w domain.Wallet
	if err := row.Scan(
		&w.ID,
		&w.UserID,
		&w.Balance,
		&w.WalletNumber,
		&w.CreatedAt,
	); err != nil {
		return nil, err
	}

	return &w, nil
}

func (r *walletRepository) UpdateWalletBalance(ctx context.Context, tx *sql.Tx, walletID int64, newBalance int64) error {

	query := "UPDATE wallet SET balance = $1 WHERE id = $2"
	_, err := tx.ExecContext(ctx, query, newBalance, walletID)

	return err
}



func (r *walletRepository) Create(ctx context.Context, tx *sql.Tx, wallet *domain.Wallet) error {

	query := "INSERT INTO wallet (user_id, balance, wallet_number) VALUES ($1, $2, $3)"
	_, err := tx.ExecContext(ctx, query, wallet.UserID, wallet.Balance, wallet.WalletNumber)

	return err
}

