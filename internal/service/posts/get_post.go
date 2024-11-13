package posts

import (
	"context"
	"errors"

	postModel "github.com/nawafilhusnul/forum/internal/model/posts"
	"github.com/rs/zerolog/log"
)

func (s *service) GetPostByID(ctx context.Context, userID int64, postID int64) (postModel.GetPostByIDResponse, error) {
	post, err := s.postRepo.GetPostByID(ctx, userID, postID)
	if err != nil {
		log.Error().Err(err).Msgf("error get post by id %d", postID)
		return postModel.GetPostByIDResponse{}, err
	}

	if post == nil {
		log.Error().Msgf("post with id %d not found", postID)
		return postModel.GetPostByIDResponse{}, errors.New("post not found")
	}

	likeCount, err := s.postRepo.CountLikesByPostID(ctx, postID)
	if err != nil {
		log.Error().Err(err).Msgf("error count likes by post id %d", postID)
		return postModel.GetPostByIDResponse{}, err
	}

	comments, err := s.postRepo.GetPostComments(ctx, postID)
	if err != nil {
		log.Error().Err(err).Msgf("error get post comments with post id %d", postID)
		return postModel.GetPostByIDResponse{}, err
	}

	return postModel.GetPostByIDResponse{
		PostDetail: *post,
		LikeCount:  likeCount,
		Comments:   comments,
	}, nil
}
