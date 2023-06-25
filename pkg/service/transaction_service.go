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
	// In real world, we should check if the user exists or not.
	_, err := s.userRepository.GetUserByID(userID)
	if err != nil {
		return 0, err
	}
	newBalance, err := s.userRepository.UpdateUserBalance(userID, amount)
	if err != nil {
		return 0, err
	}

	return newBalance, nil
}

func (s *transactionService) Withdraw(userID int, amount float64) (float64, error) {
	_, err := s.userRepository.GetUserByID(userID)
	if err != nil {
		return 0, err
	}
	newBalance, err := s.userRepository.UpdateUserBalance(userID, -amount)
	if err != nil {
		return 0, err
	}

	return newBalance, nil
}
