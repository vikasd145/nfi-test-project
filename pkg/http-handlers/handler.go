package http_handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/vikasd145/nfi-test-project/pkg/service"
	"net/http"
	"strconv"
)

type handler struct {
	userService        service.UserService
	transactionService service.TransactionService
}

func RegisterHandlers(router *gin.Engine, userService service.UserService, transactionService service.TransactionService) {
	h := &handler{
		userService:        userService,
		transactionService: transactionService,
	}

	router.POST("/register", h.registerHandler)
	router.POST("/deposit", h.depositHandler)
	router.POST("/withdraw", h.withdrawHandler)
}

func (h *handler) registerHandler(c *gin.Context) {
	userID, err := h.userService.RegisterUser()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user_id": userID, "balance": 0})
}

func (h *handler) depositHandler(c *gin.Context) {
	userIDStr := c.PostForm("user_id")
	amountStr := c.PostForm("amount")

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid amount"})
		return
	}

	newBalance, err := h.transactionService.Deposit(userID, amount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to deposit amount"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"balance": newBalance})
}

func (h *handler) withdrawHandler(c *gin.Context) {
	userIDStr := c.PostForm("user_id")
	amountStr := c.PostForm("amount")

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid amount"})
		return
	}

	newBalance, err := h.transactionService.Withdraw(userID, amount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to withdraw amount"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"balance": newBalance})
}
