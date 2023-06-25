package http_handlers

import (
	"github.com/gin-gonic/gin"
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
	assert.Equal(t, `{"user_id":1,"balance":0}`, res.Body.String())

	// Verify that the RegisterUser function was called
	assert.True(t, mockUserService.registerUserCalled)
}

type mockUserService struct {
	registerUserCalled bool
}

func (m *mockUserService) RegisterUser() (int, error) {
	m.registerUserCalled = true
	return 1, nil
}
