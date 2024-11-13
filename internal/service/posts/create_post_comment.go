package posts

import (
	"context"

	postModel "github.com/nawafilhusnul/forum/internal/model/posts"
	"github.com/rs/zerolog/log"
)

func (s *service) CreatePostComment(ctx context.Context, userID int64, postID int64, req postModel.CreatePostCommentRequest) error {
	comment := postModel.PostCommentModel{
		PostID:    postID,
		Comment:   req.Comment,
		CreatedBy: userID,
		UpdatedBy: userID,
	}

	err := s.postRepo.CreatePostComment(ctx, &comment)
	if err != nil {
		log.Error().Err(err).Msgf("failed to create post comment for post %d by user %d\n", postID, userID)
		return err
	}

	return nil
}
