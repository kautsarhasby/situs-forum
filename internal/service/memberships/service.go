package memberships

import (
	"context"
	"time"

	"github.com/kautsarhasby/situs-forum/internal/configs"
	"github.com/kautsarhasby/situs-forum/internal/model/memberships"
)

type membershipRepository interface {
	GetUser(ctx context.Context, email, username string, id int64) (*memberships.UserModel, error)
	CreateUser(ctx context.Context, model memberships.UserModel) error
	GetRefreshToken(ctx context.Context, userID int64, now time.Time) (*memberships.RefreshTokenModel, error)
	InsertRefreshToken(ctx context.Context, model memberships.RefreshTokenModel) error
}

type service struct {
	cfg            *configs.Config
	membershipRepo membershipRepository
}

func NewService(membershipRepo membershipRepository, cfg *configs.Config) *service {
	return &service{
		cfg:            cfg,
		membershipRepo: membershipRepo,
	}
}
