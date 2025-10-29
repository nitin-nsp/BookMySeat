package model

import "time"

type User struct {
	ID           int64     `db:"id" json:"id"`
	Email        string    `db:"email" json:"email"`
	PasswordHash string    `db:"password_hash" json:"-"`
	FullName     string    `db:"full_name" json:"full_name"`
	Phone        string    `db:"phone" json:"phone"`
	CreatedAt    time.Time `db:"created_at" json:"created_at"`
	UpdatedAt    time.Time `db:"updated_at" json:"updated_at"`
}

type UserProfile struct {
	ID          int64     `db:"id" json:"id"`
	UserID      int64     `db:"user_id" json:"user_id"`
	DateOfBirth *time.Time `db:"date_of_birth" json:"date_of_birth,omitempty"`
	Address     string    `db:"address" json:"address"`
	City        string    `db:"city" json:"city"`
	State       string    `db:"state" json:"state"`
	Country     string    `db:"country" json:"country"`
	PostalCode  string    `db:"postal_code" json:"postal_code"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
}
