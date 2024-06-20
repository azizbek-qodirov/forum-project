package handlers

import (
	pb "api-gateway/forum-protos/genprotos"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// TagGet handles getting popular tags.
// @Summary Get Popular tags
// @Description Gets popular tags
// @Tags tag
// @Accept json
// @Produce json
// @Param limit query integer false "limit"
// @Param offset query integer false "offset"
// @Success 200 {object} pb.TagGAResOrPopularRes
// @Failure 400 {object} string "Invalid tag  ID"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /popular-tags [GET]
func (h *HTTPHandler) PopularTagsGet(c *gin.Context) {
	limitStr := c.Query("limit")
	offsetStr := c.Query("offset")

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

	tags, err := h.Tag.GetPopular(c, &pb.Pagination{Limit: int64(limit), Offset: int64(offset)})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get popular tags", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tags)
}
