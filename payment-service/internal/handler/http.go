package handler


import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/guilh/event-system/payment-service/internal/domain"
	"github.com/guilh/event-system/payment-service/internal/usecase"
	
)

type HTTPHandler struct {
	paymentUseCase *usecase.PaymentUseCase
}

func NewHTTPHandler(paymentUseCase *usecase.PaymentUseCase) *HTTPHandler {
	return &HTTPHandler{
		paymentUseCase: paymentUseCase,
	}
}

func (h *HTTPHandler) ProcessPayment(c *gin.Context) {

	var payment domain.Payment
	if err := c.ShouldBindJSON(&payment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	}

	processedPayment, err := h.paymentUseCase.ProcessPayment(c.Request.Context(), &payment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, processedPayment)

}

func (h *HTTPHandler) GetPayment(c *gin.Context) {

	id := c.Param("id")

	payment, err := h.paymentUseCase.GetPaymentByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Payment not found"})
		return
	}

	c.JSON(http.StatusOK, payment)
}