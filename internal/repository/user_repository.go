package repository

import (
	"context"
	"database/sql"

	"github.com/rizqdwan/e-wallet/internal/domain"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *domain.User) error

	GetUserById(ctx context.Context, id string) (*domain.User, error)
	GetUserByEmail(ctx context.Context, email string) (*domain.User, error)

	ExistUserByEmail(ctx context.Context, email string) (bool, error)
}

type userRepository struct {
	db *sql.DB
}


func (r *userRepository) CreateUser(ctx context.Context, user *domain.User) error {

	query := "INSERT INTO users (username, email, password) VALUES ($1, $2, $3)"
	_, err := r.db.QueryContext(ctx, query, user.Username, user.Email, user.Password)
	if err != nil {
		return err
	}

	return nil
}


func (r *userRepository) GetUserById(ctx context.Context, id string) (*domain.User, error) {
	u := domain.User{}


	query := "SELECT id, username, email, password FROM users WHERE id = $1"
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&u.ID, 
		&u.Username, 
		&u.Email, 
		&u.Password,
	)
	if err != nil {
		return &domain.User{}, err
	}

	return &u, nil
}


func (r *userRepository) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	u := domain.User{}

	query := "SELECT id, username, email, password FROM users WHERE email = $1"
	err := r.db.QueryRowContext(ctx, query, email).Scan(
		&u.ID, 
		&u.Username, 
		&u.Email, 
		&u.Password,
	)
	if err != nil {
		return &domain.User{}, err
	}

	return &u, nil
}



func (r *userRepository) ExistUserByEmail(ctx context.Context, email string) (bool, error) {
	var id string

	query := "SELECT id FROM users WHERE email = $1"
	err := r.db.QueryRowContext(ctx, query, email).Scan(&id)
	if err != nil {
		return false, err
	}

	return false, nil
}
