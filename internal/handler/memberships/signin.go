package memberships

import (
	"net/http"

	"github.com/gin-gonic/gin"
	membershipsModel "github.com/nawafilhusnul/forum/internal/model/memberships"
)

func (h *Handler) SignIn(c *gin.Context) {
	ctx := c.Request.Context()

	var req membershipsModel.SignInRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	accessToken, refreshToken, err := h.membershipSvc.SignIn(ctx, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	res := membershipsModel.SignInResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	c.JSON(http.StatusOK, res)
}
