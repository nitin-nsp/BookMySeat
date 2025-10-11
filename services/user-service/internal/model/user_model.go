package model

import (
	"database/sql"
	"fmt"
	"time"
)

type User struct {
	Id        int64  `db:"id"`
	Email     string `db:"email"`
	Password  string `db:"password"`
	Phone     string `db:"phone"`
	Name      string `db:"name"`
	Role      string `db:"role"`
	CreatedAt int64  `db:"created_at"`
	UpdatedAt int64  `db:"updated_at"`
}

type UserModel struct {
	db *sql.DB
}

func NewUserModel(db *sql.DB) *UserModel {
	return &UserModel{db: db}
}

func (m *UserModel) Insert(user *User) (int64, error) {
	user.CreatedAt = time.Now().Unix()
	user.UpdatedAt = time.Now().Unix()
	
	query := `INSERT INTO users (email, password, phone, name, role, created_at, updated_at) 
			  VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`
	
	err := m.db.QueryRow(query, user.Email, user.Password, user.Phone, user.Name, user.Role, 
		user.CreatedAt, user.UpdatedAt).Scan(&user.Id)
	
	return user.Id, err
}

func (m *UserModel) FindByEmail(email string) (*User, error) {
	var user User
	query := `SELECT id, email, password, phone, name, role, created_at, updated_at 
			  FROM users WHERE email = $1`
	
	err := m.db.QueryRow(query, email).Scan(&user.Id, &user.Email, &user.Password, 
		&user.Phone, &user.Name, &user.Role, &user.CreatedAt, &user.UpdatedAt)
	
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("user not found")
	}
	return &user, err
}

func (m *UserModel) FindById(id int64) (*User, error) {
	var user User
	query := `SELECT id, email, password, phone, name, role, created_at, updated_at 
			  FROM users WHERE id = $1`
	
	err := m.db.QueryRow(query, id).Scan(&user.Id, &user.Email, &user.Password, 
		&user.Phone, &user.Name, &user.Role, &user.CreatedAt, &user.UpdatedAt)
	
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("user not found")
	}
	return &user, err
}

func (m *UserModel) Update(user *User) error {
	user.UpdatedAt = time.Now().Unix()
	query := `UPDATE users SET name = $1, phone = $2, updated_at = $3 WHERE id = $4`
	_, err := m.db.Exec(query, user.Name, user.Phone, user.UpdatedAt, user.Id)
	return err
}
