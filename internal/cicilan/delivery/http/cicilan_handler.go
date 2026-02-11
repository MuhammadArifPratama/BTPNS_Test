package http

import (
	"net/http"

	"btpntest/domain"
	cicilan "btpntest/internal/cicilan"
	"btpntest/internal/cicilan/usecase"

	"github.com/gin-gonic/gin"
)

type CicilanHandler struct {
	usecase cicilan.CicilanUsecase
}

func NewCicilanHandler(usecaseImpl cicilan.CicilanUsecase) *CicilanHandler {
	return &CicilanHandler{usecase: usecaseImpl}
}

// CalculateInstallments godoc
// @Summary Calculate installment schedule
// @Description Returns installment calculations for available tenors.
// @Tags Installments
// @Accept json
// @Produce json
// @Param request body domain.CalculateInstallmentRequest true "Installment request"
// @Success 200 {object} domain.CalculateInstallmentResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /calculate-installments [post]
// @Router /btpn/calculate-installments [post]
func (h *CicilanHandler) CalculateInstallments(c *gin.Context) {
	var req domain.CalculateInstallmentRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	response, err := h.usecase.CalculateInstallments(&req)
	if err != nil {
		if _, ok := err.(*usecase.ValidationError); ok {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, response)
}

func (h *CicilanHandler) RegisterRoutes(router *gin.Engine) {
	router.POST("/calculate-installments", h.CalculateInstallments)
}
