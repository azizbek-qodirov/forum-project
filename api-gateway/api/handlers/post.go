// post.go
package handlers

import (
	"context"
	"net/http"
	"strconv"

	pb "api-gateway/forum-protos/genprotos"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// PostCreate handles the creation of a new post.
// @Summary Create post
// @Description Create a new post
// @Tags post
// @Accept json
// @Produce json
// @Param post body pb.PostCReqForSwagger true "Post data"
// @Success 200 {object} pb.PostCReqOrCResOrGResOrUResp
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /post [POST]
func (h *HTTPHandler) PostCreate(c *gin.Context) {
	var req pb.PostCReqOrCResOrGResOrUResp
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
	res, err := h.Post.Create(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// PostGet handles getting a post by ID.
// @Summary Get post
// @Description Get a post by ID
// @Tags post
// @Accept json
// @Produce json
// @Param id path string true "Post ID"
// @Success 200 {object} pb.PostCReqOrCResOrGResOrUResp
// @Failure 400 {object} string "Invalid post  ID"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /post/{id} [GET]
func (h *HTTPHandler) PostGet(c *gin.Context) {
	id := &pb.PostGReqOrDReq{PostId: c.Param("id")}
	res, err := h.Post.GetByID(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't get post", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// PostUpdate handles updating an existing post .
// @Summary Update post
// @Description Update an existing post
// @Tags post
// @Accept json
// @Produce json
// @Param id path string true "Post ID"
// @Param post body pb.PostCReqForSwagger true "Updated post data"
// @Success 200 {object} pb.PostCReqOrCResOrGResOrUResp
// @Failure 400 {object} string "Invalid request payload"
// @Failure 404 {object} string "Post not found"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /post/{id} [put]
func (h *HTTPHandler) PostUpdate(c *gin.Context) {
	id := c.Param("id")
	var req pb.PostUReq

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	req.PostId = id
	res, err := h.Post.Update(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't update post ", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// PostDelete handles deleting a post  by ID.
// @Summary Delete post
// @Description Delete a post  by ID
// @Tags post
// @Accept json
// @Produce json
// @Param id path string true "Post ID"
// @Success 200 {object} string "Post  deleted"
// @Failure 400 {object} string "Invalid post  ID"
// @Failure 404 {object} string "Post  not found"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /post/{id} [DELETE]
func (h *HTTPHandler) PostDelete(c *gin.Context) {
	id := &pb.PostGReqOrDReq{PostId: c.Param("id")}
	_, err := h.Post.Delete(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't delete post ", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Post deleted"})
}

// PostGetAll handles getting all post s.
// @Summary Get all posts
// @Description Get all posts
// @Tags post
// @Accept json
// @Produce json
// @Param user_id query string false "user_id"
// @Param category_id query string false "category_id"
// @Param title query string false "title"
// @Param body query string false "body"
// @Param tags query string false "tags"
// @Param limit query integer false "limit"
// @Param offset query integer false "offset"
// @Success 200 {object} pb.PostGARes
// @Failure 400 {object} string "Invalid parameters"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /posts [GET]
func (h *HTTPHandler) PostGetAll(c *gin.Context) {
	limitStr := c.Query("limit")
	offsetStr := c.Query("offset")
	categoryId := c.Query("category_id")
	userId := c.Query("user_id")
	title := c.Query("title")
	body := c.Query("body")
	tags := c.Query("tags")

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

	res, err := h.Post.GetAll(context.Background(), &pb.PostGAReq{
		Filter: &pb.PostFilter{
			UserId:     userId,
			CategoryId: categoryId,
			Title:      title,
			Body:       body,
			Tags:       tags,
		},
		Pagination: &pb.Pagination{
			Limit:  int64(limit),
			Offset: int64(offset),
		},
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't get posts", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}
