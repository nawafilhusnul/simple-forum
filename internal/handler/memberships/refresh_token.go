package memberships

import (
	"net/http"

	"github.com/gin-gonic/gin"
	membershipsModel "github.com/nawafilhusnul/forum/internal/model/memberships"
)

func (h *Handler) Refresh(c *gin.Context) {
	ctx := c.Request.Context()

	var req membershipsModel.ValidateRefreshTokenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := c.GetInt64("user_id")

	accessToken, err := h.membershipSvc.ValidateRefreshToken(ctx, userID, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	res := membershipsModel.ValidateRefreshTokenResponse{
		AccessToken: accessToken,
	}

	c.JSON(http.StatusOK, res)
}
