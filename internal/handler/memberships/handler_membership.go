package memberships

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nawafilhusnul/forum/internal/middleware"
	membershipsModel "github.com/nawafilhusnul/forum/internal/model/memberships"
)

type membershipService interface {
	SignUp(ctx context.Context, req *membershipsModel.SignUpRequest) error
	SignIn(ctx context.Context, req *membershipsModel.SignInRequest) (string, string, error)
	ValidateRefreshToken(ctx context.Context, userID int64, req membershipsModel.ValidateRefreshTokenRequest) (string, error)
}

type Handler struct {
	*gin.Engine
	membershipSvc membershipService
}

func NewHandler(api *gin.Engine, membershipSvc membershipService) *Handler {
	return &Handler{
		Engine:        api,
		membershipSvc: membershipSvc,
	}
}

func (h *Handler) RegisterRoutes() {
	route := h.Group("/memberships")
	route.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	route.POST("/signup", h.SignUp)
	route.POST("/signin", h.SignIn)
	routeRefresh := route.Group("/refresh")
	routeRefresh.Use(middleware.AuthRefreshMiddleware())
	routeRefresh.POST("", h.Refresh)
}
