package service

import (
	"github.com/stretchr/testify/assert"
	"github.com/vikasd145/nfi-test-project/pkg/repository"
	"testing"
)

func TestDeposit(t *testing.T) {
	// Create a mock UserRepository
	mockUserRepo := &mockUserRepository{}

	// Create a TransactionService with the mockUserRepo
	service := &transactionService{
		userRepository: mockUserRepo,
	}

	// Perform the deposit
	totalAmount, err := service.Deposit(1, 100.0)

	// Verify the result
	assert.NoError(t, err)
	assert.Equal(t, 100.0, totalAmount)

	// Verify that the UpdateBalance function was called with the correct arguments
	assert.True(t, mockUserRepo.UpdateBalanceCalled)
	assert.Equal(t, 1, mockUserRepo.UpdateBalanceUserID)
	assert.Equal(t, 100.0, mockUserRepo.UpdateBalanceAmount)
}

func TestWithdraw(t *testing.T) {
	// Create a mock UserRepository
	mockUserRepo := &mockUserRepository{
		UpdateBalanceAmount: 100.0,
	}

	// Create a TransactionService with the mockUserRepo
	service := &transactionService{
		userRepository: mockUserRepo,
	}

	// Perform the withdrawal
	totalAmount, err := service.Withdraw(1, 50.0)

	// Verify the result
	assert.NoError(t, err)
	assert.Equal(t, 50.0, totalAmount)

	// Verify that the UpdateBalance function was called with the correct arguments
	assert.True(t, mockUserRepo.UpdateBalanceCalled)
	assert.Equal(t, 1, mockUserRepo.UpdateBalanceUserID)
	assert.Equal(t, 50.0, mockUserRepo.UpdateBalanceAmount)
}

type mockUserRepository struct {
	UpdateBalanceCalled bool
	UpdateBalanceUserID int
	UpdateBalanceAmount float64
	RegisterUserCalled  bool
}

func (m *mockUserRepository) GetUserByID(userID int) (*repository.User, error) {
	// Mock implementation for GetUserByID
	return &repository.User{ID: userID, Balance: 0.0}, nil
}

func (m *mockUserRepository) UpdateUserBalance(userID int, amount float64) (float64, error) {
	// Mock implementation for UpdateBalance
	m.UpdateBalanceCalled = true
	m.UpdateBalanceUserID = userID
	m.UpdateBalanceAmount = m.UpdateBalanceAmount + amount
	return m.UpdateBalanceAmount, nil
}

func (m *mockUserRepository) CreateUser() (int, error) {
	m.RegisterUserCalled = true
	return 1, nil
}
