package posts

import (
	"context"

	"github.com/nawafilhusnul/forum/internal/configs"
	postModel "github.com/nawafilhusnul/forum/internal/model/posts"
)

type postRepository interface {
	CreatePost(ctx context.Context, post postModel.PostModel) error
	CreatePostComment(ctx context.Context, req *postModel.PostCommentModel) error
	CreateUserPostActivity(ctx context.Context, userID int64, postID int64, req postModel.UserPostActivityModel) error
	GetUserPostActivity(ctx context.Context, userID int64, postID int64) (*postModel.UserPostActivityModel, error)
	UpdateUserPostActivity(ctx context.Context, userID int64, postID int64, req postModel.UserPostActivityModel) error
	GetAllPosts(ctx context.Context, limit, offset int) (postModel.GetAllPostsResponse, error)
	GetPostByID(ctx context.Context, userID, postID int64) (*postModel.PostResponse, error)
	GetPostComments(ctx context.Context, postID int64) ([]postModel.CommentResponse, error)
	CountLikesByPostID(ctx context.Context, postID int64) (int64, error)
}

type service struct {
	postRepo postRepository
	cfg      *configs.Config
}

func NewService(postRepo postRepository, cfg *configs.Config) *service {
	return &service{postRepo: postRepo, cfg: cfg}
}
