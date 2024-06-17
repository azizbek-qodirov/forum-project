// reservation_order.go
package handlers

import (
	"context"
	"net/http"
	"strconv"

	pbr "api-gateway/genproto/reservation" // Import generated protobuf package

	"github.com/gin-gonic/gin"
)

// ReservationOrderCreate handles the creation of a new reservation order.
// @Summary Create reservation order
// @Description Create a new reservation order
// @Tags reservation_order
// @Accept json
// @Produce json
// @Param reservation_order body pbr.ReservationOrderUpdate true "Reservation order data"
// @Success 200 {object} pbr.ReservationOrderRes
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /reservation_order [post]
func (h *HTTPHandler) ReservationOrderCreate(c *gin.Context) {
	var req pbr.ReservationOrderUpdateReq
	if err := c.BindJSON(&req.Update); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	var reqq pbr.ReservationOrderReq

	reqq.ReservationId = req.Update.ReservationId
	reqq.MenuItemId = req.Update.MenuItemId
	reqq.Quantity = req.Update.Quantity

	res, err := h.ReservationOrder.Create(context.Background(), &reqq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// ReservationOrderGet handles getting a reservation order by ID.
// @Summary Get reservation order
// @Description Get a reservation order by ID
// @Tags reservation_order
// @Accept json
// @Produce json
// @Param id path string true "Reservation order ID"
// @Success 200 {object} pbr.ReservationOrderRes
// @Failure 400 {object} string "Invalid reservation order ID"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /reservation_order/{id} [get]
func (h *HTTPHandler) ReservationOrderGet(c *gin.Context) {
	id := &pbr.GetByIdReq{Id: c.Param("id")}
	res, err := h.ReservationOrder.Get(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't get reservation order", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// ReservationOrderUpdate handles updating an existing reservation order.
// @Summary Update reservation order
// @Description Update an existing reservation order
// @Tags reservation_order
// @Accept json
// @Produce json
// @Param id path string true "Reservation order ID"
// @Param reservation_order body pbr.ReservationOrderUpdate true "Updated reservation order data"
// @Success 200 {object} pbr.ReservationOrderUpdate
// @Failure 400 {object} string "Invalid request payload"
// @Failure 404 {object} string "Reservation order not found"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /reservation_order/{id} [put]
func (h *HTTPHandler) ReservationOrderUpdate(c *gin.Context) {
	id := c.Param("id")
	var req pbr.ReservationOrderUpdate
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	var updateReq pbr.ReservationOrderUpdateReq
	updateReq.Id = &pbr.GetByIdReq{Id: id}
	updateReq.Update = &req
	res, err := h.ReservationOrder.Update(context.Background(), &updateReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't update reservation order", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// ReservationOrderDelete handles deleting a reservation order by ID.
// @Summary Delete reservation order
// @Description Delete a reservation order by ID
// @Tags reservation_order
// @Accept json
// @Produce json
// @Param id path string true "Reservation order ID"
// @Success 200 {object} string "Reservation order deleted"
// @Failure 400 {object} string "Invalid reservation order ID"
// @Failure 404 {object} string "Reservation order not found"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /reservation_order/{id} [delete]
func (h *HTTPHandler) ReservationOrderDelete(c *gin.Context) {
	id := &pbr.GetByIdReq{Id: c.Param("id")}
	_, err := h.ReservationOrder.Delete(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't delete reservation order", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Reservation order deleted"})
}

// ReservationOrderGetAll handles getting all reservation orders.
// @Summary Get all reservation orders
// @Description Get all reservation orders
// @Tags reservation_order
// @Accept json
// @Produce json
// @Param limit query integer false "Limit"
// @Param offset query integer false "Offset"
// @Success 200 {object} pbr.GetAllReservationOrderRes
// @Failure 400 {object} string "Invalid parameters"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /reservation_orders [get]
func (h *HTTPHandler) ReservationOrderGetAll(c *gin.Context) {
	limitStr := c.Query("limit")
	offsetStr := c.Query("offset")
	var limit, offset int
	var err error
	if limitStr == "" {
		limit = 10
	} else {
		limit, err = strconv.Atoi(limitStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit parameter"})
			return
		}
	}
	if offsetStr == "" {
		offset = 0
	} else {
		offset, err = strconv.Atoi(offsetStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid offset parameter"})
			return
		}
	}

	res, err := h.ReservationOrder.GetAll(context.Background(), &pbr.GetAllReservationOrderReq{
		Filter: &pbr.Filter{
			Limit:  int32(limit),
			Offset: int32(offset),
		},
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't get reservation orders", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}
