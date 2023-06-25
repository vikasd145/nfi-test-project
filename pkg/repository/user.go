package repository

import (
	"github.com/jmoiron/sqlx"
)

type UserRepository interface {
	CreateUser() (int, error)
	GetUserByID(userID int) (*User, error)
	UpdateUserBalance(userID int, newBalance float64) error
}

type userRepository struct {
	db *sqlx.DB
}

type User struct {
	ID      int
	Balance float64
}

func NewUserRepository(db *sqlx.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) CreateUser() (int, error) {
	var userID int
	err := r.db.QueryRow("INSERT INTO users (balance) VALUES (0) RETURNING id").Scan(&userID)
	if err != nil {
		return 0, err
	}
	return userID, nil
}

func (r *userRepository) GetUserByID(userID int) (*User, error) {
	user := &User{}
	err := r.db.Get(user, "SELECT * FROM users WHERE id = $1", userID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) UpdateUserBalance(userID int, newBalance float64) error {
	_, err := r.db.Exec("UPDATE users SET balance = $1 WHERE id = $2", newBalance, userID)
	return err
}
