package memberships

import (
	"context"
	"errors"
	"time"

	membershipsModel "github.com/nawafilhusnul/forum/internal/model/memberships"
	"github.com/nawafilhusnul/forum/pkg/jwt"
	"github.com/nawafilhusnul/forum/pkg/token"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

func (s *service) SignIn(ctx context.Context, req *membershipsModel.SignInRequest) (string, string, error) {
	user, err := s.membershipRepo.GetUser(ctx, req.Email, "", 0)
	if err != nil {
		log.Error().Err(err).Msgf("failed to get user with email %s\n", req.Email)
		return "", "", errors.New("failed to get user")
	}

	if user == nil {
		log.Error().Msgf("email %s is not registered\n", req.Email)
		return "", "", errors.New("email is not registered")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		log.Error().Err(err).Msgf("email %s or password is incorrect\n", req.Email)
		return "", "", errors.New("email or password is incorrect")
	}

	accessToken, err := jwt.CreateToken(user.ID, user.UserName, s.cfg.Service.SecretJWT)
	if err != nil {
		log.Error().Err(err).Msgf("failed to create token for user %s\n", user.UserName)
		return "", "", errors.New("failed to create token")
	}

	existingRefreshToken, err := s.membershipRepo.GetRefreshTokenByUserID(ctx, user.ID)
	if err != nil {
		log.Error().Err(err).Msgf("failed to get refresh token for user %s\n", user.UserName)
		return "", "", errors.New("failed to get refresh token")
	}

	if existingRefreshToken != nil {
		return accessToken, existingRefreshToken.RefreshToken, nil
	}

	newRefreshToken := token.GenerateRefreshToken(user.ID)
	if newRefreshToken == "" {
		log.Error().Msgf("failed to generate refresh token for user %s\n", user.UserName)
		return "", "", errors.New("failed to generate refresh token")
	}
	refreshToken := membershipsModel.RefreshTokenModel{
		UserID:       user.ID,
		RefreshToken: newRefreshToken,
		ExpiredAt:    time.Now().AddDate(0, 0, 10),
		CreatedBy:    user.ID,
		UpdatedBy:    user.ID,
	}
	err = s.membershipRepo.InsertRefreshToken(ctx, refreshToken)
	if err != nil {
		log.Error().Err(err).Msgf("failed to insert refresh token for user %s\n", user.UserName)
		return "", "", errors.New("failed to insert refresh token")
	}

	return accessToken, newRefreshToken, nil
}
