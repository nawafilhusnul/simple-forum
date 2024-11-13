package posts

import (
	"context"

	postModel "github.com/nawafilhusnul/forum/internal/model/posts"
	"github.com/rs/zerolog/log"
)

func (s *service) GetAllPosts(ctx context.Context, pageSize, pageIndex int) (postModel.GetAllPostsResponse, error) {
	limit := pageSize
	offset := (pageIndex - 1) * pageSize

	res, err := s.postRepo.GetAllPosts(ctx, limit, offset)
	if err != nil {
		log.Error().Err(err).Msgf("failed to get all posts")
		return postModel.GetAllPostsResponse{}, err
	}

	return res, nil
}
