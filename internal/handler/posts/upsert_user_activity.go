package posts

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	postsModel "github.com/nawafilhusnul/forum/internal/model/posts"
)

func (h *Handler) UpsertUserPostActivity(c *gin.Context) {
	ctx := c.Request.Context()

	userID := c.GetInt64("user_id")
	if userID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not found"})
		return
	}

	postID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid post id"})
		return
	}

	var req postsModel.UpsertUserPostActivityRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.postSvc.UpsertUserPostActivity(ctx, userID, postID, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusCreated)
}
