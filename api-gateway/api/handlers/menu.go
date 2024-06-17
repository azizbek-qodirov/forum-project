// menu.go
package handlers

import (
	"context"
	"net/http"
	"strconv"

	pbr "api-gateway/genproto/reservation" // Import generated protobuf package

	"github.com/gin-gonic/gin"
)

// MenuCreate handles the creation of a new menu item.
// @Summary Create menu item
// @Description Create a new menu item
// @Tags menu
// @Accept json
// @Produce json
// @Param menu body pbr.MenuReq true "Menu data"
// @Success 200 {object} pbr.Menu
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /menu [post]
func (h *HTTPHandler) MenuCreate(c *gin.Context) {
	var req pbr.MenuReq
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	res, err := h.Menu.Create(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// MenuGet handles getting a menu item by ID.
// @Summary Get menu item
// @Description Get a menu item by ID
// @Tags menu
// @Accept json
// @Produce json
// @Param id path string true "Menu ID"
// @Success 200 {object} pbr.MenuRes
// @Failure 400 {object} string "Invalid menu item ID"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /menu/{id} [get]
func (h *HTTPHandler) MenuGet(c *gin.Context) {
	id := &pbr.GetByIdReq{Id: c.Param("id")}
	res, err := h.Menu.Get(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't get menu item", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// MenuUpdate handles updating an existing menu item.
// @Summary Update menu item
// @Description Update an existing menu item
// @Tags menu
// @Accept json
// @Produce json
// @Param id path string true "Menu ID"
// @Param menu body pbr.MenuReq true "Updated menu item data"
// @Success 200 {object} pbr.Menu
// @Failure 400 {object} string "Invalid request payload"
// @Failure 404 {object} string "Menu item not found"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /menu/{id} [put]
func (h *HTTPHandler) MenuUpdate(c *gin.Context) {
	id := c.Param("id")
	var req pbr.MenuUpdate

	if err := c.BindJSON(&req.UpdateMenu); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	req.Id = &pbr.GetByIdReq{Id: id}
	res, err := h.Menu.Update(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't update menu item", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// MenuDelete handles deleting a menu item by ID.
// @Summary Delete menu item
// @Description Delete a menu item by ID
// @Tags menu
// @Accept json
// @Produce json
// @Param id path string true "Menu ID"
// @Success 200 {object} string "Menu item deleted"
// @Failure 400 {object} string "Invalid menu item ID"
// @Failure 404 {object} string "Menu item not found"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /menu/{id} [delete]
func (h *HTTPHandler) MenuDelete(c *gin.Context) {
	id := &pbr.GetByIdReq{Id: c.Param("id")}
	_, err := h.Menu.Delete(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't delete menu item", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Menu item deleted"})
}

// MenuGetAll handles getting all menu items.
// @Summary Get all menu items
// @Description Get all menu items
// @Tags menu
// @Accept json
// @Produce json
// @Param limit query integer false "Limit"
// @Param offset query integer false "Offset"
// @Success 200 {object} pbr.GetAllMenuRes
// @Failure 400 {object} string "Invalid parameters"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /menus [get]
func (h *HTTPHandler) MenuGetAll(c *gin.Context) {
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

	res, err := h.Menu.GetAll(context.Background(), &pbr.GetAllMenuReq{
		Filter: &pbr.Filter{
			Limit:  int32(limit),
			Offset: int32(offset),
		},
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't get menu items", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}
