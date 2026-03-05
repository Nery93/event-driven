package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/guilh/event-system/inventory-service/internal/domain"
	"github.com/guilh/event-system/inventory-service/internal/usecase"
)

type HTTPHandler struct {
	inventoryUseCase *usecase.InventoryUseCase
}

func NewHTTPHandler(inventoryUseCase *usecase.InventoryUseCase) *HTTPHandler {
	return &HTTPHandler{
		inventoryUseCase: inventoryUseCase,
	}
}

func (h *HTTPHandler) CreateInventory(c *gin.Context) {

	var inventory domain.Inventory
	if err := c.ShouldBindJSON(&inventory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	created, err := h.inventoryUseCase.CreateInventory(c.Request.Context(), &inventory)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, created)
}

func (h *HTTPHandler) GetProduct(c *gin.Context) {

	id := c.Param("id")

	inventory, err := h.inventoryUseCase.GetProductByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if inventory == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
		return
	}

	c.JSON(http.StatusOK, inventory)
}
