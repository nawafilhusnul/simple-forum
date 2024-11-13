package memberships

import (
	"context"
	"errors"
	"strings"

	membershipsModel "github.com/nawafilhusnul/forum/internal/model/memberships"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

func (s *service) SignUp(ctx context.Context, req *membershipsModel.SignUpRequest) error {
	req.Email = strings.ToLower(req.Email)
	req.UserName = strings.ToLower(req.UserName)

	user, err := s.membershipRepo.GetUser(ctx, req.Email, req.UserName, 0)
	if err != nil {
		log.Error().Err(err).Msgf("failed to get user with email %s or username %s\n", req.Email, req.UserName)
		return errors.New("failed to get user")
	}

	if user != nil {
		log.Error().Msgf("username %s or email %s already exists\n", req.UserName, req.Email)
		return errors.New("username or email already exists")
	}

	pass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Error().Err(err).Msgf("failed to generate password for user %s\n", req.UserName)
		return errors.New("failed to generate password")
	}

	model := &membershipsModel.UserModel{
		Name:      req.Name,
		UserName:  req.UserName,
		Email:     req.Email,
		Password:  string(pass),
		CreatedBy: 0,
		UpdatedBy: 0,
	}

	err = s.membershipRepo.CreateUser(ctx, model)
	if err != nil {
		log.Error().Err(err).Msgf("failed to create user with email %s and username %s\n", req.Email, req.UserName)
		return errors.New("failed to create user")
	}

	return nil
}
