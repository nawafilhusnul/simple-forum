package memberships

import (
	"context"
	"errors"

	membershipsModel "github.com/nawafilhusnul/forum/internal/model/memberships"
	"github.com/nawafilhusnul/forum/pkg/jwt"
	"github.com/nawafilhusnul/forum/pkg/token"
	"github.com/rs/zerolog/log"
)

func (s *service) ValidateRefreshToken(ctx context.Context, userID int64, req membershipsModel.ValidateRefreshTokenRequest) (string, error) {
	existingRefreshToken, err := s.membershipRepo.GetRefreshTokenByUserID(ctx, userID)
	if err != nil {
		log.Error().Err(err).Msgf("failed to get refresh token for user %d\n", userID)
		return "", errors.New("refresh token not found")
	}

	if existingRefreshToken == nil {
		log.Error().Msgf("refresh token not found for user %d\n", userID)
		return "", errors.New("refresh token not found")
	}

	if existingRefreshToken.RefreshToken != req.RefreshToken {
		log.Error().Msgf("refresh token is invalid for user %d\n", userID)
		return "", errors.New("refresh token is invalid")
	}

	newRefreshToken := token.GenerateRefreshToken(userID)
	if newRefreshToken == "" {
		log.Error().Msgf("failed to generate refresh token for user %d\n", userID)
		return "", errors.New("failed to generate refresh token")
	}

	user, err := s.membershipRepo.GetUser(ctx, "", "", userID)
	if err != nil {
		log.Error().Err(err).Msgf("failed to get user for user %d\n", userID)
		return "", errors.New("failed to get user")
	}

	if user == nil {
		log.Error().Msgf("user not found for user %d\n", userID)
		return "", errors.New("user not found")
	}

	accessToken, err := jwt.CreateToken(user.ID, user.UserName, s.cfg.Service.SecretJWT)
	if err != nil {
		log.Error().Err(err).Msgf("failed to create access token for user %d\n", userID)
		return "", errors.New("failed to create access token")
	}

	return accessToken, nil
}
