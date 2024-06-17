package handlers

import (
	pbr "api-gateway/genproto/reservation"
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// CreateReservation godoc
// @Summary Create reservation
// @Description Create a new reservation
// @Tags reservation
// @Accept json
// @Produce json
// @Param reservation body pbr.ReservationReq true "Reservation data"
// @Success 200 {object} pbr.Reservation
// @Failure 401 {object} string "Unauthorized"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /reservation [post]
func (h *HTTPHandler) ReservationCreate(c *gin.Context) {
	var req pbr.ReservationReq
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	res, err := h.Reservation.Create(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// GetReservation godoc
// @Summary Get reservation
// @Description Get a reservation by ID
// @Tags reservation
// @Accept json
// @Produce json
// @Param id path string true "Reservation ID"
// @Success 200 {object} pbr.ReservationRes
// @Failure 401 {object} string "Unauthorized"
// @Failure 404 {object} string "Reservation not found"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /reservation/{id} [get]
func (h *HTTPHandler) ReservationGet(c *gin.Context) {
	id := &pbr.GetByIdReq{Id: c.Param("id")}
	res, err := h.Reservation.Get(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't get reservation", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// UpdateReservation godoc
// @Summary Update reservation
// @Description Update an existing reservation
// @Tags reservation
// @Accept json
// @Produce json
// @Param id path string true "Reservation ID"
// @Param reservation body pbr.ReservationReq true "Updated reservation data"
// @Success 200 {object} pbr.ReservationReq
// @Failure 401 {object} string "Unauthorized"
// @Failure 404 {object} string "Reservation not found"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /reservation/{id} [put]
func (h *HTTPHandler) ReservationUpdate(c *gin.Context) {
	id := c.Param("id")
	var req pbr.ReservationUpdate
	if err := c.BindJSON(&req.UpdateReservation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	req.Id = &pbr.GetByIdReq{Id: id}
	res, err := h.Reservation.Update(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't update reservation", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// DeleteReservation godoc
// @Summary Delete reservation
// @Description Delete a reservation by ID
// @Tags reservation
// @Accept json
// @Produce json
// @Param id path string true "Reservation ID"
// @Success 200 {object} string "Reservation deleted"
// @Failure 401 {object} string "Unauthorized"
// @Failure 404 {object} string "Reservation not found"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /reservation/{id} [delete]
func (h *HTTPHandler) ReservationDelete(c *gin.Context) {
	id := &pbr.GetByIdReq{Id: c.Param("id")}
	_, err := h.Reservation.Delete(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't delete reservation", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Reservation deleted"})
}

// GetAllReservations godoc
// @Summary Get all reservations
// @Description Get all reservations
// @Tags reservation
// @Accept json
// @Produce json
// @Param user_id query string false "User ID"
// @Param limit query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {object} pbr.GetAllReservationRes
// @Failure 400 {object} string "Invalid parameters"
// @Failure 401 {object} string "Unauthorized"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /reservations [get]
func (h *HTTPHandler) ReservationGetAll(c *gin.Context) {
	user_id := c.Query("user_id")
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

	res, err := h.Reservation.GetAll(context.Background(), &pbr.GetAllReservationReq{
		UserId: user_id,
		Filter: &pbr.Filter{
			Limit:  int32(limit),
			Offset: int32(offset),
		},
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't get reservations", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h *HTTPHandler) ReservationCheck(c *gin.Context) {
	type checkModel struct {
		RestaurantID string `json:"restaurant_id"`
		DateTime     string `json:"date_time"`
	}
	var check checkModel
	if err := c.BindJSON(&check); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload", "details": err.Error()})
		return
	}

	dateTimeLayout := "2006-01-02 15:04:05"

	parsedDateTime, err := time.Parse(dateTimeLayout, check.DateTime)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date time format"})
		return
	}
	resp, err := h.Reservation.CheckTime(context.Background(), &pbr.CheckTimeReq{Time: parsedDateTime.String(), RestauranId: check.RestaurantID})
	if err != nil {
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date time format"})
			return
		}
	}
	if resp.IsBooked {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Restaurant is booked"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Restaurant is available"})
		return
	}
}
