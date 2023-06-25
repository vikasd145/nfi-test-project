package service

import (
	"github.com/vikasd145/nfi-test-project/pkg/repository"
)

type TransactionService interface {
	Deposit(userID int, amount float64) (float64, error)
	Withdraw(userID int, amount float64) (float64, error)
}

type transactionService struct {
	userRepository repository.UserRepository
}

func NewTransactionService(userRepository repository.UserRepository) TransactionService {
	return &transactionService{
		userRepository: userRepository,
	}
}

func (s *transactionService) Deposit(userID int, amount float64) (float64, error) {
	user, err := s.userRepository.GetUserByID(userID)
	if err != nil {
		return 0, err
	}

	newBalance := user.Balance + amount
	err = s.userRepository.UpdateUserBalance(userID, newBalance)
	if err != nil {
		return 0, err
	}

	return newBalance, nil
}

func (s *transactionService) Withdraw(userID int, amount float64) (float64, error) {
	user, err := s.userRepository.GetUserByID(userID)
	if err != nil {
		return 0, err
	}

	if amount > user.Balance {
		return user.Balance, nil
	}

	newBalance := user.Balance - amount
	err = s.userRepository.UpdateUserBalance(userID, newBalance)
	if err != nil {
		return 0, err
	}

	return newBalance, nil
}
