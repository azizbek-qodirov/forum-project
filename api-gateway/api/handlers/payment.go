package handlers

import (
	pbp "api-gateway/genproto/payment"

	"github.com/gin-gonic/gin"
)

// GetPayment godoc
// @Summary Get payment
// @Description Get a payment by ID
// @Tags payment
// @Accept json
// @Produce json
// @Param id path string true "Payment ID"
// @Success 200 {object} pbp.PaymentGetByIdResp
// @Failure 401 {object} string "Unauthorized"
// @Failure 404 {object} string "Payment not found"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /payment/{id} [get]
func (h *HTTPHandler) GetPaymentByID(c *gin.Context) {
	id := c.Param("id")
	payment, err := h.Payment.GetById(c, &pbp.PaymentGetByIdReq{Id: id})
	if err != nil {
		c.JSON(400, gin.H{"Couldn't get the payment": err.Error()})
		return
	}
	c.JSON(200, payment)
}

// UpdatePayment godoc
// @Summary Update payment
// @Description Update a payment by id
// @Tags payment
// @Accept json
// @Produce json
// @Param id path string true "Payment ID"
// @Param payment body pbp.PaymentCreateReq true "Updated payment data"
// @Success 200 {object} pbp.PaymentGetByIdResp
// @Failure 401 {object} string "Unauthorized"
// @Failure 404 {object} string "Payment not found"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /payment/{id} [put]
func (h *HTTPHandler) UpdatePayment(c *gin.Context) {
	id := c.Param("id")
	var paym pbp.PaymentCreateReq
	if err := c.ShouldBindJSON(&paym); err != nil {
		c.JSON(400, gin.H{"Couldn't bind the payment": err.Error()})
		return
	}
	payment := &pbp.PaymentUpdateReq{}
	payment.Id = id
	payment.Payment = &paym
	res, err := h.Payment.Update(c, payment)
	if err != nil {
		c.JSON(400, gin.H{"Couldn't update the payment": err.Error()})
		return
	}
	c.JSON(200, res)
}
