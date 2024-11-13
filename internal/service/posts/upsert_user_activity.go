package posts

import (
	"context"
	"errors"

	postModel "github.com/nawafilhusnul/forum/internal/model/posts"
	"github.com/rs/zerolog/log"
)

func (s *service) UpsertUserPostActivity(ctx context.Context, userID int64, postID int64, req postModel.UpsertUserPostActivityRequest) error {

	model := postModel.UserPostActivityModel{
		UserID:    userID,
		PostID:    postID,
		IsLiked:   req.IsLiked,
		CreatedBy: userID,
		UpdatedBy: userID,
	}

	userPostActivity, err := s.postRepo.GetUserPostActivity(ctx, userID, postID)
	if err != nil {
		log.Error().Err(err).Msgf("failed to get user post activity for user %d and post %d\n", userID, postID)
		return err
	}

	if userPostActivity == nil {
		if !req.IsLiked {
			log.Info().Msgf("user %d can't unlike post %d because user never liked it before\n", userID, postID)
			return errors.New("user never liked the post before")
		}

		err = s.postRepo.CreateUserPostActivity(ctx, userID, postID, model)
		if err != nil {
			log.Error().Err(err).Msgf("failed to create user post activity for user %d and post %d\n", userID, postID)
			return err
		}

		return nil
	}

	err = s.postRepo.UpdateUserPostActivity(ctx, userID, postID, model)
	if err != nil {
		log.Error().Err(err).Msgf("failed to update user post activity for user %d and post %d\n", userID, postID)
		return err
	}

	return nil
}
