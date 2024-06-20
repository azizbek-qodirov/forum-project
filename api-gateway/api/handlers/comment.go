// comment.go
package handlers

import (
	"context"
	"net/http"
	"strconv"

	pb "api-gateway/forum-protos/genprotos"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// CommentCreate handles the creation of a new comment.
// @Summary Create comment
// @Description Create a new comment
// @Tags comment
// @Accept json
// @Produce json
// @Param comment body pb.CommentCReqForSwagger true "Comment data"
// @Success 200 {object} pb.CommentCReqOrCResOrGResOrURes
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /comment [post]
func (h *HTTPHandler) CommentCreate(c *gin.Context) {
	var req pb.CommentCReqOrCResOrGResOrURes
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	claims, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	user_id := claims.(jwt.MapClaims)["user_id"].(string)
	req.UserId = user_id
	res, err := h.Comment.Create(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// CommentGet handles getting a comment by ID.
// @Summary Get comment
// @Description Get a comment by ID
// @Tags comment
// @Accept json
// @Produce json
// @Param id path string true "Comment ID"
// @Success 200 {object} pb.CommentCReqOrCResOrGResOrURes
// @Failure 400 {object} string "Invalid comment  ID"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /comment/{id} [get]
func (h *HTTPHandler) CommentGet(c *gin.Context) {
	id := &pb.CommentGReqOrDReq{CommentId: c.Param("id")}
	res, err := h.Comment.GetByID(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't get comment", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// CommentUpdate handles updating an existing comment .
// @Summary Update comment
// @Description Update an existing comment
// @Tags comment
// @Accept json
// @Produce json
// @Param id path string true "Comment ID"
// @Param comment body pb.CommentCReqForSwagger true "Updated comment data"
// @Success 200 {object} pb.CommentCReqOrCResOrGResOrURes
// @Failure 400 {object} string "Invalid request payload"
// @Failure 404 {object} string "Comment not found"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /comment/{id} [put]
func (h *HTTPHandler) CommentUpdate(c *gin.Context) {
	id := c.Param("id")
	var req pb.CommentUReq

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	req.CommentId = id
	res, err := h.Comment.Update(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't update comment ", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// CommentDelete handles deleting a comment  by ID.
// @Summary Delete comment
// @Description Delete a comment  by ID
// @Tags comment
// @Accept json
// @Produce json
// @Param id path string true "Comment ID"
// @Success 200 {object} string "Comment  deleted"
// @Failure 400 {object} string "Invalid comment  ID"
// @Failure 404 {object} string "Comment  not found"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /comment/{id} [delete]
func (h *HTTPHandler) CommentDelete(c *gin.Context) {
	id := &pb.CommentGReqOrDReq{CommentId: c.Param("id")}
	_, err := h.Comment.Delete(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't delete comment ", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Comment deleted"})
}

// CommentGetAll handles getting all comment s.
// @Summary Get all comments
// @Description Get all comments
// @Tags comment
// @Accept json
// @Produce json
// @Param user_id query string false "user_id"
// @Param post_id query string false "post_id"
// @Param limit query integer false "Limit"
// @Param offset query integer false "Offset"
// @Success 200 {object} pb.CommentGARes
// @Failure 400 {object} string "Invalid parameters"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /comments [GET]
func (h *HTTPHandler) CommentGetAll(c *gin.Context) {
	limitStr := c.Query("limit")
	offsetStr := c.Query("offset")
	postId := c.Query("post_id")
	userId := c.Query("user_id")
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

	res, err := h.Comment.GetAll(context.Background(), &pb.CommentGAReq{
		Filter: &pb.CommentFilter{
			PostId: postId,
			UserId: userId,
		},
		Pagination: &pb.Pagination{
			Limit:  int64(limit),
			Offset: int64(offset),
		},
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't get comments", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}
