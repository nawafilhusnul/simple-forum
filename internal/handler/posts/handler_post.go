package posts

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/nawafilhusnul/forum/internal/middleware"
	postsModel "github.com/nawafilhusnul/forum/internal/model/posts"
)

type postService interface {
	CreatePost(ctx context.Context, userID int64, req postsModel.CreatePostRequest) error
	CreatePostComment(ctx context.Context, userID, postID int64, req postsModel.CreatePostCommentRequest) error
	UpsertUserPostActivity(ctx context.Context, userID, postID int64, req postsModel.UpsertUserPostActivityRequest) error
	GetAllPosts(ctx context.Context, pageSize, pageIndex int) (postsModel.GetAllPostsResponse, error)
	GetPostByID(ctx context.Context, userID, postID int64) (postsModel.GetPostByIDResponse, error)
}

type Handler struct {
	*gin.Engine
	postSvc postService
}

func NewHandler(api *gin.Engine, postSvc postService) *Handler {
	return &Handler{
		Engine:  api,
		postSvc: postSvc,
	}
}

func (h *Handler) RegisterRoutes() {
	route := h.Group("/posts")
	route.Use(middleware.AuthMiddleware())

	route.POST("", h.CreatePost)
	route.POST("/:id/comments", h.CreatePostComment)
	route.POST("/:id/user-activity", h.UpsertUserPostActivity)
	route.GET("", h.GetAllPosts)
	route.GET("/:id", h.GetPostByID)
}
