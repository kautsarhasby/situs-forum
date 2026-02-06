package posts

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kautsarhasby/situs-forum/internal/model/posts"
)

func (h *Handler) UpsertUserActivity(c *gin.Context) {
	ctx := c.Request.Context()

	var request posts.UserActivityRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	userID := c.GetInt64("userID")
	postIDStr := c.Param("postID")
	postID, err := strconv.ParseInt(postIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": errors.New("PostID on params invalid"),
		})
		return
	}

	if err := h.postSVC.UpsertUserActivity(ctx, postID, userID, request); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": errors.New("Failed to update or insert"),
		})
		return
	}

	c.Status(http.StatusOK)
}
