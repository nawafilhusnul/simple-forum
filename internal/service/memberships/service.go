package memberships

import (
	"context"

	"github.com/nawafilhusnul/forum/internal/configs"
	membershipsModel "github.com/nawafilhusnul/forum/internal/model/memberships"
)

type membershipRepository interface {
	GetUser(ctx context.Context, email, userName string, userID int64) (*membershipsModel.UserModel, error)
	CreateUser(ctx context.Context, user *membershipsModel.UserModel) error
	GetRefreshTokenByUserID(ctx context.Context, userID int64) (*membershipsModel.RefreshTokenModel, error)
	InsertRefreshToken(ctx context.Context, refreshToken membershipsModel.RefreshTokenModel) error
}

type service struct {
	membershipRepo membershipRepository
	cfg            *configs.Config
}

func NewService(membershipRepo membershipRepository, cfg *configs.Config) *service {
	return &service{membershipRepo: membershipRepo, cfg: cfg}
}
