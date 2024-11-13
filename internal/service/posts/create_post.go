package posts

import (
	"context"
	"strings"

	postsModel "github.com/nawafilhusnul/forum/internal/model/posts"
	"github.com/rs/zerolog/log"
)

func (s *service) CreatePost(ctx context.Context, userID int64, req postsModel.CreatePostRequest) error {
	postHastags := strings.Join(req.Hashtags, ",")

	post := postsModel.PostModel{
		UserID:    userID,
		Title:     req.Title,
		Content:   req.Content,
		Hashtags:  postHastags,
		CreatedBy: userID,
		UpdatedBy: userID,
	}

	err := s.postRepo.CreatePost(ctx, post)
	if err != nil {
		log.Error().Err(err).Msgf("failed to create post by user %d\n", userID)
		return err
	}

	return nil
}
