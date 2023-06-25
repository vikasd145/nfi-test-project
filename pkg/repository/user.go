package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"math"
)

type UserRepository interface {
	CreateUser() (int, error)
	GetUserByID(userID int) (*User, error)
	UpdateUserBalance(userID int, newBalance float64) (float64, error)
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

// GetUserByID In real world, we should check if the user exists or not.If we acquire lock on row without
// checking row exist or not it will wait until locck is acquired by inserting row through other
// transaction or transaction timeout
func (r *userRepository) GetUserByID(userID int) (*User, error) {
	user := &User{}
	err := r.db.Get(user, "SELECT * FROM users WHERE id = $1", userID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) UpdateUserBalance(userID int, amount float64) (float64, error) {
	tx, err := r.db.Beginx()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()
	// To acquire lock for this row
	user := &User{}
	err = tx.Get(user, "SELECT * FROM users WHERE id = $1 FOR UPDATE", userID)
	if err != nil {
		return 0, err
	}
	if amount < 0 && user.Balance < math.Abs(amount) {
		return user.Balance, fmt.Errorf("insufficient balance")
	}
	_, err = r.db.Exec("UPDATE users SET balance = $1 WHERE id = $2", user.Balance+amount, userID)
	if err != nil {
		tx.Commit()
		return user.Balance, err
	}
	return user.Balance, err
}
