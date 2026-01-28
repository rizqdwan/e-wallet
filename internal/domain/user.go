package domain

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id" db:"id"`
	Username  string 		`json:"username" db:"username"`
	Email     string 		`json:"email" db:"email"`
	Password  string 		`db:"password"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdateAt 	time.Time `json:"update_at" db:"update_at"`

}