package http_handlers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestRegisterHandler(t *testing.T) {
	// Create a new Gin router
	router := gin.Default()

	// Create mock UserService
	mockUserService := &mockUserService{}

	// Create a handler with the mockUserService
	h := &handler{
		userService: mockUserService,
	}

	// Register the registerHandler route
	router.POST("/register", h.registerHandler)

	// Create a test request
	reqBody := strings.NewReader("")
	req, _ := http.NewRequest("POST", "/register", reqBody)
	res := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(res, req)

	// Verify the response
	assert.Equal(t, http.StatusOK, res.Code)
	assert.Equal(t, "application/json; charset=utf-8", res.Header().Get("Content-Type"))

	var response struct {
		UserID  int     `json:"user_id"`
		Balance float64 `json:"balance"`
	}

	err := json.Unmarshal(res.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, 1, response.UserID)
	assert.Equal(t, 0.0, response.Balance)

	// Verify that the RegisterUser function was called
	assert.True(t, mockUserService.registerUserCalled)
}

func TestDepositHandler(t *testing.T) {
	// Create a new Gin router
	router := gin.Default()

	// Create mock UserService
	mockUserService := &mockUserService{}

	// Create a handler with the mockUserService
	h := &handler{
		userService: mockUserService,
	}

	// Register the depositHandler route
	router.POST("/deposit/:userID", h.depositHandler)

	// Create a test request
	userID := "1"
	amount := "100.0"
	reqBody := strings.NewReader(`{"amount": ` + amount + `}`)
	req, _ := http.NewRequest("POST", "/deposit/"+userID, reqBody)
	res := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(res, req)

	// Verify the response
	assert.Equal(t, http.StatusOK, res.Code)
	assert.Equal(t, "application/json; charset=utf-8", res.Header().Get("Content-Type"))

	var response struct {
		TotalAmount float64 `json:"total_amount"`
	}

	err := json.Unmarshal(res.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, 100.0, response.TotalAmount)

	// Verify that the Deposit function was called with the correct arguments
	assert.True(t, mockUserService.depositCalled)
	assert.Equal(t, 1, mockUserService.depositUserID)
	assert.Equal(t, 100.0, mockUserService.depositAmount)
}

func TestWithdrawalHandler(t *testing.T) {
	// Create a new Gin router
	router := gin.Default()

	// Create mock UserService
	mockUserService := &mockUserService{}

	// Create a handler with the mockUserService
	h := &handler{
		userService: mockUserService,
	}

	// Register the withdrawalHandler route
	router.POST("/withdrawal/:userID", h.withdrawHandler)

	// Create a test request
	userID := "1"
	amount := "50.0"
	reqBody := strings.NewReader(`{"amount": ` + amount + `}`)
	req, _ := http.NewRequest("POST", "/withdrawal/"+userID, reqBody)
	res := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(res, req)

	// Verify the response
	assert.Equal(t, http.StatusOK, res.Code)
	assert.Equal(t, "application/json; charset=utf-8", res.Header().Get("Content-Type"))

	var response struct {
		TotalAmount float64 `json:"total_amount"`
	}

	err := json.Unmarshal(res.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, 50.0, response.TotalAmount)

	// Verify that the Withdrawal function was called with the correct arguments
	assert.True(t, mockUserService.withdrawalCalled)
	assert.Equal(t, 1, mockUserService.withdrawalUserID)
	assert.Equal(t, 50.0, mockUserService.withdrawalAmount)
}

type mockUserService struct {
	registerUserCalled bool
	depositCalled      bool
	withdrawalCalled   bool
	depositUserID      int
	withdrawalUserID   int
	depositAmount      float64
	withdrawalAmount   float64
}

func (m *mockUserService) RegisterUser() (int, error) {
	m.registerUserCalled = true
	return 1, nil
}

func (m *mockUserService) Deposit(userID int, amount float64) (float64, error) {
	m.depositCalled = true
	m.depositUserID = userID
	m.depositAmount = amount
	return 100.0, nil
}

func (m *mockUserService) Withdrawal(userID int, amount float64) (float64, error) {
	m.withdrawalCalled = true
	m.withdrawalUserID = userID
	m.withdrawalAmount = amount
	return 50.0, nil
}
