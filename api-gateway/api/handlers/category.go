package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	pb "api-gateway/forum-protos/genprotos"

	"github.com/gin-gonic/gin"
)

// CategoryCreate handles the creation of a new category.
// @Summary Create category
// @Description Create a new category
// @Tags category
// @Accept json
// @Produce json
// @Param category body pb.CategoryCReqForSwagger true "Category data"
// @Success 200 {object} pb.CategoryCReqOrCResOrGResOrUReqOrURes
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /category [post]
func (h *HTTPHandler) CategoryCreate(c *gin.Context) {
	fmt.Println("Category create trigger")
	var req pb.CategoryCReqOrCResOrGResOrUReqOrURes
	if err := c.BindJSON(&req); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON", "details": err.Error()})
		return
	}
	res, err := h.Category.Create(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// CategoryGet handles getting a category by ID.
// @Summary Get category
// @Description Get a category by ID
// @Tags category
// @Accept json
// @Produce json
// @Param id path string true "Category ID"
// @Success 200 {object} pb.CategoryCReqOrCResOrGResOrUReqOrURes
// @Failure 400 {object} string "Invalid category  ID"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /category/{id} [get]
func (h *HTTPHandler) CategoryGet(c *gin.Context) {
	id := &pb.CategoryGReqOrDReq{CategoryId: c.Param("id")}
	res, err := h.Category.GetByID(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't get category", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// CategoryUpdate handles updating an existing category .
// @Summary Update category
// @Description Update an existing category
// @Tags category
// @Accept json
// @Produce json
// @Param id path string true "Category ID"
// @Param category body pb.CategoryCReqForSwagger true "Updated category data"
// @Success 200 {object} pb.CategoryCReqOrCResOrGResOrUReqOrURes
// @Failure 400 {object} string "Invalid request payload"
// @Failure 404 {object} string "Category not found"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /category/{id} [put]
func (h *HTTPHandler) CategoryUpdate(c *gin.Context) {
	id := c.Param("id")
	var req pb.CategoryCReqOrCResOrGResOrUReqOrURes

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	req.CategoryId = id
	res, err := h.Category.Update(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't update category ", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// CategoryDelete handles deleting a category  by ID.
// @Summary Delete category
// @Description Delete a category  by ID
// @Tags category
// @Accept json
// @Produce json
// @Param id path string true "Category ID"
// @Success 200 {object} string "Category  deleted"
// @Failure 400 {object} string "Invalid category  ID"
// @Failure 404 {object} string "Category  not found"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /category/{id} [delete]
func (h *HTTPHandler) CategoryDelete(c *gin.Context) {
	id := &pb.CategoryGReqOrDReq{CategoryId: c.Param("id")}
	_, err := h.Category.Delete(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't delete category ", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Category deleted"})
}

// CategoryGetAll handles getting all category s.
// @Summary Get all categories
// @Description Get all categories
// @Tags category
// @Accept json
// @Produce json
// @Param category_id query string false "category_id"
// @Param limit query integer false "limit"
// @Param offset query integer false "offset"
// @Success 200 {object} pb.CategoryGARes
// @Failure 400 {object} string "Invalid parameters"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /categories [get]
func (h *HTTPHandler) CategoryGetAll(c *gin.Context) {
	limitStr := c.Query("limit")
	offsetStr := c.Query("offset")
	categortId := c.Query("category_id")
	var limit, offset int
	var err error
	if limitStr == "" {
		limit = 0
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

	res, err := h.Category.GetAll(c, &pb.CategoryGAReq{
		Filter: &pb.CategoryFilter{
			CategoryId: categortId,
		},
		Pagination: &pb.Pagination{
			Limit:  int64(limit),
			Offset: int64(offset),
		},
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't get categories", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}
