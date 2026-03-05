package handler

// HTTPHandler holds Gin HTTP handlers.
// Depends only on usecase — never imports repository or kafka directly.

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/guilh/event-system/order-service/internal/domain"
	"github.com/guilh/event-system/order-service/internal/usecase"
)

type HTTPHandler struct {
	orderUseCase *usecase.OrderUseCase
}

func NewHTTPHandler(orderUseCase *usecase.OrderUseCase) *HTTPHandler {
	return &HTTPHandler{
		orderUseCase: orderUseCase,
	}
}

func (h *HTTPHandler) CreateOrder(c *gin.Context) {

	var order domain.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	}

	createdOrder, err := h.orderUseCase.CreateOrder(c.Request.Context(), &order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdOrder)
}

func (h *HTTPHandler) GetOrder(c *gin.Context) {

	id := c.Param("id")

	order, err := h.orderUseCase.GetOrderByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	c.JSON(http.StatusOK, order)
}
