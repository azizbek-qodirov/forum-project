package handlers

import (
	"context"
	"net/http"
	"strconv"

	pbr "api-gateway/genproto/reservation" // Import generated protobuf package

	"github.com/gin-gonic/gin"
)

// RestaurantCreate handles the creation of a new restaurant.
// @Summary Create restaurant
// @Description Create a new restaurant
// @Tags restaurant
// @Accept json
// @Produce json
// @Param restaurant body pbr.RestaurantReq true "Restaurant data"
// @Success 200 {object} pbr.Restaurant
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /restaurant [post]
func (h *HTTPHandler) RestaurantCreate(c *gin.Context) {
	var req pbr.RestaurantReq
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	res, err := h.Restaurant.Create(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// RestaurantGet handles getting a restaurant by ID.
// @Summary Get restaurant
// @Description Get a restaurant by ID
// @Tags restaurant
// @Accept json
// @Produce json
// @Param id path string true "Restaurant ID"
// @Success 200 {object} pbr.Restaurant
// @Failure 400 {object} string "Invalid restaurant ID"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /restaurant/{id} [get]
func (h *HTTPHandler) RestaurantGet(c *gin.Context) {
	id := &pbr.GetByIdReq{Id: c.Param("id")}
	res, err := h.Restaurant.Get(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't get restaurant", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// RestaurantUpdate handles updating an existing restaurant.
// @Summary Update restaurant
// @Description Update an existing restaurant
// @Tags restaurant
// @Accept json
// @Produce json
// @Param id path string true "Restaurant ID"
// @Param restaurant body pbr.RestaurantReq true "Updated restaurant data"
// @Success 200 {object} pbr.RestaurantReq
// @Failure 400 {object} string "Invalid request payload"
// @Failure 404 {object} string "Restaurant not found"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /restaurant/{id} [put]
func (h *HTTPHandler) RestaurantUpdate(c *gin.Context) {
	id := c.Param("id")
	req := &pbr.RestaurantUpdate{}

	if err := c.BindJSON(&req.UpdateRestaurant); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	req.Id = &pbr.GetByIdReq{Id: id}
	res, err := h.Restaurant.Update(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't update restaurant", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// RestaurantDelete handles deleting a restaurant by ID.
// @Summary Delete restaurant
// @Description Delete a restaurant by ID
// @Tags restaurant
// @Accept json
// @Produce json
// @Param id path string true "Restaurant ID"
// @Success 200 {object} string "Restaurant deleted"
// @Failure 400 {object} string "Invalid restaurant ID"
// @Failure 404 {object} string "Restaurant not found"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /restaurant/{id} [delete]
func (h *HTTPHandler) RestaurantDelete(c *gin.Context) {
	id := &pbr.GetByIdReq{Id: c.Param("id")}
	_, err := h.Restaurant.Delete(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't delete restaurant", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Restaurant deleted"})
}

// RestaurantGetAll handles getting all restaurants.
// @Summary Get all restaurants
// @Description Get all restaurants
// @Tags restaurant
// @Accept json
// @Produce json
// @Param limit query integer false "Limit"
// @Param offset query integer false "Offset"
// @Success 200 {object} pbr.GetAllRestaurantRes
// @Failure 400 {object} string "Invalid parameters"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /restaurants [get]
func (h *HTTPHandler) RestaurantGetAll(c *gin.Context) {
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

	res, err := h.Restaurant.GetAll(context.Background(), &pbr.GetAllRestaurantReq{
		Filter: &pbr.Filter{
			Limit:  int32(limit),
			Offset: int32(offset),
		},
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't get restaurants", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}
