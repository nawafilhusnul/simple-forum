package posts

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetAllPosts(c *gin.Context) {
	pageSize, err := strconv.Atoi(c.Query("page_size"))
	if err != nil || pageSize <= 0 {
		pageSize = 10
	}

	pageIndex, err := strconv.Atoi(c.Query("page_index"))
	if err != nil || pageIndex <= 0 {
		pageIndex = 1
	}

	res, err := h.postSvc.GetAllPosts(c.Request.Context(), pageSize, pageIndex)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}
