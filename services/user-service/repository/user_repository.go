package repository

import (
	"context"
	"database/sql"
	"ticket-booking-platform/services/user-service/model"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *model.User) error
	FindUserByEmail(ctx context.Context, email string) (*model.User, error)
	FindUserByID(ctx context.Context, id int64) (*model.User, error)
	UpdateUser(ctx context.Context, user *model.User) error
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(ctx context.Context, user *model.User) error {
	query := `
		INSERT INTO users (email, password_hash, full_name, phone)
		VALUES ($1, $2, $3, $4)
		RETURNING id, created_at, updated_at`
	
	return r.db.QueryRowContext(ctx, query,
		user.Email, user.PasswordHash, user.FullName, user.Phone,
	).Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt)
}

func (r *userRepository) FindUserByEmail(ctx context.Context, email string) (*model.User, error) {
	user := &model.User{}
	query := `
		SELECT id, email, password_hash, full_name, phone, created_at, updated_at
		FROM users WHERE email = $1`
	
	err := r.db.QueryRowContext(ctx, query, email).Scan(
		&user.ID, &user.Email, &user.PasswordHash, &user.FullName,
		&user.Phone, &user.CreatedAt, &user.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return user, err
}

func (r *userRepository) FindUserByID(ctx context.Context, id int64) (*model.User, error) {
	user := &model.User{}
	query := `
		SELECT id, email, password_hash, full_name, phone, created_at, updated_at
		FROM users WHERE id = $1`
	
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&user.ID, &user.Email, &user.PasswordHash, &user.FullName,
		&user.Phone, &user.CreatedAt, &user.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return user, err
}

func (r *userRepository) UpdateUser(ctx context.Context, user *model.User) error {
	query := `
		UPDATE users 
		SET full_name = $1, phone = $2, updated_at = CURRENT_TIMESTAMP
		WHERE id = $3
		RETURNING updated_at`
	
	return r.db.QueryRowContext(ctx, query,
		user.FullName, user.Phone, user.ID,
	).Scan(&user.UpdatedAt)
}
